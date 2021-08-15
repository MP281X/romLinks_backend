package userHandler

import (
	"context"
	"strings"
	"testing"

	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
)

func TestValidation(t *testing.T) {
	// null input data
	user := &UserModel{
		Username:  "",
		Email:     "",
		Password:  "",
		Verified:  true,
		Moderator: true,
		Ban:       true,
	}

	err := user.Validate()
	if err == nil {
		t.Error("the input was null")
	}

	// invalid input data
	user.Username = "Mark"
	user.Email = "ciaogmail.com"
	user.Password = "TestPassword"

	err = user.Validate()
	if err == nil {
		t.Error("the input was invalid")
	}

	user.Email = "ciao@gmail.com"

	// valid data
	err = user.Validate()
	if err != nil {
		t.Error("all the field is correct")
	}

	// invalid user perm
	if user.Moderator || user.Verified || user.Ban {
		t.Error("the moderator/verified/ban field has to be set to false ")
	}
}

func TestDBReq(t *testing.T) {

	// initialize the logger and the db
	d, _ := db.InitDB("test")
	l, _ := logger.InitLogger("test")
	c := d.Collection("test_user")
	r := &DbLog{Db: c, L: l}

	// clear the test collection
	c.Drop(context.TODO())

	// add the user
	token1, err := r.signUpDB(&UserModel{
		Username:  "testUsername",
		Email:     "test@email.com",
		Password:  "TestPassword",
		Verified:  true,
		Moderator: true,
		Ban:       true,
	})
	if err != nil {
		t.Error(err)
	}

	// get the user data
	userData1, err := r.getUserDB(token1)
	if err != nil {
		t.Error(err)
	}

	// check the fields
	if userData1.Username != strings.ToLower("testUsername") || userData1.Email != strings.ToLower("test@email.com") || userData1.Password == "TestPassword" {
		t.Error("invalid user data")
	}

	// generate the token with invalid data
	_, err = r.logInDB("testUsername", "TestIncorrectPassword")
	if err == nil {
		t.Error("the password was incorrect")
	}

	// delete the test db
	d.Drop(context.TODO())

}

func TestUserPerm(t *testing.T) {

	// initialize the logger and the db
	d, _ := db.InitDB("test")
	l, _ := logger.InitLogger("test")
	c := d.Collection("test_user")
	r := &DbLog{Db: c, L: l}

	// clear the test collection
	c.Drop(context.TODO())

	// generate a token with the moderator permission
	token, _ := encryption.GenerateJwt("test_user_id", &encryption.TokenData{Verified: true, Moderator: false, Username: "mp281x"})

	// add the user
	token2, err := r.signUpDB(&UserModel{
		Username:  "testUsername",
		Email:     "test@email.com",
		Password:  "TestPassword",
		Verified:  true,
		Moderator: true,
		Ban:       true,
	})
	if err != nil {
		t.Error(err)
	}

	// give the user mod perm
	err = r.userPermDB(token, "testUsername", "moderator", true)
	if err != nil {
		t.Error(err)
	}

	// get the user data
	userData, err := r.getUserDB(token2)
	if err != nil {
		t.Error(err)
	}

	// check if the moderator perm has changed
	if !userData.Moderator {
		t.Error("the user is a moderator")
	}

	// ban a user
	err = r.userPermDB(token, "testUsername", "ban", true)
	if err != nil {
		t.Error(err)
	}

	// check if the user is banned
	_, err = r.getUserDB(token2)
	if err == nil {
		t.Error("the user is banned but it can login")
	}

	// clear the test collection
	c.Drop(context.TODO())
}
