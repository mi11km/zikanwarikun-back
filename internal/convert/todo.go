package convert

import (
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
	parsetime "github.com/mi11km/zikanwarikun-back/pkg/time"
)

func ToGraphQlTodo(dbTodo *models.Todo) *model.Todo {
	graphTodo := &model.Todo{
		ID:         strconv.Itoa(int(dbTodo.ID)),
		Kind:       dbTodo.Kind,
		Deadline:   parsetime.TimeToString(dbTodo.Deadline),
		IsDone:     dbTodo.IsDone,
		Memo:       dbTodo.Memo,
		IsRepeated: dbTodo.IsRepeated,
	}
	return graphTodo
}

func ToGraphQLTodos(dbTodos []*models.Todo) []*model.Todo {
	if len(dbTodos) == 0 {
		return nil
	}
	var graphTodos []*model.Todo
	for _, dt := range dbTodos {
		gt := ToGraphQlTodo(dt)
		graphTodos = append(graphTodos, gt)
	}
	return graphTodos
}
