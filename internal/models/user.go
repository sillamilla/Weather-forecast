package models

type User struct {
	ID    string `json:"id,omitempty"`
	Pined PinedCity
}

type PinedCity struct {
	Name   string `json:"name,omitempty"`
	Lat    string `json:"lat,omitempty"`
	Lon    string `json:"lon,omitempty"`
	Status int    `json:"status,omitempty"`
}
