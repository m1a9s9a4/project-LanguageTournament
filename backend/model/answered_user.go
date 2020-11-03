package model

import "gorm.io/gorm"

type AnsweredUser struct {
	gorm.Model
	UID        string `json:"uid"`
	QuestionID int    `json:"question_id"`
}

const AnsweredUserTable = "answered_users"

func (au *AnsweredUser) FirstByMatch(db *gorm.DB, uid string, qid int64) error {
	if rslt := db.
		Table(AnsweredUserTable).
		Where("uid = ?", uid).
		Where("question_id = ?", qid).
		First(&au); rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

func (au *AnsweredUser) Save(db *gorm.DB) error {
	if rslt := db.Table(AnsweredUserTable).Create(&au); rslt.Error != nil {
		return rslt.Error
	}
	return nil
}

type AnsweredUserQuestionId struct {
	QuestionID int `json:"question_id"`
}

type AnsweredUsers []AnsweredUserQuestionId

func (aus *AnsweredUsers) GetQIdById(db *gorm.DB, uid string) error {
	if rslt := db.Table(AnsweredUserTable).Select("question_id").Where("uid = ?", uid).Find(&aus); rslt.Error != nil {
		return rslt.Error
	}
	return nil
}
