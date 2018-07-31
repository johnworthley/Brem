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

// Add new auditor to db
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

// Get all BREM auditors
func getAllAuditors(c *gin.Context) {
	auditors, err := data.GetAllAuditors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, auditors)
}

// Add new ICO
func addICO(c *gin.Context) {
	var ico data.ICO
	err := c.BindJSON(&ico)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = ico.Developer.GetDeveloper()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = ico.AddICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func getCreatedICOs (c *gin.Context) {
	icos, err := data.GetCreatedICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}
