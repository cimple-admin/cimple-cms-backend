package main

import (
	_ "github.com/cimple-admin/cimple-cms-backend/docs"
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @BasePath /

// PingExample godoc
// @Summary ping example
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"ping": 1}"
// @Router /ping [get]
func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pong": 1,
	})
}

// 整个项目的入口
func main() {
	app := pkg.GetInstance()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/ping", pong)

	app.Run()
}
