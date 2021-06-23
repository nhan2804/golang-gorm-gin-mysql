package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Avatar string `json:"avatar"`
	Video  string `json:"video"`
}

func GetFeedback(db *gorm.DB, Feedback *[]Feedback) (err error) {
	err = db.Find(Feedback).Error
	if err != nil {
		return err
	}
	return nil
}
