package model

import (
	"gorm.io/gorm"
)

type Answer struct {
	Id               int64 `json:"id"; gorm:"column:id"`
	BattleId         int64 `json:"battle_id"; gorm:"column:battle_id"`
	QuestionId       int64 `json:"question_id"; gorm:"column:question_id"`
	SelectedPlayerId int64 `json:"selected_player_id"; gorm:"column:selected_player_id"`
}

const AnswerTable = "answer"

func (a *Answer) Save(db *gorm.DB) error {
	rslt := db.Table(AnswerTable).Create(&a)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

type AnswerQuestions struct {
	Id               int64 `json:"id"; gorm:"column:id"`
	BattleId         int64 `json:"battle_id"; gorm:"column:battle_id"`
	QuestionId       int64 `json:"question_id"; gorm:"column:question_id"`
	SelectedPlayerId int64 `json:"selected_player_id"; gorm:"column:selected_player_id"`
	Question         `json:"question"`
}

type Answers []AnswerQuestions

func (as *Answers) GetByBId(db *gorm.DB, bid int64) error {
	rslt := db.
		Table(AnswerTable).
		Where("answer.battle_id = ?", bid).
		Joins("left join question on question.id = answer.question_id").
		Scan(&as)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

type CountAnswer struct {
	Target int64 `json:"target_id"`
	Count  int64 `json:"count"`
}

type CountAnswers []CountAnswer

func (as *CountAnswers) GetCountForEachQuestion(db *gorm.DB, bid int64) error {
	rslt := db.
		Table(AnswerTable).
		Select("question_id as target_id", "count(question_id) as count").
		Where("battle_id = ?", bid).
		Group("question_id").
		Scan(&as)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

func (as *CountAnswers) CountEachQuestionId(db *gorm.DB, bid int64, qid int64) error {
	rslt := db.
		Table(AnswerTable).
		Select("selected_player_id as target", "count(selected_player_id) as count").
		Where("battle_id = ?", bid).Where("question_id = ?", qid).
		Group("selected_player_id").
		Scan(&as)
	if rslt.Error != nil {
		return rslt.Error
	}
	return nil
}
