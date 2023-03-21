package repository

import (
	"github.com/wujiyu98/gqframe/database"
	"github.com/wujiyu98/gqframe/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func New() Repository {
	return Repository{
		DB: database.DB,
	}
}

type Repository struct {
	*gorm.DB
}

func (r Repository) GetArticles() (rows []model.Article) {
	r.Find(&rows)
	return
}

func (r Repository) GetArticleByID(id uint) (row model.Article, err error) {
	err = r.First(&row).Error
	return
}

func (r Repository) GetArticleByPathname(pathname string) (row model.Article, err error) {
	err = r.First(&row, &model.Article{Pathname: pathname}).Error
	return
}

func (r Repository) GetProducts() (rows []model.Article) {
	r.Find(&rows)
	return
}

func (r Repository) GetProductByID(id uint) (row model.Article, err error) {
	err = r.Preload(clause.Associations).First(&row).Error
	return
}

func (r Repository) GetProductByPathname(pathname string) (row model.Article, err error) {
	err = r.First(&row, &model.Article{Pathname: pathname}).Error
	return
}
