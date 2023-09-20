package main

import (
	"fmt"
	"github.com/Bottlist/bottlist/cmd/di"
	"github.com/Bottlist/bottlist/external/mysql"
	"github.com/Bottlist/bottlist/middlewares"
	"github.com/Bottlist/bottlist/pkg/adapter/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	db := mysql.NewMySQLConnector()
	defer func() {
		err := db.Conn.Close()
		if err != nil {
			panic(fmt.Sprintf("DB connection close failed: %v", err))
		}
	}()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.CORS())

	app, _ := di.InitializeApp()

	// ルートを設定
	router.NewAuthRouter(e, app.AuthHandler).Router()
	router.NewHelloRouter(e).Router()

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":4000"))
}
