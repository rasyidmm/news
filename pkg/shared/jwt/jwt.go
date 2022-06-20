package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtCustomClaims struct {
	LevelId  string `json:"level_id"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	jwt.RegisteredClaims
}

func CreateJwtToken(username, lvid string) string {
	SecretKey := "hala"
	claim := jwt.NewWithClaims(jwt.SigningMethodHS512, JwtCustomClaims{
		LevelId:  "",
		Username: "",
		UUID:     "",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour*8).Unix(), 0)),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	})

	token, err := claim.SignedString([]byte(SecretKey))
	if err != nil {

	}
	return token
}
