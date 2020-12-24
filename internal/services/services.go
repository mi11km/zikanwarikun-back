package services

import (
	"context"

	"github.com/mi11km/zikanwarikun-back/graph/model"
)

type UserService interface {
	CreateUser(input model.NewUser) (string, error)
	UpdateUser(input model.UpdateUser, ctx context.Context) (string, error)
	DeleteUser(input model.DeleteUser, ctx context.Context) (bool, error)
	Login(input model.Login) (string, error)
	RefreshToken(input model.RefreshTokenInput) (string, error)
}
