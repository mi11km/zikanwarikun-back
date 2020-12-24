package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mi11km/zikanwarikun-back/graph/generated"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models/users"
	"github.com/mi11km/zikanwarikun-back/internal/middleware/auth"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	auth := auth.ForContext(ctx)
	if auth != nil {
		err := &users.AuthenticateUserCanNotDoThisActionError{}
		log.Printf("action=update user, status=failed, err=%s", err.Error())
		return "", err
	}
	return r.UserService.CreateUser(input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (string, error) {
	auth := auth.ForContext(ctx)
	if auth.User == nil {
		err := &users.UnauthenticatedUserAccessError{}
		log.Printf("action=update user, status=failed, err=%s", err.Error())
		return "", err
	}
	return r.UserService.UpdateUser(&input, *auth.User)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUser) (bool, error) {
	auth := auth.ForContext(ctx)
	if auth.User == nil {
		err := &users.UnauthenticatedUserAccessError{}
		log.Printf("action=delete user, status=failed, err=%s", err.Error())
		return false, err
	}
	return r.UserService.DeleteUser(input, *auth.User)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	auth := auth.ForContext(ctx)
	if auth != nil {
		err := &users.AuthenticateUserCanNotDoThisActionError{}
		log.Printf("action=update user, status=failed, err=%s", err.Error())
		return "", err
	}
	return r.UserService.Login(input)
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (string, error) {
	auth := auth.ForContext(ctx)
	if auth.Token == nil {
		err := &users.TokenIsNotSetError{}
		log.Printf("action=refresh token, status=failed, err=%s", err)
		return "", err
	}
	return r.UserService.RefreshToken(*auth.Token)
}

func (r *mutationResolver) CreateTimetable(ctx context.Context, input model.NewTimetable) (*model.Timetable, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTimetable(ctx context.Context, input model.UpdateTimetable) (*model.Timetable, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTimetable(ctx context.Context, input string) (bool, error) {

	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateClassTime(ctx context.Context, input model.NewClassTime) (*model.ClassTime, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateClassTime(ctx context.Context, input model.UpdateClassTime) (*model.ClassTime, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateClass(ctx context.Context, input model.NewClass) (*model.Class, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateClass(ctx context.Context, input model.UpdateClass) (*model.Class, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteClass(ctx context.Context, input string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	auth := auth.ForContext(ctx)
	if auth.User == nil {
		err := &users.UnauthenticatedUserAccessError{}
		log.Printf("action=get current user data, status=failed, err=%s", err.Error())
		return nil, err
	}
	graphqlUser := &model.User{ // todo timetablesデータも送れるようにする
		ID:     auth.User.ID,
		Email:  auth.User.Email,
		School: &auth.User.School,
		Name:   &auth.User.Name,
	}
	log.Printf("action=get current user data, status=success")
	return graphqlUser, nil
}

func (r *queryResolver) Timetable(ctx context.Context) (*model.Timetable, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Timetables(ctx context.Context) ([]*model.Timetable, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
