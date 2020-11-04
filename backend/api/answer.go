package api

import (
	"net/http"
	"strconv"

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

		if err := answer.Save(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "error occured")
		}

		return c.JSON(http.StatusOK, "saved")
	}
}

func GetAnswersByBattleId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		bid, err := strconv.ParseInt(c.Param("bid"), 0, 64)
		if err != nil {
			return err
		}
		tx := c.Get("Tx").(*gorm.DB)
		as := new(model.Answers)
		if err := as.GetByBId(tx, bid); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, as)
	}
}

func GetCountEachQuestions() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		bid, err := strconv.ParseInt(c.Param("bid"), 0, 64)
		if err != nil {
			return err
		}
		tx := c.Get("Tx").(*gorm.DB)
		as := new(model.CountAnswers)
		if err := as.GetCountForEachQuestion(tx, bid); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, as)
	}
}

func CountEachQuestionId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		bid, err := strconv.ParseInt(c.Param("bid"), 0, 64)
		if err != nil {
			return err
		}
		qid, err := strconv.ParseInt(c.Param("qid"), 0, 64)
		if err != nil {
			return err
		}
		tx := c.Get("Tx").(*gorm.DB)
		as := new(model.CountAnswers)
		if err := as.CountEachQuestionId(tx, bid, qid); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, as)
	}
}
