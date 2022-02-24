package handler

import (
	"github.com/go-chi/jwtauth"
	"time"
)

const (
	tokenTime = 12 * time.Hour
	signKey   = "dkr3!#mc349x#s3&74f12d"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(signKey), nil)
}

func generateJWT(id int) string {
	exp := time.Now().Add(tokenTime).Unix()
	claims := map[string]interface{}{"id": id, "exp": exp}

	_, tokenString, _ := tokenAuth.Encode(claims)

	return tokenString
}
