package repository

import (
	"github.com/wujiyu98/gqframe/database"
	"github.com/wujiyu98/gqframe/model"
	"gorm.io/gorm"
)

func New() Repository {
	return Repository{
		DB: database.DB,
	}
}

type Repository struct {
	*gorm.DB
}

func (r Repository) GetAtticles() (rows []model.Article) {
	r.Find(&rows)
	return
}

func (r Repository) GetAtticleByID(id uint) (row model.Article, err error) {
	err = r.First(&row).Error
	return
}

func (r Repository) GetAtticleByPathname(pathname string) (row model.Article, err error) {
	err = r.First(&row, &model.Article{Pathname: pathname}).Error
	return
}
