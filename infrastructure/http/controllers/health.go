package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() HealthController {
	return HealthController{}
}

func (c *HealthController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/health", GetStatus)
}

func GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
