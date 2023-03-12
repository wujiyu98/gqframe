package model

import "time"

// User
type User struct {
	ID              uint   `gorm:"primarykey"`
	Email           string `gorm:"type:varchar(255)"`
	Name            string `gorm:"type:varchar(60)"`
	Country         string `gorm:"type:varchar(60)"`
	Company         string `gorm:"type:varchar(255)"`
	Password        string `gorm:"type:varchar(255)"`
	RememberToken   string `gorm:"type:varchar(255)"`
	MobilePhone     string `gorm:"type:varchar(20)"`
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Address
type Address struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `gorm:"type:int(11);not null"`
	CountryID uint   `gorm:"type:int(11);not null"`
	StateID   uint   `gorm:"type:int(11);not null"`
	CityID    uint   `gorm:"type:int(11);not null"`
	Firstname string `gorm:"type:varchar(20);not null"`
	Lastname  string `gorm:"type:varchar(20);not null"`
	Company   string `gorm:"type:varchar(255);default:''"`
	Postcode  string `gorm:"type:varchar(30);default:''"`
	Address1  string `gorm:"type:varchar(255);default:''"`
	Address2  string `gorm:"type:varchar(255);default:''"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Message struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"varchar(60);not null"`
	Email       string `gorm:"varchar(60);not null"`
	Country     string `gorm:"varchar(60);not null"`
	AreaCode    string `gorm:"varchar(60);not null"`
	MobilePhone string `gorm:"varchar(30);not null"`
	Company     string `gorm:"varchar(120);default:''"`
	Comment     string `gorm:"varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Enquiry struct {
	ID          uint          `gorm:"primarykey"`
	Name        string        `gorm:"varchar(60);not null"`
	Email       string        `gorm:"varchar(60);not null"`
	Country     string        `gorm:"varchar(60);not null"`
	AreaCode    string        `gorm:"varchar(30);not null"`
	MobilePhone string        `gorm:"varchar(30);not null"`
	Company     string        `gorm:"varchar(120);default:''"`
	Comment     string        `gorm:"varchar(255);not null"`
	Items       []EnquiryItem `gorm:"type:json"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type EnquiryItem struct {
	EnquiryID    uint    `gorm:"not null;index"`
	Title        string  `gorm:"type:varchar(255);not null"`
	Manufacturer string  `gorm:"type:varchar(255);not null;default:''"`
	Summary      string  `gorm:"type:varchar(255);not null;default:''"`
	Qty          uint    `gorm:"type:int(11);not null;default:0"`
	Price        float64 `gorm:"decimal(10,2);not null;default:0"`
}

type Order struct {
	ID                uint    `gorm:"primarykey"`
	UserID            uint    `gorm:"type:int(11);not null"`
	OrderNumber       string  `gorm:"type:char(15);not null"`
	Address           string  `gorm:"size:500;not null"`
	TransactionNumber string  `gorm:"size:255;not null;default ''"`
	PaypalFee         float64 `gorm:"type:decimal(10,2);not null;default:0"`
	Total             float64 `gorm:"type:decimal(10,2);not null"`
	Payment           byte    `gorm:"not null;default:0"`
	Status            byte    `gorm:"not null;default:0"`
	Freight           float64 `gorm:"type:decimal(10,2);not null;default:35"`
	Express           string  `gorm:"size:255;not null;default:''"`
	TrackingNumber    string  `gorm:"size:255;not null;default:''"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type OrderItem struct {
	OrderID      uint    `gorm:"not null;index"`
	ProductID    uint    `gorm:"type:int(11);not null"`
	ProductUrl   string  `gorm:"not null;default:''"`
	Title        string  `gorm:"type:varchar(255);not null"`
	Image        string  `gorm:"type:varchar(255);not null;default:''"`
	Manufacturer string  `gorm:"type:varchar(255);not null;default:''"`
	Summary      string  `gorm:"type:varchar(255);not null;default:''"`
	Qty          uint    `gorm:"type:int(11);not null"`
	Price        float64 `gorm:"decimal(10,2);not null"`
}
