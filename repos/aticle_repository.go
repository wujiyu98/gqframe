package repos

import (
	"github.com/wujiyu98/gqframe/database"
	"gorm.io/gorm"
)

var Article = &ArticleRepo{database.DB}

type ArticleRepo struct {
	*gorm.DB
}
