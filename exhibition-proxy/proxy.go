package exhibition_proxy

import (
	"fmt"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy/igdb"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy/jsonUtils"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy/jsonUtils/jsonModels"
	"github.com/PhoebeSoftware/exhibition-proxy-library/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type Proxy struct {
	SettingsPath string
	Port int
}

func (p *Proxy) StartServer() {
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

	router.GET("/game/:id", returnJsonGameDataByID(apiManager))
	router.GET("/game/", returnJsonGameDataListByName(apiManager))

	err = router.Run(":" + portInString)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func returnJsonGameDataListByName(apiManager *igdb.APIManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")
		if name == "" {
			ctx.JSON(http.StatusBadRequest, models.Error{
				ErrorMessage: "No search query",
				StatusCode:   http.StatusBadRequest,
			})
			return
		}
		gameDataList, err := apiManager.GetGames(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gameDataList)
	}
}
func returnJsonGameDataByID(apiManager *igdb.APIManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Param("igdbid")
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println(err)
			return
		}
		gameData, err := apiManager.GetGameData(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gameData)
	}
}
