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

// Auditors requests middleware
func AuditorAuth() gin.HandlerFunc {
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

		auditor := data.Auditor{Address: address}
		err = auditor.GetAuditor()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Set("auditor", auditor)
		c.Next()
	}
}

// Developer requests middleware
func DeveloperAuth() gin.HandlerFunc {
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

		developer := data.Developer{Address: address}
		isExists, err := developer.IsExists()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if !isExists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		err = developer.GetDeveloper()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Set("dev", developer)
		c.Next()
	}
}