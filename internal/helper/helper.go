package helper

import (
	models2 "Weather/pkg/forecast_client/models"
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

//func GetPinedLocations(w http.ResponseWriter, r *http.Request) (string, error) {
//	cookie, err := r.Cookie("location")
//	if err != nil {
//		if errors.Is(err, http.ErrNoCookie) {
//			return "", nil
//		} else {
//			return "", err
//		}
//	}
//
//	return cookie.Value, err
//}
