package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/viky1sr/latihan_go-lang.git/dto"
	"net/http"
)

func ValidationUpdate(context *gin.Context, userUpdateDTO dto.UserUpdateDTO) {
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := BuildErrorResponse("Failed to process request", errDTO.Error(), EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
}
