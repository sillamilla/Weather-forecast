package weather_client

import (
	"Weather/pkg/weather_client/models"
	"encoding/json"
	"net/url"
	"strconv"
)

func (c *Client) FindLocation(city string) ([]models.GeoLocation, error) {
	url, err := url.Parse(c.apiUrl + "/geo/1.0/direct")
	if err != nil {
		return nil, err
	}

	url.Query().Add("appid", c.apiToken)
	url.Query().Add("q", city)
	url.Query().Add("limit", "5")
	url.RawQuery = url.Query().Encode()

	body, err := c.get(url.String())
	if err != nil {
		return []models.GeoLocation{}, err
	}

	var locations []models.GeoLocation
	if err = json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func (c *Client) WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error) {
	url, err := url.Parse(c.apiUrl + "/data/2.5/weather")
	if err != nil {
		return models.WeatherInfo{}, err
	}

	latStr := strconv.FormatFloat(geoLocation.Lat, 'f', 6, 64)
	lonStr := strconv.FormatFloat(geoLocation.Lon, 'f', 6, 64)

	url.Query().Add("appid", c.apiToken)
	url.Query().Add("lat", latStr)
	url.Query().Add("lon", lonStr)
	url.RawQuery = url.Query().Encode()

	body, err := c.get(url.String())
	if err != nil {
		return models.WeatherInfo{}, err
	}

	var weatherInfo models.WeatherInfo
	if err = json.Unmarshal(body, &weatherInfo); err != nil {
		return models.WeatherInfo{}, err
	}

	return weatherInfo, nil
}
