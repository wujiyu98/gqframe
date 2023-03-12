package repository

import (
	"github.com/wujiyu98/gqframe/model"
)

func GetArticleByID(id uint) (row model.Article, err error) {
	err = db.First(&row, id).Error
	return
}

func GetAllArticle(id uint) (row model.Article, err error) {
	err = db.First(&row, id).Error
	return
}
