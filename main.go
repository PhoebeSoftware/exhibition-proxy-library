package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/test", test)
	err := router.Run(":3030")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"wassupwigga": "yo",
	})
}
