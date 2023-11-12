package main

import (
	externals "covid-cases/externals/rest"
	"covid-cases/models"

	"covid-cases/routers"
	"covid-cases/services"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {


	cfg := initConfig()

	covid := externals.NewCovidInfo(cfg.Url)
	svc := services.NewSummaryService(covid)

	g := routers.InitRouter(svc)

	err := g.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is starting")
}

func initConfig() models.Config {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.AddConfigPath(".")      // Path to look for the config file in
	viper.ReadInConfig()          // Read the config file

	url := viper.GetString("external.covid.url")
	port := viper.GetInt("app.port")

	if url == "" || port == 0 {
		log.Fatal("Failed to read config.")
	}

	config := models.Config{
		Url:  url,
		Port: port,
	}
	return config
}
