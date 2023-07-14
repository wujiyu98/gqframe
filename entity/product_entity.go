package entity

// Category
type Category struct {
	Model
	Name        string      `gorm:"not null"`
	Path        string      `gorm:"not null;unique"`
	ParentID    uint32      `gorm:"smallint;not null;default:0"`
	Sort        uint32      `gorm:"smallint;default:0"`
	Quantity    uint32      `gorm:"default:0"`
	Image       string      `gorm:"default:''"`
	Description string      `gorm:"type:text;"`
	Attributes  []Attribute `gorm:"many2many:category_attriubtes"`
	Specs       []Attribute `gorm:"many2many:category_specs"`
	Meta
}

// Brand
type Brand struct {
	Model
	Name        string `gorm:"not null;unique"`
	Letter      string `gorm:"size:1;default:''"`
	Path        string `gorm:"not null;index"`
	Sort        uint32 `gorm:"default:0"`
	Quantity    uint32 `gorm:"default:0"`
	Image       string `gorm:"default:''"`
	SameAs      string `gorm:"default:''"`
	Description string `gorm:"type:text;"`
	Meta
}

type Attribute struct {
	Model
	Name        string
	Image       string
	Description string
	SameAs      string
}

type Option struct {
	Model
	Name        string
	Image       string
	Description string
	SameAs      string
}

type Item struct {
	Model
	Name        string
	CategoryID  uint32
	BrandID     uint32
	ImageID     uint32
	Showed      byte
	Quantity    uint32
	Price       float64
	Images      []Image `gorm:"many2many:item_images;"`
	Description string
}

type ItemDetail struct {
	ID      uint32
	ItemID  uint32
	Section string `gorm:"type:json"`
}

type Product struct {
	ID               uint `gorm:"primarykey"`
	Name             string
	ImageID          uint
	Sku              string
	Description      string
	Availability     byte
	AvailabilityDate uint
	Quantity         uint
	Price            float64
	Currency         string
	SalePrice        float64
	Unit             string
	Min              uint
	Max              uint
}
