package convert

import (
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
)

func ToGraphQLTimetable(dbTimetable *models.Timetable) *model.Timetable {
	graphTimetable := &model.Timetable{
		ID:        strconv.Itoa(int(dbTimetable.ID)),
		Name:      dbTimetable.Name,
		Days:      dbTimetable.ClassDays,
		Periods:   dbTimetable.ClassPeriods,
		CreatedAt: dbTimetable.CreatedAt.String(),
		UpdatedAt: dbTimetable.UpdatedAt.String(),
		IsDefault: dbTimetable.IsDefault,
		//Classes: , todo
		//RowData: ,
	}
	if graphClassTimes :=ToGraphQLClassTimes(dbTimetable.ClassTimes); graphClassTimes != nil {
		graphTimetable.Classtimes = graphClassTimes
	}
	return graphTimetable
}

func ToGraphQLTimetables(dbTimetables []*models.Timetable) []*model.Timetable {
	if len(dbTimetables) == 0 {
		return nil
	}
	var graphTimetables []*model.Timetable
	for _, dt := range dbTimetables {
		gt := ToGraphQLTimetable(dt)
		graphTimetables = append(graphTimetables, gt)
	}
	return graphTimetables
}
