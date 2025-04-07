package jsonModels

import (
	"time"
)

type ProxySettings struct {
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
	s.IgdbClient = "fill-in-pls"
	s.IgdbSecret = "fill-in-pls"
	s.IgdbAuth = "auto-generated"
	s.ExpiresIn = 0
	s.ExpiresAt = time.Time{}
}
