package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"srv/databases"
)

type (
	Controllers struct{}
)

func (ctrl Controllers) SayHi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi",
		"err":     nil,
		"data":    nil,
	})

	return
}

func (ctrl Controllers) GetAllContent(c *gin.Context) {
	d, err := databases.FindAllContent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": nil,
			"err":     err.Error(),
			"data":    nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": nil,
		"err":     nil,
		"data":    d,
	})

	return
}
