package controllers

import (
	"gorm-test/database"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoVideo struct {
	Db *gorm.DB
}

func NewVideo() *RepoVideo {
	db := database.InitDb()
	db.AutoMigrate(&models.Video{})
	return &RepoVideo{Db: db}
}
func (v *RepoVideo) CreateVideo(g *gin.Context) {
	var video models.Video
	g.BindJSON(&video)
	err := models.CreateVideo(v.Db, &video)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Có lỗi xảy ra"})
		return
	}
	g.JSON(http.StatusOK, video)

}
func (v *RepoVideo) GetVideo(g *gin.Context) {
	var videos []models.Video

	err := models.GetVideos(v.Db, &videos)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Có lỗi xảy ra"})
		return
	}
	g.JSON(http.StatusOK, videos)

}
