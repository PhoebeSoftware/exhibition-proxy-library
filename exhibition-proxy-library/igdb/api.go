package igdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/jsonUtils"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/jsonUtils/jsonModels"
	"net/http"
	"time"
)

type Image struct {
	ImageID string `json:"image_id" gorm:"primaryKey"`
}

type Genre struct {
	GenreID int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
}

type Metadata struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"summary"`
	// Coverid is foreign key for local db
	CoverID     string    `json:"-"`
	Cover       Image   `json:"cover" gorm:"foreignKey:CoverID"`
	Artworks    []Image `json:"artworks" gorm:"many2many:metadata_artworks"`
	Screenshots []Image `json:"screenshots" gorm:"many2many:metadata_screenshots"`
	Genres      []Genre `json:"genres" gorm:"many2many:metadata_genres"`
}

type APIManager struct {
	client   *http.Client
	settings *jsonModels.ProxySettings
}

func (a *APIManager) SetupHeader(request *http.Request) {
	request.Header.Set("Client-ID", a.settings.IgdbSettings.IgdbClient)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.settings.IgdbSettings.IgdbAuth))
}

func NewAPI(settings *jsonModels.ProxySettings, settingsManager *jsonUtils.JsonManager) (*APIManager, error) {
	apiManager := &APIManager{
		client:   &http.Client{},
		settings: settings,
	}

	// Generate new auth token if needed
	if time.Now().After(settings.IgdbSettings.ExpiresAt) {
		fmt.Println("Generating new auth token because the old one has expired")
		_, err := apiManager.GetAndSetNewAuthToken()
		if err != nil {
			fmt.Println("error fetching acces token:", err)
			return nil, err
		}
		// Save new auth token
		err = settingsManager.Save()
		if err != nil {
			return nil, err
		}
	}

	return apiManager, nil
}

func (a *APIManager) GetGameData(id int) (*Metadata, error) {
	header := fmt.Sprintf(`fields id, name, summary, cover.*, artworks.*, screenshots.*, genres.*; where id = %d;`, id)

	request, err := http.NewRequest("POST", "https://api.igdb.com/v4/games/", bytes.NewBuffer([]byte(header)))
	if err != nil {
		return &Metadata{}, err
	}

	a.SetupHeader(request)

	response, err := a.client.Do(request)
	if err != nil {
		return &Metadata{}, err
	}
	defer response.Body.Close()

	var gameDataList []Metadata
	jsonErr := json.NewDecoder(response.Body).Decode(&gameDataList)
	if jsonErr != nil {
		return &Metadata{}, err
	}

	if len(gameDataList) == 0 {
		return &Metadata{}, fmt.Errorf("no games found with id %d", id)
	}

	firstGameData := &gameDataList[0]
	return firstGameData, nil
}

func (a *APIManager) GetGames(query string) ([]Metadata, error) {
	header := fmt.Sprintf(`fields id, name, summary, cover.*, artworks.*, screenshots.*, genres.*; search "%s";`, query)

	request, err := http.NewRequest("POST", "https://api.igdb.com/v4/games/", bytes.NewBuffer([]byte(header)))
	if err != nil {
		return []Metadata{}, err
	}

	a.SetupHeader(request)

	response, err := a.client.Do(request)
	if err != nil {
		return []Metadata{}, err
	}
	defer response.Body.Close()

	var games []Metadata
	err = json.NewDecoder(response.Body).Decode(&games)
	if err != nil {
		return []Metadata{}, err
	}

	return games, nil
}
