package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Name    string
	Url     string
	ClassID uint
}
