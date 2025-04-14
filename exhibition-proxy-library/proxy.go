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
	Port         int
}

func (p *Proxy) StartBaseServer() {
	portInString := strconv.Itoa(p.Port)
	router := gin.Default()

	settings := &jsonModels.ProxySettings{}
	settingsManager, err := jsonUtils.NewJsonManager(filepath.Join(p.SettingsPath), settings)
	if err != nil {
		fmt.Println(err)
		return
	}

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
