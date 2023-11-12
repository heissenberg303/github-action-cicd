package handlers

import (
	"covid-cases/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type CovidHandler interface {
	Summary(c *gin.Context)
}

type summaryService struct {
	covidService services.CovidServices
}

func NewCovidHandler(covidService services.CovidServices) CovidHandler {
	return summaryService{covidService: covidService}
}

func (svc summaryService) Summary(c *gin.Context) {
	response, err := svc.covidService.Summary()
	if err != nil {
		log.Errorf("Failed to get response from Summary got err: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error!"})
		return
	}
	c.JSON(http.StatusOK, response)
}
