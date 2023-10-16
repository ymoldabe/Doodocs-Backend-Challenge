package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/models"
)

func (h Handler) file(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	emails, ok := c.GetPostForm("emails")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := h.service.SendLetters(file, emails); err != nil {
		if errors.Is(err, models.ErrBadRequest) {
			c.Status(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)

}
