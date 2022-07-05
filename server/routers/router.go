package routers

import (
	"net/http"
	"server/config"
	"server/middlewares"

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

	// Route /api/todos
	todo := apiGroups.Group("/todos")
	{
		todo.GET("", GetTodos)
		todo.POST("", AddTodo)
		todo.PUT("/:id", UpdateTodo)
		todo.DELETE("/:id", DeleteTodo)
		todo.GET("/:id", GetTodo)
	}

	// Route /api/auth
	auth := apiGroups.Group("/auth")
	{
		auth.GET("/me", middlewares.Protect(), GetMe)
		auth.POST("/register", Register)
	}

	PORT := env.PORT

	if PORT == "" {
		PORT = "5001"
	}

	r.Run("0.0.0.0:" + PORT)

	return r
}
