package handler

import (
	"gobdd/adapter/gateway"
	"gobdd/adapter/middleware/persistence"
	"gobdd/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func HandleNotice(g *echo.Echo) {
	g.POST("/notice", createNotice)
}

func createNotice(c echo.Context) error {
	contextVal := c.Get(PersistenceKey)
	rdbMiddleware, ok := contextVal.(persistence.RDBMiddleware)
	if !ok {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	// HTTPリクエストパラメータ（JSON形式のBodyを想定）を構造体にマッピング
	var noticeParam *noticeForm
	if err := c.Bind(&noticeParam); err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	id, err := usecase.NewNotice(gateway.NewNotice(rdbMiddleware)).Create()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, http.StatusText(http.StatusInternalServerError))
}

type noticeForm struct {
	Title       string `json:"title"`        // お知らせのタイトル
	Text        string `json:"text"`         // お知らせの文章（現時点はテキストのみサポート）
	PublishFrom int    `json:"publish_from"` // お知らせの掲載開始日時
	PublishTo   int    `json:"publish_to"`   // お知らせの掲載終了日時
}

type response struct {
	Code    string `json:"code"` // HTTPステータスコード
	Message string `json:"text"` // HTTPステータスメッセージ

}
