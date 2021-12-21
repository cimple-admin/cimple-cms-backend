package main

import (
	"fmt"
	_ "github.com/cimple-admin/cimple-cms-backend/docs"
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
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
	hmacSampleSecret := []byte("abcd")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println("err: ", err, " singStr: ", tokenString)
	c.JSON(http.StatusOK, gin.H{
		"pong":        1,
		"tokenString": tokenString,
	})
}

// 整个项目的入口
func main() {
	app := pkg.GetInstance()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/ping", pong)

	app.Run()
}
