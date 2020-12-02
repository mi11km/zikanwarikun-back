package timetables

import "github.com/mi11km/zikanwarikun-back/internal/db/models/users"

type Timetable struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	ClassDays    int         `json:"class_days"`
	ClassPeriods int         `json:"class_periods"`
	IsDefault    bool        `json:"is_default"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	User         *users.User `json:"user"`
}
