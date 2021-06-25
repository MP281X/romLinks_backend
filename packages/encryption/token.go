package encryption

import (
	"os"
	"time"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/dgrijalva/jwt-go"
)

// struct for the data to put in the token
type TokenData struct {
	Verified  bool
	Moderator bool
	Username  string
}

// generate a jwt token from the username
func GenerateJwt(userId string, tokenData *TokenData) (string, error) {

	claims := &jwt.MapClaims{
		"iss": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]interface{}{
			"verified":  tokenData.Verified,
			"moderator": tokenData.Moderator,
			"username":  tokenData.Username,
		},
	}

	// create the jwt token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sing the token
	token, err := jwtToken.SignedString([]byte(os.Getenv("jwtKey")))
	if err != nil {
		return "", logger.ErrTokenGen
	}

	return token, nil
}

// get the token data from the token
func GetTokenData(token string) (*TokenData, error) {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtKey")), nil
	})
	if err != nil {
		return nil, logger.ErrTokenRead
	}

	// get the token data from the claims
	tokenClaims := jwtToken.Claims.(jwt.MapClaims)
	tokenClaim := tokenClaims["data"].(map[string]interface{})

	// return a struct with the token data
	return &TokenData{
		Username:  tokenClaim["username"].(string),
		Moderator: tokenClaim["moderator"].(bool),
		Verified:  tokenClaim["verified"].(bool),
	}, nil
}
