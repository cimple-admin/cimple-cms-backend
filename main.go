package main

import (
	_ "github.com/cimple-admin/cimple-cms-backend/docs"
	"github.com/cimple-admin/cimple-cms-backend/internal/pkg"
	"github.com/cimple-admin/cimple-cms-backend/router"
)

// 整个项目的入口
func main() {
	app := pkg.AppInstance
	router.InitRoute()
	err := app.Run()
	if err != nil {
		return
	}
}
