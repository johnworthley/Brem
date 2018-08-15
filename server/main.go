package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
	router.POST("/login", login)

	superuserGroup := router.Group("/super")
	superuserGroup.Use(SuperuserAuth())
	{
		superuserGroup.GET("/audit", getAllAuditors)
		superuserGroup.POST("/ico/audit", addAuditorToICO)
		superuserGroup.PUT("/ico/open", publishICO)
	}

	developerGroup := router.Group("/dev")
	developerGroup.Use(DeveloperAuth())
	{
		developerGroup.GET("/ico", getDevelopersICOs)
		developerGroup.POST("/ico", addICO)
		developerGroup.POST("/image", addICOImage)
		developerGroup.PUT("/ico/request", setICORequestedStatus)
	}

	auditorGroup := router.Group("/auditor")
	auditorGroup.Use(AuditorAuth())
	{
		auditorGroup.GET("/ico", getAuditorICOs)
	}

	//router.POST("/audit", addAuditor)
	//router.GET("/ico/audit", getICOAuditors)

	icoGroup := router.Group("/ico")
	{
		icoGroup.GET("/", getICO)
		icoGroup.GET("/all", getAllICOs)
		icoGroup.GET("/created", getCreatedICOs)
		icoGroup.GET("/opened", getOpennedICOs)
		icoGroup.GET("/success", getSuccessICOs)
		icoGroup.GET("/failed", getFailedICOs)
		icoGroup.GET("/withdrawn", getWithdrawnICOs)
		icoGroup.GET("/overdue", getOverdueICOs)
		icoGroup.GET("/image", getICOImage)
	}

}
