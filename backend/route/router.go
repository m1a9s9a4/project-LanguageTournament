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
	e.Use(echoMw.CORS())
	e.Use(myMw.TransactionHandler(db.Init()))

	v1 := e.Group("/api/v1")
	{
		v1.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "welcome to language tournament")
		})
		v1.GET("/player/:id", api.FirstPlayerById())
		v1.GET("/players/type/:type_id", api.GetPlayersByTypeId())
		v1.GET("/player/english/:english", api.FirstPlayerByEnglish())
		v1.GET("/player_type/:id", api.FirstPlayerTypeById())
		v1.GET("/questions/type/:type_id", api.GetQuestionsByPlayerTypeId())
		v1.GET("/questions/type/:type_id", api.GetQuestionsByPlayerTypeId())
		v1.GET("/battle/:id1/:id2", api.GetBattleByIds())
		v1.GET("/battle/player/:id", api.GetBattleById())
		v1.GET("/answer/user", api.GetAnsweredUser())
		v1.GET("/answer/users", api.GetAnsweredUserById())

		v1.POST("/answer", api.SaveAnswer())
		v1.POST("/answer/user", api.SaveAnsweredUser())
	}

	return e
}
