package forecast_client

import (
	"Weather/pkg/forecast_client/models"
	"encoding/json"
	"net/url"
	"strconv"
)

func (c *Client) FindLocation(city string) ([]models.GeoLocation, error) {
	url, err := url.Parse(c.apiUrl + "/geo/1.0/direct")
	if err != nil {
		return nil, err
	}

	query := url.Query()
	query.Add("appid", c.apiToken)
	query.Add("q", city)
	query.Add("limit", "5")
	url.RawQuery = query.Encode()

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

	query := url.Query()
	query.Add("appid", c.apiToken)
	query.Add("lat", latStr)
	query.Add("lon", lonStr)
	url.RawQuery = query.Encode()

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
