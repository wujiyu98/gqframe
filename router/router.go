package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	e := gin.Default()
	webRouter(e)
	apiRouter(e)
	return e
}
