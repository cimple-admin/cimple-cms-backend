package router

import (
	"fmt"
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg"
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg/response"
	"github.com/cimple-admin/cimple-cms-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func InitRoute() {
	app := pkg.AppInstance
	// 全局使用检测安装中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	app.Use(cors.New(config))
	app.Use(middleware.CheckIsUnstall())
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/ping", pong)
	app.GET("/checkInstalled", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"installed": app.IsInstalled(),
		})
	})
	app.GET("/install", middleware.CheckIsInstalled(), func(c *gin.Context) {
		response.NewJson().OK(c, 10003, nil, "安装界面")
	})
	app.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}

// @BasePath /
// PingExample godoc
// @Summary ping example
// @Description do ping
// @Accept json4
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
