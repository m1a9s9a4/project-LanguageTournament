package middleware

import (
	"log"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			tx := db.Begin()
			log.Println(tx)
			c.Set(TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debug("Transaction Rollback ", err)
				return err
			}
			logrus.Debug("Transaction Commit")
			tx.Commit()

			return nil
		})
	}
}
