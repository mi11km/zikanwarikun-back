package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/mi11km/zikanwarikun-back/graph/generated"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models/users"
	"github.com/mi11km/zikanwarikun-back/internal/middleware/auth"
	"github.com/mi11km/zikanwarikun-back/pkg/jwt"
)

// todo emailの形式かどうかのバリデーションしてない
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (string, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return "", fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *model.DeleteUser) (bool, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	id, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateTimetable(ctx context.Context, input model.NewTimetable) (*model.Timetable, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Timetable{}, fmt.Errorf("access denied")
	}

	// todo ちゃんと実装する
	return &model.Timetable{
		ID:        strconv.FormatInt(1, 10),
		Name:      input.Name,
		Days:      5,
		Periods:   5,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		IsDefault: true,
		User:      &model.User{ID: user.ID, Email: user.Email, School: &user.School, Name: &user.Name}}, nil
}

func (r *mutationResolver) UpdateTimetable(ctx context.Context, input model.UpdateTimetable) (*model.Timetable, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Timetable{}, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTimetable(ctx context.Context, input string) (bool, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateClass(ctx context.Context, input model.NewClass) (*model.Class, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Class{}, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateClass(ctx context.Context, input model.UpdateClass) (*model.Class, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Class{}, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteClass(ctx context.Context, input string) (bool, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateClassTime(ctx context.Context, input model.UpdateClassTime) (*model.ClassTime, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.ClassTime{}, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, input string) (*model.User, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Timetable(ctx context.Context) (*model.Timetable, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Timetable{}, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Timetables(ctx context.Context) ([]*model.Timetable, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Classtimes(ctx context.Context, input string) ([]*model.ClassTime, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Classes(ctx context.Context, input string) ([]*model.Class, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
