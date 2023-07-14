package entity

// ArticleCategory
type ArticleCategory struct {
	Model
	Name        string `gorm:"not null"`
	Path        string `gorm:"not null;index"`
	ParentID    uint32 `gorm:"not null;default:0"`
	Sort        uint32 `gorm:"default:0"`
	Image       string `gorm:"default:''"`
	Description string `gorm:"size:1000"`
	Meta        Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}

// Article
type Article struct {
	Model
	ArticleCategoryID uint32 `gorm:"type:smallint;not null"`
	Name              string `gorm:"not null;uniqueIndex"`
	Path              string `gorm:"not null;index"`
	Sort              uint32 `gorm:"default:0"`
	Showed            *byte  `gorm:"default:1;index"`
	Author            string `gorm:"size:60;default:''"`
	Description       string `gorm:"size:1000"`
	Image             string `gorm:"default:''"`
}

type ArticleDetail struct {
	ID        uint32 `gorm:"primarykey"`
	ArticleID uint32 `gorm:"not null"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
	Content   string `gorm:"type:longtext"`
}
