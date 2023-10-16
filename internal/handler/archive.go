package handler

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	// "github.com/alexmullins/zip"

	"github.com/gin-gonic/gin"
)

func (h Handler) information(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Only POST requests are allowed.",
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zipInfo, err := h.service.ExtractArhiveInfo(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, zipInfo)
}

func (h Handler) files(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Only POST requests are allowed.",
		})
		return
	}

	err := c.Request.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	files := c.Request.MultipartForm.File["files[]"]

	zipFile := new(bytes.Buffer)
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		switch file.Header.Get("Content-Type") {
		case "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			"application/xml",
			"image/jpeg",
			"image/png":
			f, err := file.Open()
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			defer f.Close()

			filename := file.Filename
			fmt.Println(filename)

			h := &zip.FileHeader{Name: filename, Method: zip.Deflate, Flags: 0x800, Modified: time.Now()}

			zipFileInArchive, err := zipWriter.CreateHeader(h)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			_, err = io.Copy(zipFileInArchive, f)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		default:
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	err = zipWriter.Close()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Content-Disposition", "attachment; filename=archive.zip")
	c.Header("Content-Type", "application/zip")

	_, err = io.Copy(c.Writer, zipFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
