package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model

	Title     string `json:"title"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
	Author    string `json:"author"`
}

func CreateVideo(db *gorm.DB, Video *Video) (err error) {
	err = db.Create(Video).Error
	if err != nil {
		return err
	}
	return nil
}
func GetVideos(db *gorm.DB, Video *[]Video) (err error) {
	err = db.Find(Video).Error
	if err != nil {
		return err
	}
	return nil
}
