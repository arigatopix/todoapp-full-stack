package routers

import (
	"net/http"
	"server/pkg/app"
	"server/pkg/e"
	"server/pkg/utils"
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

	// check type of email
	httpCode, errCode := app.BindAndValid(ctx, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := services.User{
		Email: form.Email,
	}

	isExist, _ := authService.ExistByEmail()

	if isExist {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_EXISTED, nil)
		return
	}

	user, err := authService.Add()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_REGISTER_USER, nil)
		return
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	resData := map[string]string{
		"token": token,
		"email": user.Email,
	}

	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

// @desc    Get information user
// @route   GET /api/auth/me
// @access  Private
// @Success 200
// @Failure 400
func GetMe(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	userId, err := strconv.Atoi(ctx.GetString("userId"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := services.User{}

	user, err := authService.Get(userId)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_EXISTED, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}
