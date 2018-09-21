package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/auth0/go-jwt-middleware"

	"github.com/dgrijalva/jwt-go"
)

var GetJwt = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// JWSヘッダーの登録
	token := jwt.New(jwt.SigningMethodHS256)

	// JWSペイロードのclaimsをセットしていく
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "taji"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 署名をしていく
	jwt, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	// JWTを返却
	w.Write([]byte(jwt))
})

// 以下でtokenをチェックする
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
