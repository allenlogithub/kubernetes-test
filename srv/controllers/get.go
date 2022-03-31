package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controllers struct{}
)

func (ctrl Controllers) SayHi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi",
		"err":     nil,
	})

	return
}

func (ctrl Controllers) SayHiHi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "HiHi",
		"err":     nil,
	})

	return
}
