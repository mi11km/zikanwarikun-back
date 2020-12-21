package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mi11km/zikanwarikun-back/graph/generated"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"github.com/mi11km/zikanwarikun-back/internal/db/models/users"
	"github.com/mi11km/zikanwarikun-back/internal/middleware/auth"
	"github.com/mi11km/zikanwarikun-back/pkg/jwt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	hashedPassword, err := users.HashPassword(input.Password)
	if err != nil {
		log.Printf("action=failed to generate hashpassword, err=%s", err)
		return "", err
	}

	user := users.User{
		ID:       uuid.New().String(), // todo uuidのDBへの保存方法の最適化(現在は36文字のVARCHAR)
		Email:    input.Email,
		Password: hashedPassword,
		School:   input.School,
		Name:     input.Name,
	}
	result := database.Db.Create(&user)
	if result.Error != nil {
		log.Printf("action=failed to create user, err=%s", result.Error)
		return "", result.Error
	}

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
	result := database.Db.Select("id", "password").Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		log.Printf("action=failed to select user from email, err=%s", result.Error)
		return "", result.Error
	}

	correct := users.CheckPasswordHash(input.Password, user.Password)
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

func (r *mutationResolver) CreateClassTime(ctx context.Context, input model.NewClassTime) (*model.ClassTime, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateClassTime(ctx context.Context, input model.UpdateClassTime) (*model.ClassTime, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.ClassTime{}, fmt.Errorf("access denied")
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

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
