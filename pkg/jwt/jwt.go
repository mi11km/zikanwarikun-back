package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
// todo decide secret key from env variable?
var (
	SecretKey = []byte("CLKVHA|2MxH9~t6grxYB3JdB")
)

// GenerateToken generates a jwt token and assign a id to it's claims and return it
func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix() // Expiration is 5 hours
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Printf("action=Error in Generating key, err=%s", err)
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the id in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		return id, nil
	} else {
		return "", err
	}
}