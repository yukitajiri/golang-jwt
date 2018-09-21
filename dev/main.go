package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yukitajiri/sample_golang_jwt/dev/auth"

	"github.com/gorilla/mux"
)

type content struct {
	Title string `json:"title""`
	Body  string `json:"body"`
}

func main() {
	r := mux.NewRouter()
	// localhost:8080/publicでpublicハンドラーを実行
	r.Handle("/", public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetJwt)

	//サーバー起動
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &content{
		Title: "こんにちは",
		Body:  "さよなら",
	}
	json.NewEncoder(w).Encode(post)
})

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	content := &content{
		Title: "おはようございます",
		Body:  "今日もよろしくお願いします",
	}
	json.NewEncoder(w).Encode(content)
})
