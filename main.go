package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kojikawazu/backend/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// セットアップ
func firstSetup() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	firstSetup()
	router.SetupRoutes(e)

	// ルーティングの設定
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバーの起動
	e.Logger.Fatal(e.Start(":8080"))
}
