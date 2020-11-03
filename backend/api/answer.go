package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/m1a9s9a4/language_tournament/model"
	"gorm.io/gorm"
)

func SaveAnswer() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		answer := new(model.Answer)
		tx := c.Get("Tx").(*gorm.DB)
		if err := c.Bind(answer); err != nil {
			return err
		}

		if err := answer.Save(tx, answer); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "error occured")
		}

		return c.JSON(http.StatusOK, "saved")
	}
}
