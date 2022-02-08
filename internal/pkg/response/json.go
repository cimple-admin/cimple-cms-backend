package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type json struct{}

func NewJson() json {
	return json{}
}

func (j json) OK(c *gin.Context, code int, data map[string]interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}
