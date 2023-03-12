package entity

import (
	"time"

	"gorm.io/gorm"
)

type Meta struct {
	Title       string
	Keywords    string
	Discription string
}

type Language struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:60;not null"`
	Domain    string `gorm:"size:120;not null"`
	Code      string `gorm:"size:10;not null"`
	Image     string `gorm:"size:255;default:''"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Site
type Site struct {
	ID          uint   `gorm:"primarykey"`
	LanguageID  uint   `gorm:"type:int(8);not null"`
	SiteName    string `gorm:"default:''"`
	Email       string `gorm:"default:''"`
	Company     string `gorm:"default:''"`
	Contact     string `gorm:"default:''"`
	Phone       string `gorm:"size:20;default:''"`
	Phone2      string `gorm:"size:20;default:''"`
	MobilePhone string `gorm:"size:20;default:''"`
	Skype       string `gorm:"size:60;default:''"`
	QQ          string `gorm:"size:20;default:''"`
	Whatsapp    string `gorm:"size:20;default:''"`
	Address     string `gorm:"default:''"`
	Summary     string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Banner
type Banner struct {
	ID         uint   `gorm:"primarykey"`
	LanguageID uint   `gorm:"smallint;not null"`
	Name       string `gorm:"default:''"`
	Image      string `gorm:"default:''"`
	Href       string `gorm:"default:''"`
	Title      string `gorm:"default:''"`
	Summary    string `gorm:"default:''"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Seo struct {
	ID         uint   `gorm:"primarykey"`
	LanguageID uint   `gorm:"smallint;not null"`
	Name       string `gorm:"size:255"`
	Path       string `gorm:"size:255"`
	Meta       Meta   `gorm:"embedded;embeddedPrefix:meta_"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// ArticleCategory
type ArticleCategory struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Pathname  string `gorm:"not null;uniqueIndex"`
	ParentID  uint   `gorm:"not null;default:0"`
	SortOrder uint   `gorm:"default:0"`
	Image     string `gorm:"default:''"`
	Summary   string `gorm:"type:text;"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
	Type      string `gorm:"type:enum('about','info','article');default:'article'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Article
type Article struct {
	ID                uint   `gorm:"primarykey"`
	LanguageID        uint   `gorm:"type:smallint;not null"`
	ArticleCategoryID uint   `gorm:"type:smallint;not null"`
	Title             string `gorm:"not null"`
	Pathname          string `gorm:"not null;uniqueIndex"`
	SortOrder         uint   `gorm:"default:0"`
	Showed            *uint  `gorm:"type:tinyint;default:1;index"`
	View              uint   `gorm:"tinyint;default:0"`
	Summary           string `gorm:"type:text;"`
	Image             string `gorm:"default:''"`
	Author            string `gorm:"size:60;default:''"`
	Content           string `gorm:"type:longtext"`
	Meta              Meta   `gorm:"embedded;embeddedPrefix:meta_"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// Category
type Category struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Pathname  string `gorm:"not null;unique"`
	ParentID  uint   `gorm:"smallint;not null;default:0"`
	SortOrder uint   `gorm:"smallint;default:0"`
	Qty       uint   `gorm:"default:0"`
	Image     string `gorm:"default:''"`
	Summary   string `gorm:"type:text;"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Product
type Product struct {
	ID             uint                `gorm:"primarykey"`
	LanguageID     uint                `gorm:"type:smallint;not null"`
	CategoryID     uint                `gorm:"type:smallint;not null"`
	ManufacturerID uint                `gorm:"type:smallint;not null"`
	Title          string              `gorm:"not null;unique"`
	Pathname       string              `gorm:"not null;unique"`
	SortOrder      uint                `gorm:"default:0"`
	Showed         *byte               `gorm:"default:1"`
	Summary        string              `gorm:"type:text;"`
	Image          string              `gorm:"default:''"`
	Images         []string            `gorm:"type:json;serializer:json"`
	Content        string              `gorm:"type:longtext"`
	Stock          uint                `gorm:"default:0"`
	Hot            byte                `gorm:"default:0"`
	New            byte                `gorm:"default:0"`
	Special        byte                `gorm:"default:0"`
	Price          float64             `gorm:"type:decimal(10,4)"`
	Prices         []map[string]string `gorm:"type:json;serializer:json"`
	Meta           Meta                `gorm:"embedded;embeddedPrefix:meta_"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Attribute
type Attribute struct {
	ID         uint `gorm:"primarykey;"`
	CategoryID uint
	Name       string
}

// ProductAttribute
type ProductAttribute struct {
	ID        uint   `gorm:"primarykey;"`
	ProductID uint   `gorm:"not null"`
	Key       string `gorm:"not null"`
	Value     string `gorm:"default:''"`
}
