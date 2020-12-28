package models

import (
	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	Attend  int
	Absent  int
	Late    int
	ClassID uint
}
