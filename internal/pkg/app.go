package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自行封装的 app 实例，基础就是 gin 的 engine
type app struct {
	*gin.Engine
}

// AppInstance 对外暴露的 app 实例
var AppInstance *app
var cfg config

func init() {
	cfg = initConfig()
	gin.SetMode(cfg.App.RunMode)

	AppInstance = &app{
		gin.Default(),
	}

}

func (app *app) Run() error {
	server := &http.Server{
		Addr:         ":" + cfg.App.HttpPort,
		Handler:      app,
		ReadTimeout:  time.Duration(cfg.App.ReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(cfg.App.WriteTimeOut) * time.Second,
	}
	return server.ListenAndServe()
}
