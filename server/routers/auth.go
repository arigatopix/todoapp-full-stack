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
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3"`
	PasswordConfirm string `form:"passwordConfirm" json:"passwordConfirm" binding:"required,eqfield=Password"`
}

type ResponseUser struct {
	ID    int
	Email string
}

type UserLogin struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3"`
}

func sendTokenResponse(httpCode int, user *models.User, appG app.Gin) {

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	// set cookie for browser
	appG.C.SetCookie("token", token, 60*60*5, "/", "", false, true)

	resData := map[string]string{
		"token": token,
		"email": user.Email,
	}

	appG.Response(httpCode, e.SUCCESS, resData)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

	hashedPassword,_ := HashPassword(form.Password)

	authService := services.User{
		Email: form.Email,
		Password: hashedPassword,
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
		Password: form.Password,
	}

	user, err := authService.Login()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_NOT_EXIST, nil)
		return
	}

	// Check password
	// 1) Retrieve user from database

	// 2) Compare password from db and hash
	match := CheckPasswordHash(form.Password, user.Password)

	if !match {
		appG.Response(http.StatusBadRequest, e.ERROR_WRONG_PASSWORD, nil)
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

	appG.C.SetCookie("token", "", 0, "", "", false, true)

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": "",
	})
}
