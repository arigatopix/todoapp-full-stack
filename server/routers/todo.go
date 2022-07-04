package routers

import (
	"net/http"
	"server/pkg/app"
	"server/pkg/e"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	todoService := services.Todo{}

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

func AddTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	var form AddTodoForm

	httpCode, errCode := app.BindAndValid(ctx, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	todoService := services.Todo{
		Title:     form.Title,
		Completed: *form.Completed,
	}

	todo, err := todoService.Add()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_TODO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, todo)
}

func GetTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	todoService := services.Todo{ID: id}

	todo, err := todoService.Get(id)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_TODO_NOT_EXIST, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, todo)
}

type UpdateTodoForm struct {
	Title     string `form:"title" binding:"required"`
	Completed *bool  `form:"completed" json:"completed" binding:"required" validate:"boolean"`
}

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

	todoService := services.Todo{
		ID:        id,
		Title:     form.Title,
		Completed: *form.Completed,
	}

	todo, err := todoService.Update(id)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_UPDATE_TODO, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, todo)
}

func DeleteTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	todoService := services.Todo{ID: id}

	if err := todoService.Delete(id); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DELETE_TODO_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
