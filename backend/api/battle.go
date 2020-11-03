package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/m1a9s9a4/language_tournament/model"
	"gorm.io/gorm"
)

func GetBattleByIds() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id1, err1 := strconv.ParseInt(c.Param("id1"), 0, 64)
		if err1 != nil {
			return err1
		}
		id2, err2 := strconv.ParseInt(c.Param("id2"), 0, 64)
		if err2 != nil {
			return err2
		}
		tx := c.Get("Tx").(*gorm.DB)

		battle := new(model.Battle)
		if err = battle.FirstByPlayerIds(tx, id1, id2); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		return c.JSON(http.StatusOK, battle)
	}
}
