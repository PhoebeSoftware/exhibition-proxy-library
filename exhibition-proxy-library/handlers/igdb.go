package handlers

import (
	"fmt"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/igdb"
	"github.com/PhoebeSoftware/exhibition-proxy-library/exhibition-proxy-library/proxy_models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleSearchByName(apiManager *igdb.APIManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")
		if name == "" {
			ctx.JSON(http.StatusBadRequest, proxy_models.Error{
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
func HandleSearchByID(apiManager *igdb.APIManager) gin.HandlerFunc {
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
