package main

import (
	"net/http"

	"../server/data"

	"github.com/gin-gonic/gin"
)

// Add new developer to db
func addDeveloper(c *gin.Context) {
	var developer data.Developer
	err := c.BindJSON(&developer)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = developer.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func addAuditor(c *gin.Context) {
	var auditor data.Auditor
	err := c.BindJSON(&auditor)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = auditor.AddAuditor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
