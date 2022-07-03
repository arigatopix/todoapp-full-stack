package app

import (
	"net/http"

	"server/pkg/e"

	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return 200, 400
	}

	return http.StatusOK, e.SUCCESS
}
