package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自行封装的 app 实例，基础就是 gin 的 engine
type app struct {
	*gin.Engine
}

// AppInstance 对外暴露的 app 实例
var AppInstance *app

func init() {
	AppInstance = &app{
		gin.Default(),
	}
}

func (app *app) Run() error {
	server := &http.Server{Addr: ":8080", Handler: app}
	return server.ListenAndServe()
}
