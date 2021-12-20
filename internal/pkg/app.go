package pkg

import "github.com/gin-gonic/gin"

// 自行封装的 app 实例，基础就是 gin 的 engine
type app struct {
	*gin.Engine
}

var appInstance *app

// GetInstance 获取实例，单例模式
func GetInstance() *app {
	if appInstance == nil {
		appInstance = &app{
			gin.Default(),
		}
	}

	return appInstance
}
