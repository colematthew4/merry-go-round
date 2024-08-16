package data

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Full  string `gorm:"unique;column:full_url"`
	Short string `gorm:"unique;column:short_url"`
}
