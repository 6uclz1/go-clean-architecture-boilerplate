// main.go

package main

import (
	"go-clean-architecture-boilerplate/internal/config"
	"go-clean-architecture-boilerplate/internal/delivery/http"
	"net/http"
)

func main() {
    // アプリケーション設定の読み込み
    appConfig := config.LoadConfig()

    // HTTPハンドラの初期化
    handler := http.NewHandler()

    // サーバーの起動
    httpServer := http.Server{
        Addr:    appConfig.ServerAddress,
        Handler: handler,
    }
    httpServer.ListenAndServe()
}
