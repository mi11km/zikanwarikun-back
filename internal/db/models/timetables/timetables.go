package timetables

import (
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models/users"
)

type Timetable struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	ClassDays    int         `json:"class_days"`
	ClassPeriods int         `json:"class_periods"`
	IsDefault    bool        `json:"is_default"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	User         *users.User `json:"user"`
}

func (t *Timetable) CreateTimetable(input model.NewTimetable) (*model.Timetable, error) {
	return nil, nil
}

func (t *Timetable) UpdateTimetable(input model.UpdateTimetable) (*model.Timetable, error) {
	return nil, nil
}

func (t *Timetable) DeleteTimetable(input int) (bool, error) {
	return false, nil
}
