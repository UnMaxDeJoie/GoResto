package managers

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var secretKey = []byte("secret-key")

func CreateToken(email string, uid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    uid,
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
