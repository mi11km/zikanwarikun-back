package models

import (
	"time"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	"gorm.io/gorm"
)

type ClassTime struct {
	gorm.Model
	Periods     uint
	StartTime   time.Time
	EndTime     time.Time
	TimetableID uint
}

func (classTime *ClassTime) CreateClassTime(input model.NewClassTime) (*model.ClassTime, error) {
	return nil, nil
}
func (classTime *ClassTime) UpdateClassTime(input model.UpdateClassTime) (*model.ClassTime, error) {
	return nil, nil
}
