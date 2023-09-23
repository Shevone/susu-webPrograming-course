package handlers

import (
	"github.com/gin-gonic/gin"
	"web-programing-susu/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	router.GET("/", h.redirectToIndexHandler)
	books := router.Group("/books")
	{
		books.GET("/", h.redirectToIndexHandler)
		books.GET("/:page", h.indexHandler)
		books.GET("/delete/:id", h.deleteHandler)
		create := books.Group("/create")
		{
			create.GET("/", h.getPageToCreatHandler)
			create.POST("/", h.createBookHandler)
		}
		edit := books.Group("/edit")
		{
			edit.GET("/:id", h.editPage)
			edit.POST("/:id", h.editHandler)
		}
	}
	return router
}
