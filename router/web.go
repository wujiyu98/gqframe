package router

import "github.com/gin-gonic/gin"

func web(e *gin.Engine) {
	r := e.Group("/")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	r.GET("/contact", func(ctx *gin.Context) {
		ctx.HTML(200, "contact.html", nil)
	})

}
