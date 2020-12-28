package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mi11km/zikanwarikun-back/config"
)

// secret key being used to sign tokens
var secretKey = []byte(config.Cfg.Server.JwtSecretKey)

// GenerateToken generates a jwt token and assign a id to it's claims and return it
func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix() // Expiration is 5 hours
	tokenString, err := token.SignedString(secretKey)
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
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		return id, nil
	} else {
		return "", err
	}
}

// RefreshToken reissue jwt token from given token
func RefreshToken(token string) (string, error) {
	id, err := ParseToken(token)
	if err != nil {
		log.Printf("action=refresh token, status=failed, err=failed to parse token")
		return "", fmt.Errorf("failed to parse given token")
	}
	refreshToken, err := GenerateToken(id)
	if err != nil {
		log.Printf("action=refresh token, status=failed, err=failed to generate token")
		return "", err
	}

	log.Printf("action=refresh token, status=success")
	return refreshToken, nil
}
