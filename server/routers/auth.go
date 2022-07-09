package routers

import (
	"net/http"
	"server/models"
	"server/pkg/app"
	"server/pkg/e"
	"server/pkg/utils"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRegisterForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	// TODO password string `form:"password" json:"password" binding:"required,min=3"`
}

type ResponseUser struct {
	ID    int
	Email string
}

type UserLogin struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	// TODO password string `form:"password" json:"password" binding:"required,min=3"`
}

func sendTokenResponse(httpCode int, user *models.User, appG app.Gin) {

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	// set cookie for browser
	appG.C.SetCookie("token", token, 60*60*5, "/", "", true, true)

	resData := map[string]string{
		"token": token,
		"email": user.Email,
	}

	appG.Response(httpCode, e.SUCCESS, resData)
}

// @desc    Register new user
// @route   POST /api/auth/register
// @payload {email, password}
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

	sendTokenResponse(http.StatusOK, user, appG)
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

// @desc    Login
// @route   GET /api/auth/login
// @payload {email, password}
// @access  Public
// @Success 200
// @Failure 400
func Login(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	var form UserLogin

	httpCode, errCode := app.BindAndValid(ctx, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := services.User{
		Email: form.Email,
	}

	user, err := authService.Login()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_NOT_EXIST, nil)
		return
	}

	sendTokenResponse(http.StatusOK, user, appG)
}

// @desc    Logout
// @route   POST /api/auth/logout
// @access  Private
// @Success 200
// @Failure 400
func Logout(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	appG.C.SetCookie("token", "", 0, "", "", true, true)

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": "",
	})
}
