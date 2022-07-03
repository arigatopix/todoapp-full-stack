package app

import (
	"net/http"

	"server/pkg/e"

	"github.com/gin-gonic/gin"
)

func BindAndValid(ctx *gin.Context, form interface{}) (int, int) {
	err := ctx.ShouldBindJSON(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
