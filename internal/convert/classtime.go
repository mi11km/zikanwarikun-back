package convert

import (
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
)

func ToGraphQLClassTime(dbClassTime *models.ClassTime) *model.ClassTime {
	graphClassTime := &model.ClassTime{
		ID:        strconv.Itoa(int(dbClassTime.ID)),
		Period:    dbClassTime.Periods,
		StartTime: dbClassTime.StartTime,
		EndTime:   dbClassTime.EndTime,
	}
	return graphClassTime
}

func ToGraphQLClassTimes(dbClassTimes []*models.ClassTime) []*model.ClassTime {
	if len(dbClassTimes) == 0 {
		return nil
	}
	var graphClassTimes []*model.ClassTime
	for _, dct := range dbClassTimes {
		gct := ToGraphQLClassTime(dct)
		graphClassTimes = append(graphClassTimes, gct)
	}
	return graphClassTimes
}
