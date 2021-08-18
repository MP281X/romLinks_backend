package userHandler

import (
	"strings"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	SavedRom  []string           `bson:"savedRom" json:"savedRom"`
	Verified  bool               `bson:"verified" json:"verified"`
	Moderator bool               `bson:"moderator" json:"moderator"`
	Ban       bool               `bson:"ban" json:"ban"`
}

//validate the user data
func (user *UserModel) Validate() error {

	// validate username
	if user.Username == "" || len(user.Username) < 4 || len(user.Username) > 15 {
		return logger.ErrInvUsername
	}

	if strings.Contains(user.Username, "_") {
		return logger.ErrInvUsername
	}

	// validate email
	if user.Email == "" || !strings.Contains(user.Email, "@") || len(user.Email) < 10 {
		return logger.ErrInvEmail
	}
	// validate password
	if user.Password == "" || len(user.Password) < 6 || len(user.Password) > 20 {
		return logger.ErrInvPassword
	}

	// reset the other field
	user.SavedRom = []string{}
	user.Moderator = false
	user.Ban = false
	user.Verified = false

	return nil
}
