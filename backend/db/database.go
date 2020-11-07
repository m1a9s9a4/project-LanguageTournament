package db

import (
	"context"
	"time"

	"github.com/m1a9s9a4/language_tournament/conf"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Session struct {
	DryRun            bool
	PrepareStmt       bool
	WithConditions    bool
	AllowGlobalUpdate bool
	Context           context.Context
	Logger            logger.Interface
	NowFunc           func() time.Time
}

type Player struct {
	gorm.Model
	ID       int    `gorm:"primary_key" json:"id"`
	Japanese string `json:"japanese"`
	English  string `json:"english"`
	Img      string `json:"img"`
	TypeID   int    `json:"type_id"`
	// Type     PlayerType
}

func Init() *gorm.DB {
	session := createSession()

	return session
}

func createSession() *gorm.DB {
	dsn := _createDsn()
	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}

	return db
}

func _createDsn() string {
	return "host=" + conf.HOST + " user=" + conf.USER + " password=" + conf.PASS + " dbname=" + conf.DB + " port=" + conf.PORT + " sslmode=disable TimeZone=Asia/Tokyo"
}
