package exhibition_proxy_library

import (
	"fmt"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/handlers"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/igdb"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/jsonUtils"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/jsonUtils/jsonModels"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
)

type Proxy struct {
	SettingsPath string
	Settings *jsonModels.ProxySettings
	SettingsManger *jsonUtils.JsonManager
}

func (p *Proxy) Init() {
	settings := &jsonModels.ProxySettings{}
	settingsManager, err := jsonUtils.NewJsonManager(filepath.Join(p.SettingsPath), settings)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Settings = settings
	p.SettingsManger = settingsManager
}

func (p *Proxy) StartBaseServer() {
	gin.SetMode(gin.ReleaseMode)
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
