package classtimes

import (
	"time"

	"github.com/mi11km/zikanwarikun-back/graph/model"
)

type ClassTime struct {
	ID          int       `json:"id"`
	Period      int       `json:"period"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	TimetableID int       `json:"timetable_id"`
}

func (classTime *ClassTime) CreateClassTime(input model.NewClassTime) (*model.ClassTime, error) {
	return nil, nil
}
func (classTime *ClassTime) UpdateClassTime(input model.UpdateClassTime) (*model.ClassTime, error) {
	return nil, nil
}
