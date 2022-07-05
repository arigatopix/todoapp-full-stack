package routers

import (
	"net/http"
	"server/pkg/app"
	"server/pkg/e"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @desc    Get all todos
// @route   GET /api/todos
// @access  Private
// @Success 200
// @Failure 400
func GetTodos(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	userId, err := strconv.Atoi(ctx.GetString("userId"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	todoService := services.Todo{UserId: userId}

	todos, err := todoService.GetAll()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, todos)
}

type AddTodoForm struct {
	Title     string `form:"title" json:"title" binding:"required" validate:"min=1"`
	Completed *bool  `form:"completed" json:"completed" binding:"required" validate:"boolean"`
}

// @desc    Add todo
// @route   POST /api/todos
// @access  Private
// @Success 200
// @Failure 400
func AddTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	var form AddTodoForm

	httpCode, errCode := app.BindAndValid(ctx, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userId, err := strconv.Atoi(ctx.GetString("userId"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	todoService := services.Todo{
		Title:     form.Title,
		Completed: *form.Completed,
		UserId:    userId,
	}

	todo, err := todoService.Add()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_TODO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, todo)
}

// @desc    Get single todo
// @route   GET /api/todos/:id
// @access  Private
// @Success 200
// @Failure 400
func GetTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userId, err := strconv.Atoi(ctx.GetString("userId"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	todoService := services.Todo{ID: id, UserId: userId}

	todo, err := todoService.Get()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_TODO_NOT_EXIST, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, todo)
}

type UpdateTodoForm struct {
	Title     string `form:"title" binding:"required"`
	Completed *bool  `form:"completed" json:"completed" binding:"required" validate:"boolean"`
}

// @desc    Update todo
// @route   PUT /api/todos/:id
// @access  Private
// @Success 200
// @Failure 400
func UpdateTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	var form UpdateTodoForm

	httpCode, errCode := app.BindAndValid(ctx, &form)

	if httpCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userId, err := strconv.Atoi(ctx.GetString("userId"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	todoService := services.Todo{
		ID:        id,
		Title:     form.Title,
		Completed: *form.Completed,
		UserId:    userId,
	}

	todo, err := todoService.Update(id)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_UPDATE_TODO, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, todo)
}

// @desc    DELELTE todo
// @route   DELETE /api/todos/:id
// @access  Private
// @Success 200
// @Failure 400
func DeleteTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userId, err := strconv.Atoi(ctx.GetString("userId"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	todoService := services.Todo{ID: id, UserId: userId}

	// Find todo with userId
	if _, err := todoService.Get(); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_DOES_NOT_BELONG_TODO, nil)
		return
	}

	todoService.Delete()

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
