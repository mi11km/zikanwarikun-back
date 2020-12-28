package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Kind       string
	Deadline   time.Time
	IsDone     bool
	Memo       *string
	IsRepeated bool
	ClassID    uint
}
