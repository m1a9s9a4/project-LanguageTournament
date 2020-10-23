package route

import (
	"language_tournament/api"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Debug()

	v1 := e.Group("/api/v1")
	{
		v1.POST("/languages", api.GetLanguages())
	}
}
