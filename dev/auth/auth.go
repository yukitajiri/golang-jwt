package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)


getJwt := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// JWSヘッダーの登録
	token := jwt.New(jwt.SigningMethodHS256)

	// JWSペイロードのclaimsをセットしていく
	claims := token.Claims.(jwt.MapClaims)

})