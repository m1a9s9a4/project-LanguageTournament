package model

import "gorm.io/gorm"

const QuestionTable = "question"

type Question struct {
	gorm.Model
	ID           int64  `json:"id"`
	Japanese     string `json:"japanese"`
	English      string `json:"english"`
	PlayerTypeID int64  `json:"player_type_id"`
}

type Questions []Question

func (qs *Questions) GetByPlayerTypeId(db *gorm.DB, id int64) error {
	rslt := db.Table(QuestionTable).Where("player_type_id = ?", id).Find(&qs)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}
