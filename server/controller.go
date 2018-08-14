package main

import (
	"net/http"

	"../server/data"

	"github.com/gin-gonic/gin"
	"os"
	"log"
	"io/ioutil"
	"image/jpeg"
	"github.com/nfnt/resize"
)

const imagesDir = "./ico_images/"

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
	if auditors == nil {
		auditors = make([]data.Auditor, 0)
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

// Upload ICO image
func addICOImage(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	address := c.Request.FormValue("address")
	if len(address) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	log.Println(address)
	ico := data.ICO{Address: address}
	err = ico.GetICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// Resize file
	img, err := jpeg.Decode(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	file.Close()
	resizedImg := resize.Resize(400, 260, img, resize.Lanczos3)

	out, err := os.Create(imagesDir + address + ".jpeg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer out.Close()
	err = jpeg.Encode(out, resizedImg, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Get ICO image
func getICOImage(c *gin.Context) {
	var ico data.ICO
	ico.Address = c.Query("address")
	if len(ico.Address) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	path := imagesDir + ico.Address + ".jpeg"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Data(http.StatusOK, "image/jpeg", b)
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
	if icos == nil {
		icos = make([]data.ICO, 0)
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
	if icos == nil {
		icos = make([]data.ICO, 0)
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status opened
func getOpennedICOs (c *gin.Context) {
	icos, err := data.GetOpenedICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with statuses success and requested
func getSuccessICOs(c *gin.Context) {
	icos, err := data.GetSuccessICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status failed
func getFailedICOs(c *gin.Context) {
	icos, err := data.GetFailedICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status withdrawn
func getWithdrawnICOs(c *gin.Context) {
	icos, err := data.GetWithdrawnICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if icos == nil {
		icos = make([]data.ICO, 0)
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

// Change ICO status to success
func setICOSucccessStatus(c *gin.Context) {
	var ico data.ICO
	err := c.BindJSON(&ico)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = ico.SetSuccessStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Change ICO status to requested
func setICORequestedStatus(c *gin.Context) {
	var ico data.ICO
	err := c.BindJSON(&ico)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = ico.SetRequestedStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Change ICO status to withdrawn
func setICOWithdrawnStatus(c *gin.Context) {
	var ico data.ICO
	err := c.BindJSON(&ico)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = ico.SetWithdrawnStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}