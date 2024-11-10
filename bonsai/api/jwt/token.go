package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	Roles string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(email string, roleid string, secretkey string) (string, error) {
	var mySigningKey = []byte(secretkey)

	claims := Claims{
		Email: email,
		Roles: roleid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Token hết hạn sau 72 giờ
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretkey string) (*Claims, error) {
	var mySigningKey = []byte(secretkey)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
