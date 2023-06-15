package helper

import (
	models2 "Weather/pkg/forecast_client/models"
	"github.com/google/uuid"
	"net/http"
)

func SameCheck(locationsInfo []models2.GeoLocation) []models2.GeoLocation {

	sortMap := make(map[string]models2.GeoLocation)
	for i, value := range locationsInfo {
		sortMap[value.Name] = locationsInfo[i]
	}

	var location []models2.GeoLocation
	for _, value := range sortMap {
		location = append(location, sortMap[value.Name])
	}

	return location
}

func GenerateAndSetCookie(w http.ResponseWriter) string {
	cookieValue := uuid.New().String()

	cookie := http.Cookie{
		Name:  "user",
		Value: cookieValue,
	}

	http.SetCookie(w, &cookie)
	return cookie.Value
}

func GetCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("user")
	if cookie == nil {
		return "", err
	}

	return cookie.Value, nil
}
