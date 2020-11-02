package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/m1a9s9a4/language_tournament/model"
	"gorm.io/gorm"
)

func GetQuestionsByPlayerTypeId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		typeID, err := strconv.ParseInt(c.Param("type_id"), 0, 64)
		tx := c.Get("Tx").(*gorm.DB)
		questions := new(model.Questions)
		if err = questions.GetByPlayerTypeId(tx, typeID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "error occured")
		}

		return c.JSON(http.StatusOK, questions)
	}
}
