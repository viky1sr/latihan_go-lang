package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func (c authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success login",
	})
}

func (c authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success register data.",
	})
}

func NewAuthController() AuthController {
	return &authController{}
}
