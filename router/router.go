package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/gqframe/config"
	"github.com/wujiyu98/gqframe/view"
)

var e *gin.Engine

func init() {
	e = gin.Default()
	e.HTMLRender = view.LoadTemplates()
	// e.HTMLRender = view.LoadTemplatesFs2()
	web(e)
	api(e)
}

func Run() {
	e.Run(config.V.GetString("server.port"))
}
