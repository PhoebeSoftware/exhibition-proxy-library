package proxy

import (
	"fmt"
	"github.com/PhoebeSoftware/exhibition-proxy-library/igdb"
	"github.com/PhoebeSoftware/exhibition-proxy-library/jsonUtils"
	"github.com/PhoebeSoftware/exhibition-proxy-library/jsonUtils/jsonModels"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type Proxy struct {

}

func (p *Proxy) StartServer() {
	router := gin.Default()

	settings := &jsonModels.ProxySettings{}
	settingsManager, err := jsonUtils.NewJsonManager(filepath.Join("settings.json"), settings)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("test")

	apiManager, err := igdb.NewAPI(settings, settingsManager)
	if err != nil {
		fmt.Println(err)
		return
	}

	router.GET("/game/id/:igdbid", returnJsonGameDataByID(apiManager))
	router.GET("/game/name/:name", returnJsonGameDataListByName(apiManager))

	err = router.Run(":3030")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func returnJsonGameDataListByName(apiManager *igdb.APIManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
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
