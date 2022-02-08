// Package middleware
// 检查是否正确安装中间件
// 错误码起始数字 1
package middleware

import (
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg"
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func CheckIsUnstall() gin.HandlerFunc {
	return func(c *gin.Context) {
		installed := pkg.AppInstance.IsInstalled()
		if installed {
			c.Next()
		} else {
			response.NewJson().OK(c, 10001, nil, "系统未安装，即将进入安装页面")
			c.Abort()
		}
	}
}

func CheckIsInstalled() gin.HandlerFunc {
	return func(c *gin.Context) {
		installed := pkg.AppInstance.IsInstalled()
		if installed {
			response.NewJson().OK(c, 10002, nil, "系统已安装")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
