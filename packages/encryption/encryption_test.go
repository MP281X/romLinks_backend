package encryption

import (
	"os"
	"testing"
)

func TestHash(t *testing.T) {

	// encryption
	const password = "o$c8Ss!69@J*rtfU5&uE*V3!^"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Error(err)
	}

	// input and output
	if password == hashedPassword {
		t.Error("the hash is equal to the password")
	}

	// validation
	err = ValidatePassword(password, hashedPassword)
	if err != nil {
		t.Error("the password is invalid")
	}
}

func TestJWT(t *testing.T) {

	os.Setenv("jwtKey", "mk9SAmRoa7aEVTXGSw1Ptz6MP7B135")
	userId := "60d30a6d31b62f8cb6dc9dc9"

	tokenData := &TokenData{
		Username:  "testUsername",
		Verified:  true,
		Moderator: false,
	}

	token, err := GenerateJwt(userId, tokenData)
	if err != nil {
		t.Error(err)
	}

	data, err := GetTokenData(token)
	if err != nil {
		t.Error(err)
	}

	if data.Username != tokenData.Username || data.Moderator != tokenData.Moderator || data.Verified != tokenData.Verified {
		t.Error("the input data has changed")
	}
}
