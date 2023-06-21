package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/gqframe/controller"
)

func web(e *gin.Engine) {
	r := e.Group("/")
	c := controller.NewFront()

	r.GET("/", c.Index)
	r.GET("/manufacturers/", c.Index)
	r.GET("/manufacturers/:pathname", c.Index)
	r.GET("/category/", c.Index)
	r.GET("/category/:pathname/*pathname2", c.Index)
	r.GET("/news-category/", c.Index)
	r.GET("/news-category/:pathname", c.Index)
	r.GET("/article/:pathname", c.Index)
	r.GET("/about/:pathname", c.Index)
	r.GET("/info/:pathname", c.Index)
	r.GET("/inquiry", c.Index)
	r.GET("/contact", c.Index)
	r.POST("/message", c.Index)
	r.POST("/inquiry", c.Index)
}
