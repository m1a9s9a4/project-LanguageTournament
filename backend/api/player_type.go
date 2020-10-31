package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/m1a9s9a4/language_tournament/model"
	"gorm.io/gorm"
)

func FirstPlayerTypeById() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ID, err := strconv.ParseInt(c.Param("id"), 0, 64)
		tx := c.Get("Tx").(*gorm.DB)

		pt := new(model.PlayerType)
		if err = pt.FirstById(tx, ID); err != nil {
			// change error message to each situaction
			return echo.NewHTTPError(http.StatusNotFound, "error occured")
		}

		return c.JSON(http.StatusOK, pt)
	}
}
