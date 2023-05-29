package jwt

import (
	"bookstore/core/environment"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	env            *environment.Environment
	expirationTime time.Time
}

func New(env *environment.Environment) *JWT {
	expiresAt := time.Now().Add(12 * time.Hour)
	return &JWT{env: env, expirationTime: expiresAt}
}

type Claim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

type Response struct {
	ExpirationTime time.Time `json:"expirationTime"`
	Token          string    `json:"token"`
}

func (j *JWT) Generate(name string, email string, id string) (res Response, err error) {
	claims := &Claim{
		Id:    id,
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: j.expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.env.JWTSecret))
	if err != nil {
		return Response{}, err
	}
	res = Response{ExpirationTime: j.expirationTime, Token: tokenString}
	return res, nil
}

func (j *JWT) Validate(signedToken string) (email string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.env.JWTSecret), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", err
	}
	return claims.Email, nil

}
