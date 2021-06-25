package userHandler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// add a user to the db
func (r *DbLog) signUpDB(user *UserModel) (string, error) {

	// validate the input data
	err := user.Validate()
	if err != nil {
		return "", err
	}

	user.Username = strings.ToLower(user.Username)

	// hash the password
	user.Password, _ = encryption.HashPassword(user.Password)

	// add the user to the db
	id, err := r.Db.InsertOne(context.TODO(), user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			return "", logger.ErrDuplicateKey
		}
		return "", logger.ErrDbWrite
	}

	r.L.DbWrite("created a new user")

	// convert the id to a string
	userId := fmt.Sprintf("%v", id.InsertedID)
	userId = userId[10 : len(userId)-2]

	// generate the jwt
	token, err := encryption.GenerateJwt(userId, &encryption.TokenData{Verified: false, Moderator: false, Username: user.Username})
	if err != nil {
		return "", errors.New("unable generate the user token")
	}

	// return the jwt token
	return token, nil
}

// get the user data from the db
func (r *DbLog) getUserDB(token string) (*UserModel, error) {

	// get the token data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return nil, err
	}

	var user UserModel

	// find the user with the same username of the token
	err = r.Db.FindOne(context.TODO(), bson.M{"username": tokenData.Username}).Decode(&user)
	if err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.DbRead("readed the data of a user")

	// remove the password from the return
	user.Password = ""

	// check if the user is banned
	if user.Ban {
		return nil, logger.ErrUnauthorized
	}

	// return the user data
	return &user, nil
}

// check the username and the password and generate the user token
func (r *DbLog) logInDB(username string, password string) (string, error) {

	var user UserModel

	// convert the username to lowercase
	username = strings.ToLower(username)

	// get the user data
	err := r.Db.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", logger.ErrDbRead
	}

	r.L.DbRead("readed the data of a user")

	// check if the user is banned
	if user.Ban {
		return "", logger.ErrUnauthorized
	}

	// check if the password is correct
	err = encryption.ValidatePassword(password, user.Password)
	if err != nil {
		return "", err
	}

	// generate the jwt
	token, err := encryption.GenerateJwt(user.ID.Hex(), &encryption.TokenData{Verified: user.Dev.Verified, Moderator: user.Moderator, Username: username})
	if err != nil {
		return "", err
	}

	// return the token
	return token, nil
}

// edit the permission of a user
func (r *DbLog) userPermDB(token string, username string, perm string, value bool) error {

	// get the user data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// check if the user has the permission
	if !tokenData.Moderator && !tokenData.Verified {
		return logger.ErrUnauthorized
	}

	// check if the perm to modify is correct
	if perm == "ban" || perm == "dev.verified" || perm == "moderator" {

		// edit the user perm
		_, err := r.Db.UpdateOne(context.TODO(), bson.M{"username": strings.ToLower(username)}, bson.D{
			bson.E{"$set", bson.M{perm: value}},
		})
		if err != nil {
			return errors.New("unable to edit the user data")
		}

		return nil
	}

	return logger.ErrInvalidKey

}
