package main

import (
	"fmt"
	"github.com/Bottlist/bottlist/cmd/di"
	"github.com/Bottlist/bottlist/middlewares"
	"github.com/Bottlist/bottlist/pkg/adapter/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	app, _ := di.InitializeApp()
	defer func() {
		err := app.Db.Close()
		if err != nil {
			panic(fmt.Sprintf("DB connection close failed: %v", err))
		}
	}()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.CORS())

	noAuth := e.Group("/api")
	reqAuth := e.Group("/api")
	reqAuth.Use(app.Session.Session)

	// ルートを設定
	router.NewAuthRouter(noAuth, reqAuth, *app.AuthHandler).Router()
	router.NewHelloRouter(noAuth).Router()

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":4000"))
}
