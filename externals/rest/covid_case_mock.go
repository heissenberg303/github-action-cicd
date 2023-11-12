package externals

import (
	"covid-cases/models"

	"github.com/stretchr/testify/mock"
)

type covidCasesMock struct {
	mock.Mock
}

func NewCovidCasesMock() *covidCasesMock {
	return &covidCasesMock{}
}

func (m *covidCasesMock) GetInfo() (models.Cases, error) {
	args := m.Called()
	return args.Get(0).(models.Cases), args.Error(1)
}
