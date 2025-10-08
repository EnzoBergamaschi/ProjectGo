package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte
var jwtExpHours = 24

// Configura o segredo e tempo de expiração do token JWT
func Configure(secret string, expHours int) {
	jwtSecret = []byte(secret)
	if expHours > 0 {
		jwtExpHours = expHours
	}
}

// Claims representa as informações contidas no token JWT
type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Tipo  string `json:"tipo"`
	jwt.RegisteredClaims
}

// GerarJWT cria um token JWT com ID, email e tipo de usuário (admin/cliente)
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

// ValidarJWT valida o token JWT e retorna as claims (dados do usuário)
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
