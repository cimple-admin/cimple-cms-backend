package pkg

import "github.com/gin-gonic/gin"

type app struct {
	*gin.Engine
}

var appInstance *app

func GetInstance() *app {
	if appInstance == nil {
		appInstance = &app{
			gin.Default(),
		}
	}

	return appInstance
}
