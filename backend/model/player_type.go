package model

import "gorm.io/gorm"

type PlayerType struct {
	gorm.Model
	ID       int64  `json:"id"`
	Japanese string `json:"japanese"`
	English  string `json:"english"`
	ParentID int64  `json:"parent_id"`
}

const PlayerTypeTable = "player_type"

func (pt *PlayerType) FirstById(db *gorm.DB, ID int64) error {
	rslt := db.Table(PlayerTypeTable).Where("id = ?", ID).First(&pt)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}
