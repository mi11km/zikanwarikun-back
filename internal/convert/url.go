package convert

import (
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
)

func ToGraphQlUrl(dbUrl *models.Url) *model.URL {
	graphUrl := &model.URL{
		ID:   strconv.Itoa(int(dbUrl.ID)),
		Name: dbUrl.Name,
		URL:  dbUrl.Url,
	}
	return graphUrl
}

func ToGraphQLUrls(dbUrls []*models.Url) []*model.URL {
	if len(dbUrls) == 0 {
		return nil
	}
	var graphUrls []*model.URL
	for _, dt := range dbUrls {
		gt := ToGraphQlUrl(dt)
		graphUrls = append(graphUrls, gt)
	}
	return graphUrls
}