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
	if graphTimetables := ToGraphQLTimetables(dbUser.Timetables); graphTimetables != nil {
		graphUser.Timetables = graphTimetables
	}
	return graphUser
}
