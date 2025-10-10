package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte
var jwtExpHours = 24

func Configure(secret string, expHours int) {
	jwtSecret = []byte(secret)
	if expHours > 0 {
		jwtExpHours = expHours
	}
}

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Tipo  string `json:"tipo"`
	jwt.RegisteredClaims
}

func GerarJWT(id int, email, tipo string) (string, error) {
	claims := &Claims{
		ID:    id,
		Email: email,
		Tipo:  tipo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtExpHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ProjectGoAPI",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func ValidarJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
