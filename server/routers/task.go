package routers

import (
	"net/http"
	"strconv"

	"server/pkg/app"
	"server/pkg/e"
	s "server/services"

	"github.com/gin-gonic/gin"
)

type AddTaskForm struct {
	Text     string `form:"text" json:"text" binding:"required" validate:"min=1,max=500"`
	Day      string `form:"day" json:"day" binding:"required"`
	Reminder bool   `form:"reminder" json:"reminder" validate:"boolean"`
}

// @Summary Create Task
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/ [post]
func AddTask(c *gin.Context) {
	// pass context to controller
	appG := app.Gin{C: c}

	var form *AddTaskForm

	err := c.ShouldBindJSON(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	todoService := s.Task{
		Text:     form.Text,
		Day:      form.Day,
		Reminder: form.Reminder,
	}

	task, err := todoService.Add()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ADD_TODO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, task)
}

// @Summary Get task with ID
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/:id [get]
func GetTask(c *gin.Context) {

	// converse string to int
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	// create instance
	todoService := s.Task{ID: id}

	// call service with instance
	task, err := todoService.Get()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// @Summary Get multiple Tasks
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks [get]
func GetTasks(c *gin.Context) {
	todoService := s.Task{}

	tasks, err := todoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tasks,
	})
}

// @Summary Update Task by id
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/:id [put]
func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var form *AddTaskForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	todoService := s.Task{
		ID:       id,
		Text:     form.Text,
		Day:      form.Day,
		Reminder: form.Reminder,
	}

	task, err := todoService.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// @Summary Delete Task by id
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/:id [delete]
func DeleteTask(c *gin.Context) {
	// converse string to int
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	todoService := s.Task{ID: id}

	if err := todoService.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
