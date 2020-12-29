package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/mi11km/zikanwarikun-back/graph/generated"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
	"github.com/mi11km/zikanwarikun-back/internal/middleware/auth"
	"github.com/mi11km/zikanwarikun-back/internal/myerrors"
	"github.com/mi11km/zikanwarikun-back/pkg/jwt"
)

func (r *mutationResolver) Signup(ctx context.Context, input model.NewUser) (*model.Auth, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth != nil {
		err := &myerrors.AuthenticateUserCanNotDoThisActionError{}
		log.Printf("action=signup, status=failed, err=%s", err.Error())
		return nil, err
	}

	dbUser := &models.User{}
	err := dbUser.CreateUser(input)
	if err != nil {
		log.Printf("action=signup, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(dbUser.ID)
	if err != nil {
		log.Printf("action=signup, status=failed, err=%s", err)
		return nil, err
	}
	graphUser := &model.User{
		ID: dbUser.ID, Email: dbUser.Email, Name: dbUser.Name, School: dbUser.School}
	log.Printf("action=signup, status=success")
	return &model.Auth{User: graphUser, Token: token}, nil
}

func (r *mutationResolver) UpdateLoginUser(ctx context.Context, input model.UpdateUser) (*model.Auth, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update login user, status=failed, err=%s", err.Error())
		return nil, err
	}
	err := auth.User.UpdateLoginUser(&input)
	if err != nil {
		log.Printf("action=update login user, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(auth.User.ID)
	if err != nil {
		log.Printf("action=update login user, status=failed, err=%s", err)
		return nil, err
	}
	graphUser := &model.User{
		ID: auth.User.ID, Email: auth.User.Email, Name: auth.User.Name, School: auth.User.School}
	log.Printf("action=update login user, status=success")
	return &model.Auth{User: graphUser, Token: token}, nil
}

func (r *mutationResolver) DeleteLoginUser(ctx context.Context, input model.DeleteUser) (bool, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=delete user, status=failed, err=%s", err.Error())
		return false, err
	}
	return auth.User.DeleteLoginUser(input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.Auth, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth != nil {
		err := &myerrors.AuthenticateUserCanNotDoThisActionError{}
		log.Printf("action=login, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbUser := &models.User{}
	err := dbUser.Login(input)
	if err != nil {
		log.Printf("action=login, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(dbUser.ID)
	if err != nil {
		log.Printf("action=login, status=failed, err=%s", err)
		return nil, err
	}
	graphUser := &model.User{
		ID: dbUser.ID, Email: dbUser.Email, Name: dbUser.Name, School: dbUser.School} // todo timetablesも入れる
	log.Printf("action=login, status=success")
	return &model.Auth{User: graphUser, Token: token}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (string, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=refresh token, status=failed, err=%s", err.Error())
		return "", err
	}
	return jwt.RefreshToken(*auth.Token)
}

func (r *mutationResolver) CreateTimetable(ctx context.Context, input model.NewTimetable) (*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=create timetable, status=failed, err=%s", err.Error())
		return nil, err
	}
	panic("not implement")
}

func (r *mutationResolver) UpdateTimetable(ctx context.Context, input model.UpdateTimetable) (*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update timetable, status=failed, err=%s", err.Error())
		return nil, err
	}
	panic("not implement")
}

func (r *mutationResolver) DeleteTimetable(ctx context.Context, input string) (bool, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=delete timetable, status=failed, err=%s", err.Error())
		return false, err
	}
	panic("not implement")
}

func (r *mutationResolver) CreateClass(ctx context.Context, input model.NewClass) (*model.Class, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=create class, status=failed, err=%s", err.Error())
		return nil, err
	}
	panic("not implement")
}

func (r *mutationResolver) UpdateClass(ctx context.Context, input model.UpdateClass) (*model.Class, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update class, status=failed, err=%s", err.Error())
		return nil, err
	}
	panic("not implement")
}

func (r *mutationResolver) DeleteClass(ctx context.Context, input string) (bool, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=delete class, status=failed, err=%s", err.Error())
		return false, err
	}
	panic("not implement")
}

func (r *mutationResolver) CreateClassTime(ctx context.Context, input model.NewClassTime) (*model.ClassTime, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=create class time, status=failed, err=%s", err.Error())
		return nil, err
	}
	panic("not implement")
}

func (r *mutationResolver) UpdateClassTime(ctx context.Context, input model.UpdateClassTime) (*model.ClassTime, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update class time, status=failed, err=%s", err.Error())
		return nil, err
	}
	panic("not implement")
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=get current user data, status=failed, err=%s", err.Error())
		return nil, err
	}
	graphqlUser := &model.User{
		ID:     auth.User.ID,
		Email:  auth.User.Email,
		School: auth.User.School,
		Name:   auth.User.Name,
	}
	dbTimetables, err := models.FetchTimetablesByUserId(auth.User.ID)
	if err != nil {
		log.Printf("action=get current user data, status=failed, err=%s", err)
		return nil, err
	}
	graphTimetables := models.ConvertTimetablesFromDbToGraph(dbTimetables, graphqlUser)
	graphqlUser.Timetables = graphTimetables
	log.Printf("action=get current user data, status=success")
	return graphqlUser, nil
}

func (r *queryResolver) Timetable(ctx context.Context) (*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=get default timetable, status=failed, err=%s", err.Error())
		return nil, err
	}
	graphUser := &model.User{
		ID:     auth.User.ID,
		Email:  auth.User.Email,
		School: auth.User.School,
		Name:   auth.User.Name,
	}
	dbTimetables, err := models.FetchDefaultTimetableByUserId(auth.User.ID)
	if err != nil {
		log.Printf("action=get default timetable, status=failed, err=%s", err)
		return nil, err
	}
	graphTimetable := models.ConvertTimetableFromDbToGraph(dbTimetables, graphUser)
	log.Printf("action=get default timetable, status=success")
	return graphTimetable, nil
}

func (r *queryResolver) Timetables(ctx context.Context) ([]*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=get all timetables, status=failed, err=%s", err.Error())
		return nil, err
	}
	graphUser := &model.User{
		ID:     auth.User.ID,
		Email:  auth.User.Email,
		School: auth.User.School,
		Name:   auth.User.Name,
	}
	dbTimetables, err := models.FetchTimetablesByUserId(auth.User.ID)
	if err != nil {
		log.Printf("action=get all timetables, status=failed, err=%s", err)
		return nil, err
	}
	graphTimetables := models.ConvertTimetablesFromDbToGraph(dbTimetables, graphUser)
	log.Printf("action=get all timetables, status=success")
	return graphTimetables, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
