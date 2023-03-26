package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/gqframe/config"
	"github.com/wujiyu98/gqframe/view"
)

var e *gin.Engine

func init() {
	e = gin.Default()
	if config.V.GetString("env") == "dev" {
		e.HTMLRender = view.LoadTemplates(funcMap)
	} else {
		e.HTMLRender = view.LoadTemplatesFs(funcMap)
	}

	web(e)
	api(e)
}

func Run() {
	e.Run(config.V.GetString("port"))
}
