package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	//go RunUpdater()

	// TODO: File server

	router = gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(corsConfig))

	initAPI()

	router.Run("127.0.0.1:8080")
}

func initAPI() {
	router.POST("/dev", addDeveloper)
	router.GET("/audit", getAllAuditors)
	router.POST("/audit", addAuditor)
	router.GET("/ico/dev", getDevelopersICOs)
	router.GET("/audit/ico", getAuditorICOs)
	router.GET("/ico/audit", getICOAuditors)
	router.GET("/ico/created", getCreatedICOs)
	router.GET("/ico/opened", getOpennedICOs)
	router.GET("/ico/success", getSuccessICOs)
	router.GET("/ico/failed", getFailedICOs)
	router.GET("/ico/withdrawn", getWithdrawnICOs)
	router.POST("/ico", addICO)
	router.POST("/ico/audit", addAuditorToICO)
	router.PUT("/ico/open", publishICO)
	router.PUT("/ico/request", setICORequestedStatus)
	router.PUT("/ico/withdrawn", setICOWithdrawnStatus)
}
