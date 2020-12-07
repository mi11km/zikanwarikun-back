package classes

import "github.com/mi11km/zikanwarikun-back/internal/db/models/timetables"

type Class struct {
	ID        int                `json:"id"`
	Name      string                `json:"name"`
	Day       int                   `json:"day"`
	Period    int                   `json:"period"`
	Style     string                `json:"style"`
	Color     string                `json:"color"`
	Teacher   string                `json:"teacher"`
	Credit    int                   `json:"credit"`
	RoomOrUrl string                `json:"room_or_url"`
	Memo      string                `json:"memo"`
	Timetable *timetables.Timetable `json:"timetable"`
}
