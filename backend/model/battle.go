package model

import (
	"gorm.io/gorm"
)

type Battle struct {
	ID        int `json:"id"`
	Player1ID int `json:"player_1_id"`
	Player2ID int `json:"player_2_id"`
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
