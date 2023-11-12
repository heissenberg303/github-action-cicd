package routers

import (
	"covid-cases/handlers"
	"covid-cases/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(covidService services.CovidServices) *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "helloworld.html", gin.H{
			"title": "Hello, World!",
		})
	})
	r.GET("/health", healthCheck)
	initCovidRouter(r, covidService)

	return r
}

func initCovidRouter(r *gin.Engine, covidService services.CovidServices) {
	handler := handlers.NewCovidHandler(covidService)

	covid := r.Group("covid")
	covid.GET("/summary", handler.Summary)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running~",
	})
}
