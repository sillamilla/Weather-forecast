package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GetConfig()
}

var (
	c *Config
)

type Config struct {
	WeatherApiToken string
	WeatherApiUrl   string
}

func GetConfig() *Config {
	if c == nil {
		weatherApiToken := os.Getenv("WEATHER_API_TOKEN")
		if weatherApiToken == "" {
			panic("WEATHER_API_TOKEN is not set")
		}

		weatherApiUrl := os.Getenv("WEATHER_API_URL")
		if weatherApiUrl == "" {
			panic("WEATHER_API_URL is not set")
		}

		c = &Config{
			WeatherApiToken: weatherApiToken,
			WeatherApiUrl:   weatherApiUrl,
		}

		return c
	}

	return c
}
