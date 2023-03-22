package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/gqframe/service"
)

var (
	s = service.NewFront()
)

func NewFront() *Front {
	return &Front{}
}

type Front struct {
}

func (c Front) Index(ctx *gin.Context) {
	s.Index(ctx)

}

func (c Front) Manufacturers(ctx *gin.Context) {

}

func (c Front) Manufacturer(ctx *gin.Context) {

}

func (c Front) Categories(ctx *gin.Context) {

}

func (c Front) Category(ctx *gin.Context) {

}

func (c Front) Product(ctx *gin.Context) {

}

func (c Front) NewsCategories(ctx *gin.Context) {

}

func (c Front) NewsCategory(ctx *gin.Context) {

}

func (c Front) Article(ctx *gin.Context) {

}

func (c Front) About(ctx *gin.Context) {

}

func (c Front) Info(ctx *gin.Context) {

}

func (c Front) Inquiry(ctx *gin.Context) {

}

func (c Front) Contact(ctx *gin.Context) {

}

func (c Front) PostMessage(ctx *gin.Context) {

}

func (c Front) PostInquiry(ctx *gin.Context) {

}
