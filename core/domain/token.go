package domain

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Token struct {
	jwt *jwt.Token
}

func NewToken() *Token {
	t := &Token{
		jwt: jwt.New(jwt.SigningMethodHS256),
	}

	return t
}

func NewFromString(tokenSigned string) (*Token, error) {
	token, _ := jwt.Parse(tokenSigned, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(viper.GetString("token.secret")), nil
	})

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return &Token{jwt: token}, nil
}

func (token *Token) SetClaim(name string, value any) {
	claims := token.jwt.Claims.(jwt.MapClaims)
	claims[name] = value
}

func (token *Token) GetClaim(name string) any {
	claims := token.jwt.Claims.(jwt.MapClaims)
	return claims[name]
}

func (token *Token) SignString() (string, error) {
	var sampleSecretKey = []byte(viper.GetString("token.secret"))
	return token.jwt.SignedString(sampleSecretKey)
}
