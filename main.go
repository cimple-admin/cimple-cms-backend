package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/cimple-admin/cimple-cms-backend/docs"
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"pong":        1,
		"tokenString": tokenString,
	})
}

// 整个项目的入口
func main() {
	app := pkg.AppInstance
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/ping", pong)
	app.GET("/checkInstalled", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"installed": app.IsInstalled(),
		})
	})

	app.Run()
}
