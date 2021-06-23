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
}

// generate a jwt token from the username
func GenerateJwt(username string, tokenData *TokenData) (string, error) {
	claims := &jwt.MapClaims{
		"iss": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]interface{}{
			"verified":  tokenData.Verified,
			"moderator": tokenData.Moderator,
		},
	}
	// create the jwt token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sing the token
	token, err := jwtToken.SignedString([]byte(os.Getenv("jwtKey")))
	if err != nil {
		return "", err
	}
	logger.Jwt("generated a new token for " + username)
	return token, nil
}

// get the token data from the token
func GetTokenData(token string) (map[string]interface{}, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtKey")), nil
	})
	if err != nil {
		return nil, err
	}
	// get the token data
	tokenClaims := jwtToken.Claims.(jwt.MapClaims)
	logger.Jwt("got the token data from a token")
	return tokenClaims["data"].(map[string]interface{}), nil
}

// get the token data from the token
func GetUserIdFromToken(token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtKey")), nil
	})
	if err != nil {
		return "", err
	}
	// get the token data
	tokenClaims := jwtToken.Claims.(jwt.MapClaims)
	logger.Jwt("got the user id from a token")
	return tokenClaims["iss"].(string), nil
}
