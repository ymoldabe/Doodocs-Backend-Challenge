package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) information(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	zipInfo, err := h.service.ExtractArhiveInfo(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, zipInfo)
}

func (h Handler) files(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	files := c.Request.MultipartForm.File["files[]"]

	zipFile, code, err := h.service.CreateArchive(files)
	if err != nil {
		if code == http.StatusBadRequest {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": http.StatusText(http.StatusBadRequest),
			})
			return
		}
		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=archive.zip")
	c.Header("Content-Type", "application/zip")

	_, err = io.Copy(c.Writer, zipFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
