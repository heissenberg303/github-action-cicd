package services

import (
	externals "covid-cases/externals/rest"
	"covid-cases/models"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Summary(t *testing.T) {

	tests := []struct {
		name        string
		args        []models.Data
		expected    models.SummaryResponse
		expectedErr bool
	}{
		{
			name: "1- Happy",
			args: []models.Data{
				{
					ConfirmDate:    "2021-05-04",
					No:             0,
					Age:            51,
					Gender:         "หญิง",
					GenderEn:       "Female",
					Nation:         "",
					NationEn:       "China",
					Province:       "Phrae",
					ProvinceId:     46,
					District:       "",
					ProvinceEn:     "Phrae",
					StatQuarantine: 5,
				},
				{
					ConfirmDate:    "2021-05-04",
					No:             0,
					Age:            52,
					Gender:         "หญิง",
					GenderEn:       "Female",
					Nation:         "",
					NationEn:       "Thailand",
					Province:       "Chumphon",
					ProvinceId:     12,
					District:       "",
					ProvinceEn:     "Chumphon",
					StatQuarantine: 8,
				},
				{
					ConfirmDate:    "2021-05-04",
					No:             0,
					Age:            99,
					Gender:         "หญิง",
					GenderEn:       "Female",
					Nation:         "",
					NationEn:       "China",
					Province:       "Bangkok",
					ProvinceId:     0,
					District:       "",
					ProvinceEn:     "Bangkok",
					StatQuarantine: 18,
				},
				{
					ConfirmDate:    "2021-05-01",
					No:             0,
					Age:            25,
					Gender:         "",
					GenderEn:       "",
					Nation:         "",
					NationEn:       "India",
					Province:       "Phrae",
					ProvinceId:     46,
					District:       "",
					ProvinceEn:     "Phrae",
					StatQuarantine: 15,
				},
			},
			expected: models.SummaryResponse{
				Province: map[string]int{
					"Bangkok":  1,
					"Chumphon": 1,
					"Phrae":    2,
				},
				AgeGroup: map[string]int{
					"0-30":  1,
					"31-60": 2,
					"61+":   1,
				},
			},
			expectedErr: false,
		},
		{
			name: "2- No Province Data",
			args: []models.Data{
				{
					ConfirmDate:    "2021-05-04",
					No:             0,
					Age:            51,
					Gender:         "หญิง",
					GenderEn:       "Female",
					Nation:         "",
					NationEn:       "China",
					Province:       "",
					ProvinceId:     46,
					District:       "",
					ProvinceEn:     "",
					StatQuarantine: 5,
				},
				{
					ConfirmDate:    "2021-05-01",
					No:             0,
					Age:            25,
					Gender:         "",
					GenderEn:       "",
					Nation:         "",
					NationEn:       "India",
					ProvinceId:     46,
					District:       "",
					ProvinceEn:     "",
					StatQuarantine: 15,
				},
				{
					ConfirmDate:    "2021-05-02",
					No:             0,
					Age:            39,
					Gender:         "",
					GenderEn:       "",
					Nation:         "",
					NationEn:       "USA",
					Province:       "",
					ProvinceId:     0,
					District:       "",
					ProvinceEn:     "",
					StatQuarantine: 10,
				},
			},
			expected: models.SummaryResponse{
				Province: map[string]int{
					"N/A": 3,
				},
				AgeGroup: map[string]int{
					"0-30":  1,
					"31-60": 2,
				},
			},
			expectedErr: false,
		},
		{
			name: "3- No Age Data",
			args: []models.Data{
				{
					ConfirmDate:    "2021-05-04",
					No:             0,
					Gender:         "หญิง",
					GenderEn:       "Female",
					Nation:         "",
					NationEn:       "China",
					Province:       "",
					ProvinceId:     46,
					District:       "",
					ProvinceEn:     "",
					StatQuarantine: 5,
				},
				{
					ConfirmDate:    "2021-05-01",
					No:             0,
					Gender:         "",
					GenderEn:       "",
					Nation:         "",
					NationEn:       "India",
					Province:       "",
					ProvinceId:     46,
					District:       "",
					ProvinceEn:     "",
					StatQuarantine: 15,
				},
				{
					ConfirmDate:    "2021-05-02",
					No:             0,
					Gender:         "",
					GenderEn:       "",
					Nation:         "",
					NationEn:       "USA",
					Province:       "",
					ProvinceId:     0,
					District:       "",
					ProvinceEn:     "",
					StatQuarantine: 10,
				},
			},
			expected: models.SummaryResponse{
				Province: map[string]int{
					"N/A": 3,
				},
				AgeGroup: map[string]int{
					"N/A": 3,
				},
			},
			expectedErr: false,
		},
		{
			name: "4- No Data",
			args: []models.Data{},
			expected: models.SummaryResponse{
				Province: map[string]int{},
				AgeGroup: map[string]int{},
			},
			expectedErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock
			covidAPI := externals.NewCovidCasesMock()
			covidAPI.On("GetInfo").Return(models.Cases{Data: tt.args}, nil)

			services := NewSummaryService(covidAPI)
			got, err := services.Summary()
			if (err != nil) != tt.expectedErr {
				t.Errorf("covidAPI.Summary() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("covidAPI.Summary() = %v, want %v", got, tt.expected)
			}
		})
	}

	t.Run("Error from http request", func(t *testing.T) {
		covidAPI := externals.NewCovidCasesMock()
		covidAPI.On("GetInfo").Return(models.Cases{}, errors.New(""))

		services := NewSummaryService(covidAPI)
		_, err := services.Summary()
		assert.ErrorIs(t, err, ErrorExternalAPI)
	})
}
