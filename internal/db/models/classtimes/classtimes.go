package classtimes

import "github.com/mi11km/zikanwarikun-back/internal/db/models/timetables"

type ClassTime struct {
	ID        int                `json:"id"`
	Period    int                   `json:"period"`
	StartTime string                `json:"start_time"`
	EndTime   string                `json:"end_time"`
	Timetable *timetables.Timetable `json:"timetable"`
}
