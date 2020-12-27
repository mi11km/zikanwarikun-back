package classes

import (
	"github.com/mi11km/zikanwarikun-back/graph/model"
)

type Class struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Day         int    `json:"day"`
	Periods     int    `json:"period"`
	Style       string `json:"style"`
	Color       string `json:"color"`
	Teacher     string `json:"teacher"`
	Credit      int    `json:"credit"`
	RoomOrUrl   string `json:"room_or_url"`
	Memo        string `json:"memo"`
	TimetableID int    `json:"timetable_id"`
}

func (class *Class) CreateClass(input model.NewClass) (*model.Class, error) {
	return nil, nil
}
func (class *Class) UpdateClass(input model.UpdateClass) (*model.Class, error) {
	return nil, nil
}
func (class *Class) DeleteClass(input string) (bool, error) {
	return false, nil
}
