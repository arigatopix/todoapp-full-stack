package routers

import (
	"net/http"
	"server/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Load env
	env := config.LoadENV()

	r := gin.New()

	r.Use(cors.Default())

	apiGroups := r.Group("/api")

	apiGroups.GET("/hi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})

	task := apiGroups.Group("/tasks")
	{
		// GET /tasks
		task.GET("", GetTasks)

		// GET /tasks/:id
		task.GET("/:id", GetTask)

		// POST /tasks
		task.POST("", AddTask)

		// PUT /tasks/:id
		task.PUT("/:id", UpdateTask)

		// DELETE /tasks/:id
		task.DELETE("/:id", DeleteTask)
	}

	todo := apiGroups.Group("/todos")
	{
		// GET /api/todos
		todo.GET("", GetTodos)
		// POST /api/todos
		todo.POST("", AddTodo)
		// UPDATE /api/todos/:id
		todo.PUT("/:id", UpdateTodo)
		// DELETE /api/todos/:id
		todo.DELETE("/:id", DeleteTodo)
		// GET /api/todos/:id
		todo.GET("/:id", GetTodo)
	}

	PORT := env.PORT

	if PORT == "" {
		PORT = "5001"
	}

	r.Run("0.0.0.0:" + PORT)

	return r
}
