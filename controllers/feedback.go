package controllers

import (
	"gorm-test/database"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoFeedBack struct {
	Db *gorm.DB
}

func NewFeedBack() *RepoVideo {
	db := database.InitDb()
	// db.AutoMigrate(&models.Video{})
	return &RepoVideo{Db: db}
}

func (v *RepoVideo) GetFeedback(g *gin.Context) {
	var fbs []models.Feedback

	err := models.GetFeedback(v.Db, &fbs)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Có lỗi xảy ra"})
		return
	}
	g.JSON(http.StatusOK, fbs)

}
