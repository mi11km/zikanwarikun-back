package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
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
	if err := dbUser.Create(input); err != nil {
		log.Printf("action=signup, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(dbUser.ID)
	if err != nil {
		log.Printf("action=signup, status=failed, err=%s", err)
		return nil, err
	}
	log.Printf("action=signup, status=success")
	return &model.Auth{User: convert.ToGraphQLUser(dbUser), Token: token}, nil
}

func (r *mutationResolver) UpdateLoginUser(ctx context.Context, input model.UpdateUser) (*model.Auth, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update login user, status=failed, err=%s", err.Error())
		return nil, err
	}
	if err := auth.User.Update(&input); err != nil {
		log.Printf("action=update login user, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(auth.User.ID)
	if err != nil {
		log.Printf("action=update login user, status=failed, err=%s", err)
		return nil, err
	}
	log.Printf("action=update login user, status=success")
	return &model.Auth{User: convert.ToGraphQLUser(auth.User), Token: token}, nil
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
	if err := dbUser.Login(input); err != nil {
		log.Printf("action=login, status=failed, err=%s", err)
		return nil, err
	}
	token, err := jwt.GenerateToken(dbUser.ID)
	if err != nil {
		log.Printf("action=login, status=failed, err=%s", err)
		return nil, err
	}
	log.Printf("action=login, status=success")
	return &model.Auth{User: convert.ToGraphQLUser(dbUser), Token: token}, nil
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
	if err := dbTimetable.Create(input, *auth.User); err != nil {
		log.Printf("action=create timetable, status=failed, err=%s", err)
		return nil, err
	}
	return convert.ToGraphQLTimetable(dbTimetable), nil
}

func (r *mutationResolver) UpdateTimetable(ctx context.Context, input model.UpdateTimetable) (*model.Timetable, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update timetable, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbTimetable := models.FetchTimetableById(input.ID)
	if err := dbTimetable.Update(input, *auth.User); err != nil {
		log.Printf("action=update timetable, status=failed, err=%s", err)
		return nil, err
	}
	return convert.ToGraphQLTimetable(dbTimetable), nil
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
	dbClass := &models.Class{}
	if err := dbClass.Create(input); err != nil {
		log.Printf("action=create class, status=failed, err=%s", err)
		return nil, err
	}
	return convert.ToGraphQLClass(dbClass), nil
}

func (r *mutationResolver) UpdateClass(ctx context.Context, input model.UpdateClass) (*model.Class, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update class, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbClass := models.FetchClassById(input.ID)
	if err := dbClass.Update(input); err != nil {
		log.Printf("action=update class, status=failed, err=%s", err)
		return nil, err
	}
	return convert.ToGraphQLClass(dbClass), nil
}

func (r *mutationResolver) DeleteClass(ctx context.Context, input string) (bool, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=delete class, status=failed, err=%s", err.Error())
		return false, err
	}
	dbClass := &models.Class{}
	return dbClass.Delete(input)
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

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=create todo, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbTodo := &models.Todo{}
	if err := dbTodo.Create(input); err != nil {
		log.Printf("action=create todo, status=failed, err=%s", err)
		return nil, err
	}
	return convert.ToGraphQlTodo(dbTodo), nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update todo, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbTodo := models.FetchTodoById(input.ID)
	if err := dbTodo.Update(input); err != nil {
		log.Printf("action=update todo, status=failed, err=%s", err.Error())
		return nil, err
	}
	return convert.ToGraphQlTodo(dbTodo), nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input string) (bool, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=delete todo, status=failed, err=%s", err.Error())
		return false, err
	}
	dbTodo := &models.Todo{}
	return dbTodo.Delete(input)
}

func (r *mutationResolver) CreateURL(ctx context.Context, input model.NewURL) (*model.URL, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=create url, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbUrl := &models.Url{}
	if err := dbUrl.Create(input); err != nil {
		log.Printf("action=create url, status=failed, err=%s", err)
		return nil, err
	}
	return convert.ToGraphQlUrl(dbUrl), nil
}

func (r *mutationResolver) UpdateURL(ctx context.Context, input model.UpdateURL) (*model.URL, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=update url, status=failed, err=%s", err.Error())
		return nil, err
	}
	dbUrl := models.FetchUrlById(input.ID)
	if err := dbUrl.Update(input); err != nil {
		log.Printf("action=update url, status=failed, err=%s", err.Error())
		return nil, err
	}
	return convert.ToGraphQlUrl(dbUrl), nil
}

func (r *mutationResolver) DeleteURL(ctx context.Context, input string) (bool, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=delete url, status=failed, err=%s", err.Error())
		return false, err
	}
	dbUrl := &models.Url{}
	return dbUrl.Delete(input)
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
	models.SetClassTimesToEachTimetable(timetables)
	models.SetClassesToEachTimetable(timetables)
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
	// todo classやclass_timeをセットしてない　というかそもそもこのAPIつかう？
	return convert.ToGraphQLTimetable(dbDefaultTimetable), nil
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
	models.SetClassTimesToEachTimetable(dbTimetables)
	models.SetClassesToEachTimetable(dbTimetables)
	graphTimetables := convert.ToGraphQLTimetables(dbTimetables)
	return graphTimetables, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	auth := auth.GetAuthInfoFromCtx(ctx)
	if auth == nil {
		err := &myerrors.UnauthenticatedUserAccessError{}
		log.Printf("action=get login user todo data, status=failed, err=%s", err.Error())
		return nil, err
	}
	// todo
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
