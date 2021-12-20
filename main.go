package main

import (
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := pkg.GetInstance()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": 1,
		})
	})

	app.Run()
}
