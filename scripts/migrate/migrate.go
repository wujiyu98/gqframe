package main

import (
	_ "github.com/wujiyu98/gqframe/config"
	"github.com/wujiyu98/gqframe/database"
	"github.com/wujiyu98/gqframe/entity"
)

var db = database.DB

func migrate() {
	db.AutoMigrate(
		&entity.Language{},
		&entity.Site{},
		&entity.Banner{},
		&entity.Seo{},
		&entity.ArticleCategory{},
		&entity.Article{},
		&entity.Category{},
		&entity.Product{},
		&entity.Attribute{},
		&entity.ProductAttribute{},
		&entity.User{},
		&entity.Address{},
		&entity.Message{},
		&entity.Enquiry{},
		&entity.EnquiryItem{},
		&entity.Order{},
		&entity.OrderItem{},
	)

}

func main() {
	migrate()

}
