package model

import (
	"gorm.io/gorm"
)

type Battle struct {
	ID          int `json:"id"`
	Player_1_ID int `json:"player_1_id"`
	Player_2_ID int `json:"player_2_id"`
}

const BattleTable = "battle"

func (b *Battle) FirstByPlayerIds(db *gorm.DB, id1 int64, id2 int64) error {
	rslt := db.Table(BattleTable).
		Where(db.Where("player_1_id = ?", id1).Where("player_2_id = ?", id2)).
		Or(db.Where("player_1_id = ?", id2).Where("player_2_id = ?", id1)).
		First(&b)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

type Battles []Battle

func (bs *Battles) GetByPlayerId(db *gorm.DB, id int64) error {
	rslt := db.Table(BattleTable).Where("player_1_id = ?", id).Or(db.Where("player_2_id = ?", id)).Find(&bs)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}
