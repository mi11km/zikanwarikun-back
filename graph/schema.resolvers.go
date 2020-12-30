package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/mi11km/zikanwarikun-back/graph/generated"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/convert"
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
	err := dbUser.Create(input)
	if err != nil {
		log.Printf("action=signup, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(dbUser.ID)
	if err != nil {
		log.Printf("action=signup, status=failed, err=%s", err)
		return nil, err
	}
	graphUser := convert.ToGraphQLUser(dbUser)
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
	err := auth.User.Update(&input)
	if err != nil {
		log.Printf("action=update login user, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(auth.User.ID)
	if err != nil {
		log.Printf("action=update login user, status=failed, err=%s", err)
		return nil, err
	}
	graphUser := convert.ToGraphQLUser(auth.User)
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
	return auth.User.Delete(input)
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
	graphUser := convert.ToGraphQLUser(dbUser)
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
	dbTimetable := &models.Timetable{}
	err := dbTimetable.Create(input, *auth.User)
	if err != nil {
		log.Printf("action=create timetable, status=failed, err=%s", err)
		return nil, err
	}
	graphTimetable := convert.ToGraphQLTimetable(dbTimetable)
	return graphTimetable, nil
}

func (r *mutationResolver) UpdateTimetable(ctx context.Context, input model.UpdateTimetable) (*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update timetable, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbTimetable := models.FetchTimetableById(input.ID)
	err := dbTimetable.Update(input, *auth.User)
	if err != nil {
		log.Printf("action=update timetable, status=failed, err=%s", err)
		return nil, err
	}
	graphTimetable := convert.ToGraphQLTimetable(dbTimetable)
	return graphTimetable, nil
}

func (r *mutationResolver) DeleteTimetable(ctx context.Context, input string) (bool, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=delete timetable, status=failed, err=%s", err.Error())
		return false, err
	}
	dbTimetable := &models.Timetable{}
	return dbTimetable.Delete(input)
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
	dbClassTime := &models.ClassTime{}
	if err := dbClassTime.Create(input); err != nil {
		log.Printf("action=create class time, status=failed, err=%s", err)
		return nil, err
	}
	return convert.ToGraphQLClassTime(dbClassTime), nil
}

func (r *mutationResolver) UpdateClassTime(ctx context.Context, input model.UpdateClassTime) (*model.ClassTime, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update class time, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbClassTime := models.FetchClassTimeById(input.ID)
	if err := dbClassTime.Update(input); err != nil {
		log.Printf("action=update class time, status=failed, err=%s", err.Error())
		return nil, err
	}
	return convert.ToGraphQLClassTime(dbClassTime), nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=get login user data, status=failed, err=%s", err.Error())
		return nil, err
	}
	timetables, err := models.FetchTimetablesByUser(*auth.User)
	if err != nil {
		log.Printf("action=get login user data, status=failed, err=%s", err)
		return nil, err
	}
	auth.User.Timetables = timetables
	graphUser := convert.ToGraphQLUser(auth.User)
	log.Printf("action=get login user data, status=success")
	return graphUser, nil
}

func (r *queryResolver) Timetable(ctx context.Context) (*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=get default timetable of login user, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbDefaultTimetable, err := models.FetchDefaultTimetableByUser(*auth.User)
	if err != nil {
		log.Printf("action=get default timetable of login user, status=failed, err=%s", err)
		return nil, err
	}
	graphTimetable := convert.ToGraphQLTimetable(dbDefaultTimetable)
	return graphTimetable, nil
}

func (r *queryResolver) Timetables(ctx context.Context) ([]*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=get all timetables of login user, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbTimetables, err := models.FetchTimetablesByUser(*auth.User)
	if err != nil {
		log.Printf("action=get all timetables of login user, status=failed, err=%s", err.Error())
		return nil, err
	}
	graphTimetables := convert.ToGraphQLTimetables(dbTimetables)
	return graphTimetables, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
