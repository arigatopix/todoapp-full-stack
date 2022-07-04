package routers

import (
	"net/http"
	"server/pkg/app"
	"server/pkg/e"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRegisterForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

// @desc    Register new user
// @route   POST /api/auth/register
// @access  Public
// @Success 200
// @Failure 400
func Register(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	var form UserRegisterForm

	httpCode, errCode := app.BindAndValid(ctx, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := services.User{
		Email: form.Email,
	}

	user, err := userService.Add()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_REGISTER_USER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}

// @desc    Get information user
// @route   GET /api/auth/me
// @access  Private
// @Success 200
// @Failure 400
func GetMe(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userService := services.User{ID: id}

	user, err := userService.Get(id)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_NOT_EXIST, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}
