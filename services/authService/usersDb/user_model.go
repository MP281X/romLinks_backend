package usersdb

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	SavedRom  []string           `bson:"savedRom" json:"savedRom"`
	Dev       *DevModel          `bson:"dev" json:"dev"`
	Moderator bool               `bson:"moderator" json:"moderator"`
	Image     string             `bson:"image" json:"image"`
	Ban       bool               `bson:"ban" json:"ban"`
}

type DevModel struct {
	Link     []string `bson:"link" json:"link"`
	Verified bool     `bson:"verified" json:"verified"`
}

//validate the user data
func (user *UserModel) ValidateUserData() error {
	// validate username
	if user.Username == "" {
		return errors.New("enter a username")
	} else if len(user.Username) < 4 {
		return errors.New("username is too short")
	}

	// validate email
	if user.Email == "" {
		return errors.New("enter an email")
	} else if !strings.Contains(user.Email, "@") || len(user.Email) < 4 {
		return errors.New("invalid email")
	}

	// validate password
	if user.Password == "" {
		return errors.New("enter a password")
	} else if len(user.Password) < 6 {
		return errors.New("password is too short")
	}

	user.SavedRom = []string{}
	user.Moderator = false
	user.Ban = false
	user.Dev = &DevModel{
		Verified: false,
	}

	if len(user.Dev.Link) == 0 {
		user.Dev.Link = []string{}
	}

	return nil
}
