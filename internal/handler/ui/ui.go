package ui

import (
	"Weather/internal/models"
	"Weather/internal/service"
	models2 "Weather/pkg/forecast_client/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	srv  service.Service
	home *template.Template
}

func NewHandler(service service.Service) Handler {
	home, err := template.ParseFiles("./internal/templates/homePage.html")
	if err != nil {
		panic(err)
	}
	return Handler{
		srv:  service,
		home: home,
	}
}

func (h Handler) MainPageProcess(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("city")
	if cityName == "" {
		cityName = "London"
	}
	userRequest := models.UserRequest{
		City: cityName,
	}

	locations, err := h.srv.LocationSrv.FiendLocation(userRequest)
	if err != nil {
		log.Fatal(err)
	}

	var weathersInfo []models2.WeatherInfo
	var errors []error

	for _, location := range locations {
		weatherInfo, err := h.srv.WeatherSrv.WeatherStatus(location)
		if err != nil {
			errors = append(errors, err)
		}

		weathersInfo = append(weathersInfo, weatherInfo)
	}
	if len(errors) != 0 {
		fmt.Println(errors)
	}

	//pined, err := helper.GetPinedLocations(w, r)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//strPined := strings.Split(pined, ",")

	homeData := models.PageData{
		//PinedLocations: strPined,
		WeatherInfo: weathersInfo,
	}

	err = h.home.Execute(w, homeData)
	if err != nil {
		log.Println(err)
		return
	}
}

//
//func (h Handler) Pin(w http.ResponseWriter, r *http.Request) {
//	cityName := r.URL.Query().Get("city")
//	cookie, err := r.Cookie("location")
//	if err == nil {
//		cookie.Value = cookie.Value + "," + cityName
//	} else {
//		cookie = &http.Cookie{
//			Name:  "location",
//			Value: cityName,
//		}
//	}
//
//	http.SetCookie(w, cookie)
//	http.Redirect(w, r, "/", http.StatusSeeOther)
//}
//
//func (h Handler) Unpin(w http.ResponseWriter, r *http.Request) {
//	cityName := r.URL.Query().Get("city")
//	locations, err := helper.GetPinedLocations(w, r)
//	if err != nil {
//		return
//	}
//
//	if strings.Contains(locations, cityName) {
//		locations = strings.ReplaceAll(locations, cityName, "")
//		cookie := &http.Cookie{
//			Name:  "location",
//			Value: locations,
//		}
//
//		http.SetCookie(w, cookie)
//		http.Redirect(w, r, "/", http.StatusSeeOther)
//	}
//}
