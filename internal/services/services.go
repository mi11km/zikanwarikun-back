package services

import (
	"github.com/mi11km/zikanwarikun-back/graph/model"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
)

type UserService interface {
	Signup(input model.NewUser) (string, error)
	UpdateLoginUser(input *model.UpdateUser, currentUser models.User) (string, error)
	DeleteLoginUser(input model.DeleteUser, currentUser models.User) (bool, error)
	Login(input model.Login) (string, error)
}

type TimetableService interface {
	CreateTimetable(input model.NewTimetable, user models.User) (*model.Timetable, error)
	UpdateTimetable(input model.UpdateTimetable, user models.User) (*model.Timetable, error)
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
