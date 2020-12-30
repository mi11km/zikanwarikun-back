package convert

import (
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
)

func ToGraphQLClass(dbClass *models.Class) *model.Class {
	graphClass := &model.Class{
		ID:        strconv.Itoa(int(dbClass.ID)),
		Name:      dbClass.Name,
		Day:       dbClass.Days,
		Period:    dbClass.Periods,
		Style:     dbClass.Style,
		RoomOrURL: dbClass.RoomOrUrl,
		Teacher:   dbClass.Teacher,
		Credit:    &dbClass.Credit,
		Memo:      dbClass.Memo,
		Color:     dbClass.Color,
	}
	return graphClass
}

func ToGraphQLClasses(dbClasses []*models.Class) []*model.Class {
	if len(dbClasses) == 0 {
		return nil
	}
	var graphClasses []*model.Class
	for _, dc := range dbClasses {
		gc := ToGraphQLClass(dc)
		graphClasses = append(graphClasses, gc)
	}
	return graphClasses
}
