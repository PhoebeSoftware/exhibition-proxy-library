package jsonModels

import (
	"time"
)

type ProxySettings struct {
	IgdbSettings IgdbSettings
}

type IgdbSettings struct {
	IgdbClient string `json:"igdb_client"`
	IgdbSecret string `json:"igdb_secret"`
	IgdbAuth   string `json:"igdb_auth"`

	// In seconds
	ExpiresIn int `json:"expires_in"`

	// Basic go time format
	ExpiresAt time.Time `json:"expires_at"`
}

func (s ProxySettings) GetSettings() ProxySettings {
	return s
}

func (s *ProxySettings) DefaultValues() {
	s.IgdbSettings.IgdbClient = "fill-in-pls"
	s.IgdbSettings.IgdbSecret = "fill-in-pls"
	s.IgdbSettings.IgdbAuth = "auto-generated"
	s.IgdbSettings.ExpiresIn = 0
	s.IgdbSettings.ExpiresAt = time.Time{}
}
