package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/m1a9s9a4/language_tournament/model"
	"gorm.io/gorm"
)

func GetAnsweredUserById() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)
		uid := c.QueryParam("uid")
		users := new(model.AnsweredUsers)
		if err := users.GetQIdById(tx, uid); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return c.JSON(http.StatusOK, users)
	}
}

func GetAnsweredUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*gorm.DB)
		uid := c.QueryParam("uid")
		qid, err := strconv.ParseInt(c.QueryParam("qid"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		user := new(model.AnsweredUser)
		if err := user.FirstByMatch(tx, uid, qid); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func GetAnsweredUserByBattleId() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)
		uid := c.QueryParam("uid")
		bid, err := strconv.ParseInt(c.QueryParam("bid"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		users := new(model.AnsweredUsers)
		if err := users.GetQIdByBId(tx, uid, bid); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, users)
	}
}

func SaveAnsweredUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		user := new(model.AnsweredUser)
		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}

		tx := c.Get("Tx").(*gorm.DB)
		if err := user.Save(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}

		return c.JSON(http.StatusOK, "saved")
	}
}
