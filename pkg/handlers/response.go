package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type myError struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, stratusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(stratusCode, myError{message})
}
