package services

import (
	externals "covid-cases/externals/rest"
	"covid-cases/models"
	"errors"

	"github.com/labstack/gommon/log"
)

var (
	ErrorExternalAPI = errors.New("external API error.")
)

type CovidServices interface {
	Summary() (models.SummaryResponse, error)
}

type covidAPI struct {
	covid externals.CovidCases
}

func NewSummaryService(covid externals.CovidCases) CovidServices {
	return covidAPI{covid: covid}
}

func (ext covidAPI) Summary() (models.SummaryResponse, error) {
	cases, err := ext.covid.GetInfo()
	if err != nil {
		log.Errorf("Failed to get covid information from external API got error: %s", err.Error())
		return models.SummaryResponse{}, ErrorExternalAPI
	}

	response := groupByProvinceAndAge(&cases)

	return response, nil
}

func groupByProvinceAndAge(cases *models.Cases) models.SummaryResponse {

	provinces, ageGroup := make(map[string]int), make(map[string]int)

	for _, data := range cases.Data {

		// Provice Group
		if data.Province == "" {
			data.Province = "N/A"
		}
		provinces[data.Province]++

		// Age Group
		switch {
		case data.Age > 0 && data.Age <= 30:
			ageGroup["0-30"]++
		case data.Age > 30 && data.Age <= 60:
			ageGroup["31-60"]++
		case data.Age >= 61:
			ageGroup["61+"]++
		default:
			ageGroup["N/A"]++
		}
	}

	return models.SummaryResponse{
		Province: provinces,
		AgeGroup: ageGroup,
	}
}
