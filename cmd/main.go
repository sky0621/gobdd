package main

import (
	"errors"
	"fmt"
	"gobdd/adapter/middleware/persistence"
	"net/http"
	"os"

	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

const PersistenceKey = "PERSISTENCE"

func main() {
	// データベース接続ソース文字列
	var dataSource string
	if os.Getenv("IS_LOCAL") == "" {
		// GCP - CloudSQL への接続情報は環境変数から取得
		var (
			connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
			user           = os.Getenv("CLOUDSQL_USER")
			password       = os.Getenv("CLOUDSQL_PASSWORD")
			database       = os.Getenv("CLOUDSQL_DATABASE")
		)
		dataSource = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=True", user, password, connectionName, database)
	} else {
		// ローカル環境の場合は docker-compose の設定でMySQL起動するので固定値
		dataSource = "localuser:localpass@tcp(127.0.0.1)/localdb?charset=utf8&parseTime=True&loc=Local"
	}

	persistentMiddleware, err := persistence.NewRDBMiddleware(dataSource)
	if err != nil {
		panic(err)
	}

	// WebアプリケーションフレームワークとしてEchoを利用
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(PersistenceKey, persistentMiddleware)
			return next(c)
		}
	})

	http.Handle("/", e)
	e.POST("/notice", createNotice)

	port := "8080"
	// Google App Engineにデプロイする場合、デプロイ先で適切なポートが設定される。
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func createNotice(c echo.Context) error {
	contextVal := c.Get(PersistenceKey)
	rdbMiddleware, ok := contextVal.(NewRDBMiddleware)
	if !ok {
		return errors.New("abnormal end")
	}

	// FIXME:
	return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}
