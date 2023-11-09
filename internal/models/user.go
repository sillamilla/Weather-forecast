package models

type PinedCity struct {
	Name []string `json:"name,omitempty"`
}

type UserRequest struct {
	City        string
	StateCode   string
	CountryCode string
}
