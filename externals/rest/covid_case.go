package externals

import (
	"covid-cases/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"
)

type CovidCases interface {
	GetInfo() (models.Cases, error)
}

type covidInfo struct {
	url string
}

func NewCovidInfo(config string) CovidCases {
	return covidInfo{url: config}
}

func (cfg covidInfo) GetInfo() (models.Cases, error) {

	// create http request
	resp, err := http.Get(cfg.url)
	if err != nil {
		return models.Cases{}, err
	}

	log.Infof("http client request: method[GET], path[%s]", resp.Request.URL.Path)

	// start request
	defer resp.Body.Close()

	// read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return models.Cases{}, err
	}
	// check http status
	if resp.StatusCode != http.StatusOK {
		return models.Cases{}, fmt.Errorf("%s return HTTP Status: %d", cfg.url, resp.StatusCode)
	}

	// unmarshal body to struct
	var response models.Cases
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.Cases{}, err
	}
	return response, nil
}
