package model

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	ID       int    `gorm:"primary_key" json:"id"`
	Japanese string `json:"japanese"`
	English  string `json:"english"`
	Img      string `json:"img"`
	TypeID   *int   `json:"type_id"`
}

const PlayerTable = "player"

func (p *Player) FirstById(db *gorm.DB, id int64) error {
	rslt := db.Table(PlayerTable).Where("id = ?", id).First(&p)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

func (p *Player) GetByEnglish(db *gorm.DB, english string) error {
	rslt := db.Table(PlayerTable).Where("english = ?", english).First(&p)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

type Players []Player

func (ps *Players) GetByTypeId(db *gorm.DB, typeID int64) error {
	rslt := db.Table(PlayerTable).Where("type_id = ?", typeID).Find(&ps)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}
