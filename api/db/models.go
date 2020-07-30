package db

import "github.com/jinzhu/gorm"

type Testament struct {
	gorm.Model
	Id   int    `gorm:"PRIMARY_KEY"`
	Name string `gorm:"size:45"`
}

type Books struct {
	gorm.Model
	Id        int    `gorm:"PRIMARY_KEY"`
	Name      string `gorm:"size:45"`
	Abbrev    string `gorm:"size:5"`
	Testament string `gorm:"size:5"`
}

type Verses struct {
	gorm.Model
	Id        int    `gorm:"PRIMARY_KEY"`
	Version   string `gorm:"size:10"`
	Text      string `gorm:"type:text"`
	Testament int
	Book      int
	Verses    int
}
