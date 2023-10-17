package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *HandlerType) InitRouters() *gin.Engine {

	router := gin.New()

	api := router.Group("/api")
	{
		archive := api.Group("/archive")
		{
			archive.Any("/information", h.information)
			archive.Any("files", h.files)
		}
		mail := api.Group("mail")
		{
			mail.Any("/file", h.file)
		}
	}

	return router
}
