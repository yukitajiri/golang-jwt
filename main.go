package main

import (
	"encoding/json"
	"log"
	"net/http"

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

	//サーバー起動
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	content := &content{
		Title: "おはようございます",
		Body:  "今日もよろしくお願いします",
	}
	json.NewEncoder(w).Encode(content)
})
