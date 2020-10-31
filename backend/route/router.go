package route

import (
	"net/http"

	echoMw "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"github.com/m1a9s9a4/language_tournament/api"
	"github.com/m1a9s9a4/language_tournament/db"

	myMw "github.com/m1a9s9a4/language_tournament/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(myMw.TransactionHandler(db.Init()))

	v1 := e.Group("/api/v1")
	{
		v1.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "welcome to language tournament")
		})
		v1.GET("/language/:id", getLanguage)
		v1.GET("/battle/:l1/:l2", getBattle)
	}

	return e
}

func getLanguage(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getBattle(c echo.Context) error {
	l1 := c.Param("l1")
	l2 := c.Param("l2")
	// languageに紐づく質問を全て取得して渡す
	return c.String(http.StatusOK, "battle: "+l1+" vs "+l2)
}
