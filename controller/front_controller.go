package controller

import (
	"github.com/gin-gonic/gin"
)

func NewFront() *Front {
	return &Front{}
}

type Front struct {
}

// Index
func (c Front) Index(ctx *gin.Context) {

	ctx.HTML(200, "index.html", ctx.Keys)

}

// Manufacturers
func (c Front) Manufacturers(ctx *gin.Context) {

	ctx.HTML(200, "manufacturers.html", ctx.Keys)

}

// Manufacturer
func (c Front) Manufacturer(ctx *gin.Context) {

	ctx.HTML(200, "manufacturer.html", ctx.Keys)

}

// Categories
func (c Front) Categories(ctx *gin.Context) {
	ctx.HTML(200, "categories.html", ctx.Keys)

}

// Category
func (c Front) Category(ctx *gin.Context) {
	ctx.HTML(200, "category.html", ctx.Keys)

}

// Product
func (c Front) Product(ctx *gin.Context) {
	ctx.HTML(200, "product.html", ctx.Keys)

}

// NewsCategories
func (c Front) NewsCategories(ctx *gin.Context) {
	ctx.HTML(200, "news.html", ctx.Keys)

}

// NewsCategory
func (c Front) NewsCategory(ctx *gin.Context) {
	ctx.HTML(200, "news_category.html", ctx.Keys)

}

// Article
func (c Front) Article(ctx *gin.Context) {
	ctx.HTML(200, "article.html", ctx.Keys)

}

// About
func (c Front) About(ctx *gin.Context) {
	ctx.HTML(200, "about.html", ctx.Keys)

}

// Info
func (c Front) Info(ctx *gin.Context) {
	ctx.HTML(200, "info.html", ctx.Keys)

}

// Inquiry
func (c Front) Inquiry(ctx *gin.Context) {
	ctx.HTML(200, "inquiry.html", ctx.Keys)

}

// Contact
func (c Front) Contact(ctx *gin.Context) {
	ctx.HTML(200, "contact.html", ctx.Keys)

}

// PostMessage
func (c Front) PostMessage(ctx *gin.Context) {

}

// PostInquiry
func (c Front) PostInquiry(ctx *gin.Context) {

}
