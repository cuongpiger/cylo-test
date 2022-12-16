package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID     int64          `json:"id" gorm:"column:id;primaryKey;autoIncrement;"`
	Name   string         `json:"name" gorm:"column:name;not null;type:string;"`
	Price  float64        `json:"price" gorm:"column:price;not null;type:float;"`
	Images pq.StringArray `json:"images" gorm:"column:images;type:text[];"`
}
