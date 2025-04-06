package main

import (
	"exhibtion-proxy/igdb"
	"exhibtion-proxy/jsonUtils"
	"exhibtion-proxy/jsonUtils/jsonModels"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {
	router := gin.Default()

	settings := &jsonModels.Settings{}
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

	router.GET("/game/id/:igdbid", ReturnJsonGameDataByID(apiManager))
	router.GET("/game/name/:name", ReturnJsonGameDataListByName(apiManager))

	err = router.Run(":3030")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReturnJsonGameDataListByName(apiManager *igdb.APIManager) gin.HandlerFunc {
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
func ReturnJsonGameDataByID(apiManager *igdb.APIManager) gin.HandlerFunc {
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
