package routers

import (
	"net/http"
	"server/pkg/app"
	"server/pkg/e"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int64  `json:"userId"`
}

var todos = []Todo{
	{
		ID:        1,
		Title:     "Do some yoga",
		Completed: true,
		UserId:    1,
	},
	{
		ID:        2,
		Title:     "Practice meditation",
		Completed: false,
		UserId:    1,
	},
	{
		ID:        3,
		Title:     "Download a leadership and/or business audiobook",
		Completed: true,
		UserId:    2,
	},
	{
		ID:        4,
		Title:     "Coding",
		Completed: false,
		UserId:    3,
	},
}

func GetTodos(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	appG.Response(http.StatusOK, e.SUCCESS, todos)
}

func AddTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	appG.Response(http.StatusOK, e.SUCCESS, "add todo")
}

func GetTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, "Get todo id : "+strconv.Itoa(id))
}

func UpdateTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, "UpdateTodo id : "+strconv.Itoa(id))
}

func DeleteTodo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, "DeleteTodo id : "+strconv.Itoa(id))
}
