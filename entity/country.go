package entity

import "time"

// City
type City struct {
	ID          uint    `gorm:"primarykey"`
	Name        string  `gorm:"column:name;not null"`
	StateID     string  `gorm:"index:cities_test_ibfk_1;column:state_id;type:mediumint(8) unsigned;not null"`
	StateCode   string  `gorm:"column:state_code;not null"`
	CountryID   string  `gorm:"index:cities_test_ibfk_2;column:country_id;type:mediumint(8) unsigned;not null"`
	CountryCode string  `gorm:"column:country_code;type:char(2);not null"`
	Latitude    float64 `gorm:"column:latitude;type:decimal(10,8);not null"`
	Longitude   float64 `gorm:"column:longitude;type:decimal(11,8);not null"`
	Flag        bool    `gorm:"column:flag;type:tinyint(1);not null;default:1"`
	WikiDataID  string  `gorm:"column:wikiDataId;type:varchar(255)"`
}

// Countries [...]
type Country struct {
	ID             uint    `gorm:"primarykey"`
	Name           string  `gorm:"column:name;type:varchar(100);not null"`
	Iso3           string  `gorm:"column:iso3;type:char(3)"`
	NumericCode    string  `gorm:"column:numeric_code;type:char(3)"`
	Iso2           string  `gorm:"column:iso2;type:char(2)"`
	AreaCode       string  `gorm:"column:area_code;type:varchar(255)"`
	Capital        string  `gorm:"column:capital;type:varchar(255)"`
	Currency       string  `gorm:"column:currency;type:varchar(255)"`
	CurrencyName   string  `gorm:"column:currency_name;type:varchar(255)"`
	CurrencySymbol string  `gorm:"column:currency_symbol;type:varchar(255)"`
	Tld            string  `gorm:"column:tld;type:varchar(255)"`
	Native         string  `gorm:"column:native;type:varchar(255)"`
	Region         string  `gorm:"column:region;type:varchar(255)"`
	Subregion      string  `gorm:"column:subregion;type:varchar(255)"`
	Timezones      string  `gorm:"column:timezones;type:text"`
	Translations   string  `gorm:"column:translations;type:text"`
	Latitude       float64 `gorm:"column:latitude;type:decimal(10,8)"`
	Longitude      float64 `gorm:"column:longitude;type:decimal(11,8)"`
	Emoji          string  `gorm:"column:emoji;type:varchar(191)"`
	EmojiU         string  `gorm:"column:emojiU;type:varchar(191)"`
	Flag           bool    `gorm:"column:flag;type:tinyint(1);not null;default:1"`
	WikiDataID     string  `gorm:"column:wikiDataId;type:varchar(255)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// States [...]
type State struct {
	ID          uint    `gorm:"primarykey"`
	Name        string  `gorm:"column:name;not null"`
	CountryID   uint    `gorm:"index:country_region;column:country_id;type:mediumint(8) unsigned;not null"`
	CountryCode string  `gorm:"column:country_code;type:char(2);not null"`
	FipsCode    string  `gorm:"column:fips_code;type:varchar(255)"`
	Iso2        string  `gorm:"column:iso2;type:varchar(255)"`
	Type        string  `gorm:"column:type;type:varchar(191)"`
	Latitude    float64 `gorm:"column:latitude;type:decimal(10,8)"`
	Longitude   float64 `gorm:"column:longitude;type:decimal(11,8)"`
	Flag        bool    `gorm:"column:flag;type:tinyint(1);not null;default:1"`
	WikiDataID  string  `gorm:"column:wikiDataId;type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
