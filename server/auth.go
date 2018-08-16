package main

import (
	"../server/data"
	"strconv"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
)

func validate(address string, sign string) (res bool, err error) {
	validationMsg := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(address)) + address
	hash := crypto.Keccak256([]byte(validationMsg))

	signature, err := hexutil.Decode(sign)
	signature[64] -= 27

	pk, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		return
	}

	ecdsaPK, err := crypto.UnmarshalPubkey(pk)
	if err != nil {
		return
	}

	phHolder := crypto.PubkeyToAddress((*ecdsaPK)).String()

	res = strings.EqualFold(address, phHolder)

	return
}

// Register new developer
// @Summary Зарегистрировать застройщика
// @Description Зарегистрировать застройщика
// @Tags auth
// @Accept  json
// @Produce  json
// @Param ico body data.Developer true "Dev"
// @Router /signup [post]
func addDeveloper(c *gin.Context) {
	var developer data.Developer
	err := c.BindJSON(&developer)
	if err != nil || len(developer.Username) == 0 || len(developer.Sign) == 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// Validate request
	isValid, err := validate(developer.Address, developer.Sign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !isValid {
		c.JSON(http.StatusNotAcceptable, nil)
		return
	}

	exists, err := developer.IsExists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if exists {
		c.JSON(http.StatusOK, nil)
		return
	}

	err = developer.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// Developer login
// Register new developer
// @Summary Вход для застройщика
// @Description Вход для застройщика
// @Tags auth
// @Accept  json
// @Produce  json
// @Param ico body data.Developer true "Dev"
// @Success 200 {object} data.Developer
// @Router /login [post]
func login(c *gin.Context) {
	var developer data.Developer
	err := c.BindJSON(&developer)
	if err != nil || len(developer.Address) == 0 || len(developer.Sign) == 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// Validate request
	isValid, err := validate(developer.Address, developer.Sign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !isValid {
		c.JSON(http.StatusNotAcceptable, nil)
		return
	}

	exists, err := developer.IsExists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !exists {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	err = developer.GetDeveloper()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// Set up session
	session := sessions.Default(c)
	session.Set("address", developer.Address)
	session.Set("sign", developer.Sign)

	c.JSON(http.StatusOK, developer)
}