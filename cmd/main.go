package main

import (
	"fmt"
	"gobdd/adapter/middleware/persistence"
	"gobdd/global"
	"gobdd/handler"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func init() {
	global.InitIsLocal(os.Getenv("IS_LOCAL") != "")
}

func main() {
	// データベース接続ソース文字列
	var dataSource string
	if global.IsLocal() {
		// ローカル環境の場合は docker-compose の設定でMySQL起動するので固定値
		dataSource = "localuser:localpass@tcp(127.0.0.1)/localdb?charset=utf8&parseTime=True&loc=Local"
	} else {
		// GCP - CloudSQL への接続情報は環境変数から取得
		var (
			connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
			user           = os.Getenv("CLOUDSQL_USER")
			password       = os.Getenv("CLOUDSQL_PASSWORD")
			database       = os.Getenv("CLOUDSQL_DATABASE")
		)
		dataSource = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=True", user, password, connectionName, database)
	}

	// RDB接続状態の管理用ツールを隠蔽するミドルウェア
	rdbMiddleware, err := persistence.NewRDBMiddleware(dataSource)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := rdbMiddleware.Close(); err != nil {
			panic(err)
		}
	}()

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(handler.PersistenceKey, rdbMiddleware)
			return next(c)
		}
	})

	handler.Routing(e)

	port := "8080"
	// Google App Engineにデプロイする場合、デプロイ先で適切なポートが設定される。
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}

	e.Logger.Fatal(e.Start(":" + port))
}
