package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

const PersistenceKey = "PERSISTENCE"

func Routing(e *echo.Echo) {
	http.Handle("/", e)
	HandleNotice(e)
}
