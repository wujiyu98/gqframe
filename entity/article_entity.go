package entity

import "time"

// ArticleCategory
type ArticleCategory struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"not null"`
	Path      string `gorm:"not null;index"`
	ParentID  uint   `gorm:"not null;default:0"`
	Sort      uint   `gorm:"default:0"`
	Image     string `gorm:"default:''"`
	Summary   string `gorm:"size:1000;default:''"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}

// Article
type Article struct {
	ID                uint `gorm:"primarykey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArticleCategoryID uint   `gorm:"type:smallint;not null"`
	Title             string `gorm:"not null;uniqueIndex"`
	Path              string `gorm:"not null;index"`
	Sort              uint   `gorm:"default:0"`
	Showed            *byte  `gorm:"default:1;index"`
	Author            string `gorm:"size:60;default:''"`
	Summary           string `gorm:"size:1000;default:''"`
	Image             string `gorm:"default:''"`
}

type ArticelDetail struct {
	ID        uint   `gorm:"primarykey"`
	ArticleID uint   `gorm:"not null;index"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
	Content   string `gorm:"type:longtext"`
}
