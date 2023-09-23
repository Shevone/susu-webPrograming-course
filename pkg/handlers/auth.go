package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-programing-susu/pkg/models"
)

func (h *Handler) signUp(c *gin.Context) {
	//Засунуть парсинг формы в будущем
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})

}
func (h *Handler) signIn(c *gin.Context) {

}
