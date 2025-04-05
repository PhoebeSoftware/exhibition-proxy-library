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

	apiManager, err := igdb.NewAPI(settings, settingsManager)
	if err != nil {
		fmt.Println(err)
		return
	}
	router.GET("/getGame/:igdbid", returnGameData(apiManager))
	err = router.Run(":3030")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func returnGameData(apiManager *igdb.APIManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Param("igdbid")
		id, _ := strconv.Atoi(idString)
		gameData, err := apiManager.GetGameData(id)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, gameData)
	}
}
