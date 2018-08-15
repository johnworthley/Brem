package main

import (
	"net/http"

	"../server/data"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const imagesDir = "./ico_images/"

// FILLING IN UPDATER
// Add new auditor to db
//func addAuditor(c *gin.Context) {
//	var auditor data.Auditor
//	err := c.BindJSON(&auditor)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	err = auditor.AddAuditor()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, err)
//		return
//	}
//
//	c.JSON(http.StatusCreated, nil)
//}

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
	iDeveloper, exists := c.Get("dev")
	if !exists || iDeveloper == nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	developer := iDeveloper.(data.Developer)
	if strings.EqualFold(developer.Address, ico.Developer.Address) {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	err = ico.CreateICO()
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
	ico := data.ICO{Address: address}
	err = ico.GetICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	iDeveloper, exists := c.Get("dev")
	if !exists || iDeveloper == nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	developer := iDeveloper.(data.Developer)
	if developer.ID != ico.Developer.ID {
		c.JSON(http.StatusNotAcceptable, nil)
		return
	}
	// Resize file
	img, _, err := image.Decode(file)
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
	iDeveloper, exists := c.Get("dev")
	if iDeveloper == nil || !exists {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	developer := iDeveloper.(data.Developer)
	icos, err := developer.GetICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Get current auditor's ICO's
func getAuditorICOs(c *gin.Context) {
	iAuditor, exists := c.Get("auditor")
	if iAuditor == nil || !exists {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	auditor := iAuditor.(data.Auditor)
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

// Get ICOs in page
func getAllICOs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	icos, err := data.GetAllICOs(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Returns all ICOs with status created
func getCreatedICOs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	icos, err := data.GetCreatedICOs(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status opened
func getOpennedICOs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	icos, err := data.GetOpenedICOs(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with statuses success and requested
func getSuccessICOs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	icos, err := data.GetSuccessICOs(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status failed
func getFailedICOs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	icos, err := data.GetFailedICOs(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status withdrawn
func getWithdrawnICOs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	icos, err := data.GetWithdrawnICOs(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status overdue
func getOverdueICOs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	icos, err := data.GetOverdueICOs(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, icos)
}

// Add auditor to ICO
func addAuditorToICO(c *gin.Context) {
	type Request struct {
		ICO     data.ICO     `json:"ico"`
		Auditor data.Auditor `json:"auditor"`
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
//func getICOAuditors(c *gin.Context) {
//	var ico data.ICO
//	ico.Address = c.Query("address")
//	if len(ico.Address) == 0 {
//		c.JSON(http.StatusBadRequest, nil)
//		return
//	}
//	err := ico.GetICO()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, err)
//		return
//	}
//	auditors, err := ico.GetICOAuditors()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, err)
//		return
//	}
//	if auditors == nil {
//		auditors = make([]data.Auditor, 0)
//	}
//	c.JSON(http.StatusOK, auditors)
//}

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

//// Change ICO status to success
//func setICOSucccessStatus(c *gin.Context) {
//	var ico data.ICO
//	err := c.BindJSON(&ico)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	err = ico.SetSuccessStatus()
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	c.JSON(http.StatusOK, nil)
//}

// Change ICO status to requested
func setICORequestedStatus(c *gin.Context) {
	iDeveloper, exists := c.Get("dev")
	if !exists || iDeveloper == nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	developer := iDeveloper.(data.Developer)
	var ico data.ICO
	err := c.BindJSON(&ico)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = ico.GetICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if developer.ID != ico.Developer.ID {
		c.JSON(http.StatusNotAcceptable, nil)
		return
	}
	err = ico.SetRequestedStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

//// Change ICO status to withdrawn
//func setICOWithdrawnStatus(c *gin.Context) {
//	var ico data.ICO
//	err := c.BindJSON(&ico)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	err = ico.SetWithdrawnStatus()
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	c.JSON(http.StatusOK, nil)
//}
