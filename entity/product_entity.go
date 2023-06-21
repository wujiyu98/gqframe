package entity

import (
	"time"
)

// Category
type Category struct {
	ID         uint        `gorm:"primarykey"`
	Name       string      `gorm:"not null"`
	Slug       string      `gorm:"default:''"`
	Path       string      `gorm:"not null;unique"`
	ParentID   uint        `gorm:"smallint;not null;default:0"`
	Sort       uint        `gorm:"smallint;default:0"`
	Quantity   uint        `gorm:"default:0"`
	Image      string      `gorm:"default:''"`
	Summary    string      `gorm:"type:text;"`
	Attributes []Attribute `gorm:"many2many:category_attriubtes"`
	Specs      []Attribute `gorm:"many2many:category_specs"`
	Meta
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Brand
type Brand struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"not null;unique"`
	Slug      string
	Letter    string `gorm:"size:1;default:''"`
	Path      string `gorm:"not null;index"`
	Sort      uint   `gorm:"default:0"`
	Quantity  uint   `gorm:"default:0"`
	Image     string `gorm:"default:''"`
	Summary   string `gorm:"type:text;"`
	Meta
}

type Attribute struct {
	ID   uint
	Name string
}

type Item struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	CategoryID uint
	BrandID    uint
	ImageID    uint
	SaleAble   byte
	Images     []Image `gorm:"many2many:item_images;"`
	Summary    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ItemAttribute struct {
	ID        uint
	ProductID uint
	Name      string
	Value     string
}

type ItemDetail struct {
	ID     uint `gorm:"primarykey"`
	ItemID uint
	Meta
	Description string
	Content     string
}

type Product struct {
	ID               uint `gorm:"primarykey"`
	Name             string
	ImageID          uint
	Sku              string
	Description      string
	Availability     byte
	AvailabilityDate uint
	Price            float64
	SalePrice        float64
	PriceUnit        string
	Currency         string
	QuantityLimit    string
	Quantity         uint
}

type ItemSpec struct {
	ProductID uint `gorm:"primarykey"`
	SpecID    uint `gorm:"primarykey"`
	Value     string
}

type ProductSpecOption struct {
	ProductID uint
	OptionID  uint
}
