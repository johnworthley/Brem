package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)

	go RunUpdater()

	router = gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("brem", store))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(corsConfig))

	initAPI()

	router.Run("127.0.0.1:8080")
}

func initAPI() {
	router.POST("/signup", addDeveloper)
	router.GET("/login", login)

	superuserGroup := router.Group("/super")
	superuserGroup.Use(SuperuserAuth())
	{
		superuserGroup.GET("/audit", getAllAuditors)
	}

	developerGroup := router.Group("/dev")
	developerGroup.Use(DeveloperAuth())
	{
		developerGroup.GET("/ico/dev", getDevelopersICOs)
	}

	//router.POST("/audit", addAuditor)
	router.GET("/audit/ico", getAuditorICOs)
	router.GET("/ico/audit", getICOAuditors)
	router.GET("/ico/created", getCreatedICOs)
	router.GET("/ico/opened", getOpennedICOs)
	router.GET("/ico/success", getSuccessICOs)
	router.GET("/ico/failed", getFailedICOs)
	router.GET("/ico/withdrawn", getWithdrawnICOs)
	router.POST("/ico", addICO)
	router.GET("/ico/image", getICOImage)
	router.POST("/ico/image", addICOImage)
	router.POST("/ico/audit", addAuditorToICO)
	router.PUT("/ico/open", publishICO)
	router.PUT("/ico/success", setICOSucccessStatus)
	router.PUT("/ico/request", setICORequestedStatus)
	router.PUT("/ico/withdrawn", setICOWithdrawnStatus)
}
