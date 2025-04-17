package exhibition_proxy_library

import (
	"fmt"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/handlers"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/igdb"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/jsonUtils"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/jsonUtils/jsonModels"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
)

type Proxy struct {
	Settings       *jsonModels.ProxySettings
	SettingsManger *jsonUtils.JsonManager
	DataPath       string
}

func (p *Proxy) Init() {
	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = filepath.Join("..", "data")
	}
	if err := os.MkdirAll(dataPath, 0777); err != nil {
		fmt.Println(err)
		fmt.Println("proxy.Init(): Could not create path: " + dataPath)
		return
	}
	p.DataPath = dataPath
	settingsPath := filepath.Join(dataPath, "proxy-settings.json")

	settings := &jsonModels.ProxySettings{}
	settingsManager, err := jsonUtils.NewJsonManager(settingsPath, settings)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Settings = settings
	p.SettingsManger = settingsManager
	if settings.IgdbSettings.IgdbClient == "fill-in-pls" ||
		settings.IgdbSettings.IgdbSecret == "fill-in-pls" {
		absPath, err := filepath.Abs(settingsPath)
		if err != nil {
			fmt.Println("Config file path: " + settingsPath)
		} else {
			fmt.Println("Config file path:", absPath)
		}
		panic("Failed to launch: Please fill in the IGDB client and secret")
	}
}

func (p *Proxy) StartBaseServer() {
	if !p.Settings.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	settings := p.Settings
	settingsManager := p.SettingsManger

	portInString := strconv.Itoa(settings.Port)
	router := gin.Default()

	apiManager, err := igdb.NewAPI(settings, settingsManager)
	if err != nil {
		fmt.Println(err)
		return
	}

	router.GET("/game/:igdbid", handlers.HandleSearchByID(apiManager))
	router.GET("/game/", handlers.HandleSearchByName(apiManager))

	err = router.Run(":" + portInString)
	if err != nil {
		fmt.Println(err)
		return
	}
}
