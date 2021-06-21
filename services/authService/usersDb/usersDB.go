package usersdb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection

// add a user to the db
func AddUser(user *UserModel) (string, error) {
	// validate the input data
	err := user.ValidateUserData()
	if err != nil {
		return "", err
	}
	user.Username = strings.ToLower(user.Username)
	// check is there is user with the same username
	var x interface{}
	res := UserCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&x)
	if res == nil {
		return "", errors.New("username already used")
	}

	// hash the password
	user.Password = encryption.HashPassword(user.Password)

	// add the user to the db
	id, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", errors.New("unable to add the user to the db")
	}
	userId := fmt.Sprintf("%v", id.InsertedID)
	// generate the jwt
	token, err := encryption.GenerateJwt(userId[10:len(userId)-2], &encryption.TokenData{Verified: false, Moderator: false})
	if err != nil {
		return "", errors.New("unable generate the user token")
	}
	logger.Info("new user: " + user.Username)
	// return the jwt token
	return token, nil
}

// get the user data from the db
func UserData(token string) (*UserModel, error) {
	var user UserModel
	id, err := encryption.GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}
	userId, _ := primitive.ObjectIDFromHex(id)
	err = UserCollection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user)
	user.Password = ""
	if user.Ban {
		return nil, errors.New("banned user")
	}
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	return &user, nil
}

// check the username and the password and generate the user token
func GenerateUserToken(username string, password string) (string, error) {
	var user UserModel
	// get the user data
	err := UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	// check if the user is banned
	if user.Ban {
		return "", errors.New("banned user")
	}

	// check if the password is correct
	err = encryption.ValidatePassword(password, user.Password)
	if err != nil {
		return "", err
	}

	// generate the jwt
	token, err := encryption.GenerateJwt(user.ID.Hex(), &encryption.TokenData{Verified: user.Dev.Verified, Moderator: user.Moderator})
	if err != nil {
		return "", errors.New("unable generate the user token")
	}

	// return the token
	logger.Info("user login: " + user.Username)
	return token, nil
}
