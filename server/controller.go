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

// Get current developer's ICOs
func getDevelopersICOs(c *gin.Context) {
	var developer data.Developer
	developer.Address = c.Query("address")
	if len(developer.Address) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := developer.GetDeveloper()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	icos, err := developer.GetICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Get current auditor's ICO's
func getAuditorICOs(c *gin.Context) {
	var auditor data.Auditor
	auditor.Address = c.Query("address")
	if len(auditor.Address) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := auditor.GetAuditor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	icos, err := auditor.GetICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Returns all ICO's with status created
func getCreatedICOs (c *gin.Context) {
	icos, err := data.GetCreatedICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Add auditor to ICO
func addAuditorToICO(c *gin.Context) {
	type Request struct {
		ICO 	data.ICO		`json:"ico"`
		Auditor data.Auditor	`json:"auditor"`
	}
	var req Request
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ico := req.ICO
	err = ico.GetICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	auditor := req.Auditor
	err = auditor.GetAuditor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = ico.AddAuditorToICO(auditor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Get current ICO auditors
func getICOAuditors(c *gin.Context) {
	var ico data.ICO
	ico.Address = c.Query("address")
	if len(ico.Address) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := ico.GetICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	auditors, err := ico.GetICOAuditors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if auditors == nil {
		auditors = make([]data.Auditor, 0)
	}
	c.JSON(http.StatusOK, auditors)
}

// Change ICO status to opened
func publishICO(c *gin.Context) {
	var ico data.ICO
	err := c.BindJSON(&ico)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = ico.PublishICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}