package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) file(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Only POST requests are allowed.",
		})
		return
	}
}
