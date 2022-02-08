package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Json struct{}

func NewJson() Json {
	return Json{}
}

func (j Json) OK(c *gin.Context, code int, data map[string]interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}
