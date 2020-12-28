package services

import (
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models/users"
)

type UserService interface {
	CreateUser(input model.NewUser) (string, error)
	UpdateUser(input *model.UpdateUser, u users.User) (string, error)
	DeleteUser(input model.DeleteUser, u users.User) (bool, error)
	Login(input model.Login) (string, error)
	RefreshToken(token string) (string, error)
}

type TimetableService interface {
	CreateTimetable(input model.NewTimetable, user users.User) (*model.Timetable, error)
	UpdateTimetable(input model.UpdateTimetable, user users.User) (*model.Timetable, error)
	DeleteTimetable(input string) (bool, error)
}

type ClassService interface {
	CreateClass(input model.NewClass) (*model.Class, error)
	UpdateClass(input model.UpdateClass) (*model.Class, error)
	DeleteClass(input string) (bool, error)
}

type ClassTimeService interface {
	CreateClassTime(input model.NewClassTime) (*model.ClassTime, error)
	UpdateClassTime(input model.UpdateClassTime) (*model.ClassTime, error)
}
