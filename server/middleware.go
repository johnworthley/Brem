package main

import (
	"../server/data"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)

// Superuser requests middleware
func SuperuserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		iAddress := session.Get("address")
		if iAddress == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		address := iAddress.(string)
		iSign := session.Get("sign")
		if iSign == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		sign := iSign.(string)
		// Validate session
		isValid, err := validate(address, sign)
		if err != nil || !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Check superuser address
		isSuperuser, err := data.IsSuperuser(address)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if isSuperuser {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}