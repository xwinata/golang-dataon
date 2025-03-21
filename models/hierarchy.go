package models

import "gorm.io/gorm"

type Hierarchy struct {
	gorm.Model
	Value  string `gorm:"not null"`
	Upline *uint
}
