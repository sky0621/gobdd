package handler

import (
	"errors"
	"fmt"
	"gobdd/adapter/middleware/persistence"
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
		return errors.New("abnormal end")
	}

	fmt.Println(rdbMiddleware)

	// FIXME:
	return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}
