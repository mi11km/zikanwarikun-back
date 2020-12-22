package graph

// todo テストコードを書いて本当に正しい動作をするか確認する
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
		log.Printf("action=create user, status=failed, err=%s", err)
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
		log.Printf("action=create user, status=failed, err=%s", result.Error)
		return "", result.Error
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Printf("action=create user, status=failed, err=%s", err)
		return "", err
	}

	log.Printf("action=create user, status=success")
	return token, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (string, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		err := &users.UnauthenticatedUserAccessError{}
		log.Printf("action=update user, status=failed, err=%s", err.Error())
		return "", err
	}

	updateData := make(map[string]interface{})
	if input.Email != nil {
		updateData["email"] = *input.Email
	}
	if input.School != nil {
		updateData["school"] = *input.School
	}
	if input.Name != nil {
		updateData["name"] = *input.Name
	}
	if input.Password != nil {
		if input.CurrentPassword == nil {
			log.Printf("action=update user, status=failed, err=currentPassword is not set")
			return "", fmt.Errorf("to update password, currentPassword is needed")
		}
		correct := users.CheckPasswordHash(*input.CurrentPassword, user.Password)
		if !correct {
			log.Printf("action=update user, status=failed, err=currentPassword is wrong")
			return "", fmt.Errorf("failed to update password. currentPassword is wrong")
		}

		hashedPassword, err := users.HashPassword(*input.Password)
		if err != nil {
			log.Printf("action=update user, status=failed, err=%s", err)
			return "", err
		}
		updateData["password"] = hashedPassword
	}
	if len(updateData) == 0 {
		log.Printf("action=update user, status=failed, err=update data is not set")
		return "", fmt.Errorf("update data must be set")
	}

	result := database.Db.Model(&user).Updates(updateData)
	if result.Error != nil {
		log.Printf("action=update user, status=failed, err=%s", result.Error)
		return "", result.Error
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Printf("action=update user, status=failed, err=%s", err)
		return "", err
	}

	log.Printf("action=update user, status=success")
	return token, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *model.DeleteUser) (bool, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		err := &users.UnauthenticatedUserAccessError{}
		log.Printf("action=delete user, status=failed, err=%s", err.Error())
		return false, err
	}

	correct := users.CheckPasswordHash(input.Password, user.Password)
	if !correct {
		log.Printf("action=delete user, status=failed, err=password is wrong")
		return false, fmt.Errorf("failed to delte user. password is wrong")
	}

	result := database.Db.Delete(&user)
	if result.Error != nil {
		log.Printf("action=delete user, status=failed, err=%s", result.Error)
		return false, result.Error
	}

	log.Printf("action=delete user, status=success")
	return true, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	result := database.Db.Select("id", "password").Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		log.Printf("action=login, status=failed, err=%s", result.Error)
		return "", result.Error
	}

	correct := users.CheckPasswordHash(input.Password, user.Password)
	if !correct {
		log.Printf("action=login, status=failed, err=email or password is wrong")
		return "", fmt.Errorf("failed to login. email or password is wrong")
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Printf("action=login, status=failed, err=%s", err)
		return "", err
	}

	log.Printf("action=login, status=success")
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	id, err := jwt.ParseToken(input.Token)
	if err != nil {
		log.Printf("action=refresh token, status=failed, err=failed to parse token")
		return "", fmt.Errorf("failed to parse given token")
	}
	token, err := jwt.GenerateToken(id)
	if err != nil {
		log.Printf("action=refresh token, status=failed, err=failed to generate token")
		return "", err
	}
	log.Printf("action=refresh token, status=success")
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
		err := &users.UnauthenticatedUserAccessError{}
		log.Printf("action=get current user data, status=failed, err=%s", err.Error())
		return nil, err
	}
	graphqlUser := &model.User{ // todo timetablesデータも送れるようにする
		ID:     user.ID,
		Email:  user.Email,
		School: &user.School,
		Name:   &user.Name,
	}
	log.Printf("action=get current user data, status=success")
	return graphqlUser, nil
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
