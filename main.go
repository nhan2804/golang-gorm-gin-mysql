package main

import (
	"gorm-test/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	videoRepo := controllers.NewVideo()
	r.GET("/video", videoRepo.GetVideo)
	r.POST("/video", videoRepo.CreateVideo)
	feedbackRepo := controllers.NewFeedBack()
	r.GET("/feedback", feedbackRepo.GetFeedback)

	return r
}
