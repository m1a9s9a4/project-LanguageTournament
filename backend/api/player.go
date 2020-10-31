package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/m1a9s9a4/language_tournament/model"
	"gorm.io/gorm"
)

func GetPlayersByTypeId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		typeID, err := strconv.ParseInt(c.Param("type_id"), 0, 64)
		tx := c.Get("Tx").(*gorm.DB)

		players := new(model.Players)
		if err = players.GetByTypeId(tx, typeID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "error occured")
		}

		return c.JSON(http.StatusOK, players)
	}
}

func FirstPlayerByEnglish() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		english := c.Param("english")
		log.Println("english")
		log.Println(english)
		tx := c.Get("Tx").(*gorm.DB)
		player := new(model.Player)
		if err = player.GetByEnglish(tx, english); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "error occured")
		}

		return c.JSON(http.StatusOK, player)
	}
}
