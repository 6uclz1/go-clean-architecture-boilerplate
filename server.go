package main

import (
	"fmt"
	"net/http"
)

func main() {
    // ハンドラ関数を定義
    helloHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    }

    // ルートパス ("/") へのリクエストに対するハンドラを設定
    http.HandleFunc("/", helloHandler)

    // サーバーを起動
    port := 8080
    fmt.Printf("Server is listening on port %d...\n", port)
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        panic(err)
    }
}