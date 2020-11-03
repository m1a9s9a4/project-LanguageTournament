package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	ID               int64 `json:"id"`
	BattleID         int64 `json:"battle_id"`
	QuestionID       int64 `json:"question_id"`
	SelectedPlayerID int64 `json:"selected_player_id"`
}

const AnswerTable = "answer"

func (a *Answer) Save(db *gorm.DB, answer *Answer) error {
	rslt := db.Table(AnswerTable).Create(answer)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}
