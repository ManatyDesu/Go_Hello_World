package main

import (
	"fmt"
	"net/http"
)

// Requestから情報を取り出し、HTTPレスポンスを作成 -> ResponseWriterがレスポンスを返信
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!!") // 引数としてwriterがある。
}

func main() {
	http.HandleFunc("/", handler)     // root URL (/)が呼び出されたらhandler関数を起動
	http.ListenAndServe(":8080", nil) // port 8080を監視するようサーバを起動
}
