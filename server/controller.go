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

// Get all registered BREM auditors
// @Summary Список зарегистрированных аудиторов
// @Tags superuser
// @Description Получить всех зарегистрированных аудиторов
// @Produce  json
// @Success 200 {array} data.Auditor
// @Router /super/audit [get]
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
// @Summary Добавить ICO
// @Description Добавить ICO
// @Tags developer
// @Accept  json
// @Produce  json
// @Param ico body data.ICO true "ICO (без Id)"
// @Router /dev/ico [post]
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
	if !strings.EqualFold(developer.Address, ico.Developer.Address) {
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
// @Summary Загрузить изображение ICO
// @Description Загрузить изображение ICO
// @Tags developer
// @Accept  multipart/form-data
// @Produce  json
// @Param  address path string true "ICO address"
// @Param file formData file true "account image"
// @Router /dev/image [post]
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
// @Summary Получить изображение ICO
// @Description Получить изображение ICO
// @Tags common
// @Accept  image/jpeg
// @Produce  json
// @Param  address path string true "ICO address"
// @Router /ico/image [get]
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

// Get ICO info
// @Summary Получить информацию об ICO
// @Description Получить информацию об ICO
// @Tags common
// @Accept  json
// @Produce  json
// @Param  address path string true "ICO address"
// @Success 200 {object} data.ICO
// @Router /ico [get]
func getICO(c *gin.Context) {
	var ico data.ICO
	ico.Address = c.Query("address")
	if len(ico.Address) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	ico.Address = strings.ToLower(ico.Address)
	err := ico.GetICO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ico)
}

// Get ICOs info
// @Summary Получить все ICO для открытия
// @Description Получить все ICO для открытия
// @Tags superuser
// @Accept  json
// @Produce  json
// @Success 200 {array} data.ICO
// @Router /super/ico [get]
func getSuperusersICOs(c *gin.Context) {
	icos, err := data.GetSuperuserICOs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Get ICOs info
// @Summary Получить все ICO данного застройщика
// @Description Получить все ICO данного застройщика
// @Tags developer
// @Accept  json
// @Produce  json
// @Success 200 {array} data.ICO
// @Router /dev/ico [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Get ICOs info
// @Summary Получить все ICO аудитора для подстверждения вывода (с запросами на вывод)
// @Description Получить все ICO аудитора для подстверждения вывода (с запросами на вывод)
// @Tags auditor
// @Accept  json
// @Produce  json
// @Success 200 {array} data.ICO
// @Router /auditor/ico [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Get ICOs in page
// @Summary Получить страницу ICO (6 ICO)
// @Description Получить страницу ICO
// @Tags common
// @Accept  json
// @Produce  json
// @Param  page path int true "Page number"
// @Success 200 {array} data.ICO
// @Router /ico/all [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns all ICOs with status created
// @Summary Получить страницу ICO (статус created)
// @Description Получить страницу ICO (статус created)
// @Tags common
// @Accept  json
// @Produce  json
// @Param  page path int true "Page number"
// @Success 200 {array} data.ICO
// @Router /ico/created [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status opened
// @Summary Получить страницу ICO (статус opened)
// @Description Получить страницу ICO (статус opened)
// @Tags common
// @Accept  json
// @Produce  json
// @Param  page path int true "Page number"
// @Success 200 {array} data.ICO
// @Router /ico/opened [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with statuses success and requested
// @Summary Получить страницу ICO (статус success и requested)
// @Description Получить страницу ICO (статус success и requested)
// @Tags common
// @Accept  json
// @Produce  json
// @Param  page path int true "Page number"
// @Success 200 {array} data.ICO
// @Router /ico/success [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status failed
// @Summary Получить страницу ICO (статус failed)
// @Description Получить страницу ICO (статус failed)
// @Tags common
// @Accept  json
// @Produce  json
// @Param  page path int true "Page number"
// @Success 200 {array} data.ICO
// @Router /ico/failed [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status withdrawn
// @Summary Получить страницу ICO (статус withdrawn)
// @Description Получить страницу ICO (статус withdrawn)
// @Tags common
// @Accept  json
// @Produce  json
// @Param  page path int true "Page number"
// @Success 200 {array} data.ICO
// @Router /ico/withdrawn [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Returns ICOs with status overdue
// @Summary Получить страницу ICO (статус overdue)
// @Description Получить страницу ICO (статус overdue)
// @Tags common
// @Accept  json
// @Produce  json
// @Param  page path int true "Page number"
// @Success 200 {array} data.ICO
// @Router /ico/overdue [get]
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
	if icos == nil {
		icos = make([]data.ICO, 0)
	}
	c.JSON(http.StatusOK, icos)
}

// Add auditor to ico request
type Request struct {
	ICO     data.ICO     `json:"ico"`
	Auditor data.Auditor `json:"auditor"`
}

// Add auditor to ICO
// @Summary Добавить аудитора к ICO
// @Description Добавить аудитора к ICO
// @Tags superuser
// @Accept  json
// @Produce  json
// @Param req body Request true "REQ (только адреса ICO и аудитора)"
// @Router /super/ico/audit [post]
func addAuditorToICO(c *gin.Context) {
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

// Change ICO status to opened
// @Summary Изменить статус ICO на opened
// @Description Изменить статус ICO на opened
// @Tags superuser
// @Accept  json
// @Produce  json
// @Param ico body data.ICO true "ICO (только адрес)"
// @Router /super/ico/open [post]
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

// Change ICO status to requested
// @Summary Изменить статус ICO на requested
// @Description Изменить статус ICO на requested
// @Tags developer
// @Accept  json
// @Produce  json
// @Param ico body data.ICO true "ICO (только адрес)"
// @Router /dev/ico/request [post]
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
