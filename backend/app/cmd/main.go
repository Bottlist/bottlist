package main

import (
	"fmt"
	"net/http"

	"github.com/Bottlist/bottlist/pkg/infra"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	db := infra.NewMySQLConnector()
	defer func() {
		err := db.Conn.Close()
		if err != nil {
			panic(fmt.Sprintf("DB connection close failed: %v", err))
		}
	}()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":4000"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
