package convert

import (
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
)

func ToGraphQLUser(dbUser *models.User) *model.User {
	graphUser := &model.User{
		ID:     dbUser.ID,
		Email:  dbUser.Email,
		Name:   dbUser.Name,
		School: dbUser.School,
	}
	graphTimetables := ToGraphQLTimetables(dbUser.Timetables)
	if len(graphTimetables) != 0 {
		graphUser.Timetables = graphTimetables
	}
	return graphUser
}
