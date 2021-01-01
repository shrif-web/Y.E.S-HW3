package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"yes-blog/graph/model"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("Y.E.S at its best")
)

// GenerateToken generates a jwt token and store the username as a claim
func GenerateToken(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// producing the token string
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		message:="couldn't create the token"
		return "", model.InternalServerException{Message: &message}
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username from claims
func ParseToken(tokenStr string) (string, error) {

	//creating the token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err!=nil{
		message:="couldn't read the token"
		return "", model.InternalServerException{Message: &message}
	}

	// extracting username
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", model.InternalServerException{}
	}
}
