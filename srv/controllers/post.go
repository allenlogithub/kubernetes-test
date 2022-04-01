package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"srv/databases"
)

type (
	PostContentRequest struct {
		Content string `json:"Content" binding:"required"`
	}
)

func (ctrl Controllers) PostContent(c *gin.Context) {
	var r PostContentRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
			"data":    nil,
		})

		return
	}

	d := databases.InsertContentRequest{
		Content: r.Content,
	}
	if err := databases.InsertContent(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
			"data":    nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": nil,
		"err":     nil,
		"data":    nil,
	})

	return
}
