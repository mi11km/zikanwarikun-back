package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"github.com/mi11km/zikanwarikun-back/pkg/password"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID         string `gorm:"primaryKey"`
	Email      string
	Password   string
	School     string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Timetables []*Timetable   `gorm:"many2many:user_timetables;"`
}

func (user *User) Create(input model.NewUser) error {
	hashedPassword, err := password.HashPassword(input.Password)
	if err != nil {
		log.Printf("action=create user data, status=failed, err=%s", err)
		return err
	}
	user.ID = uuid.New().String()
	user.Email = input.Email
	user.Password = hashedPassword
	user.School = input.School
	user.Name = input.Name
	if err := database.Db.Create(user).Error; err != nil {
		log.Printf("action=create user data, status=failed, err=%s", err)
		return err
	}
	return nil
}

func (user *User) Update(input *model.UpdateUser) error {
	updateData := make(map[string]interface{})
	if input.Email != nil && *input.Email != user.Email {
		updateData["email"] = *input.Email
	}
	if input.School != nil && *input.School != user.School {
		updateData["school"] = *input.School
	}
	if input.Name != nil && *input.Name != user.Name {
		updateData["name"] = *input.Name
	}
	if input.Password != nil {
		if correct := password.CheckPasswordHash(input.Password.Current, user.Password); !correct {
			log.Printf("action=update login user data, status=failed, err=currentPassword is wrong")
			return fmt.Errorf("failed to update password. currentPassword is wrong")
		}
		hashedPassword, err := password.HashPassword(input.Password.New)
		if err != nil {
			log.Printf("action=update login user data, status=failed, err=%s", err)
			return err
		}
		updateData["password"] = hashedPassword
	}
	if len(updateData) == 0 {
		log.Printf("action=update login user data, status=failed, err=update data must be set or the only same data id set")
		return fmt.Errorf("update data must be set or the only same data id set")
	}

	if err := database.Db.Model(user).Updates(updateData).Error; err != nil {
		log.Printf("action=update login user data, status=failed, err=%s", err)
		return err
	}
	return nil
}

func (user *User) Delete(input model.DeleteUser) (bool, error) {
	if correct := password.CheckPasswordHash(input.Password, user.Password); !correct {
		log.Printf("action=delete login user data, status=failed, err=password is wrong")
		return false, fmt.Errorf("failed to delte user. password is wrong")
	}

	if err := database.Db.Select(clause.Associations).Delete(user).Error; err != nil { // todo? 現在中間テーブルとUserしか削除されない
		log.Printf("action=delete login user data, status=failed, err=%s", err)
		return false, err
	}
	log.Printf("action=delete login user data, status=success")
	return true, nil
}

func (user *User) Login(input model.Login) error {
	if err := database.Db.Where("email = ?", input.Email).First(user).Error; err != nil {
		log.Printf("action=login, status=failed, err=%s", err)
		return err
	}
	if correct := password.CheckPasswordHash(input.Password, user.Password); !correct {
		log.Printf("action=login, status=failed, err=email or password is wrong")
		return fmt.Errorf("failed to login. email or password is wrong")
	}
	return nil
}
