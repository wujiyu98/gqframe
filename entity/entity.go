package entity

import "time"

type Meta struct {
	Title       string `gorm:"default:''"`
	Keywords    string `gorm:"default:''"`
	Discription string `gorm:"default:''"`
}

type Price struct {
	Min      uint
	Value    uint
	Currency string
}
type Image struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Url         string
	Description string
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Language struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"size:60;not null"`
	LocalName string `gorm:"size:120;not null"`
	Domain    string `gorm:"size:120;not null"`
	Code      string `gorm:"size:10;not null"`
	Image     string `gorm:"size:255;default:''"`
}

// WebSite
type WebSite struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
}

// Banner
type Banner struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"default:''"`
	Image     string `gorm:"default:''"`
	Url       string `gorm:"default:''"`
	Title     string `gorm:"default:''"`
	Summary   string `gorm:"default:''"`
	Color     string `gorm:"default:''"`
	Position  string `gorm:"default:''"`
}

type Seo struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"size:255"`
	Path      string `gorm:"size:255"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}
