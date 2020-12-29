package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"github.com/mi11km/zikanwarikun-back/pkg/password"
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
	Timetables []*Timetable `gorm:"many2many:user_timetables;"`
}

func (user *User) Create(input model.NewUser) error {
	hashedPassword, err := password.HashPassword(input.Password)
	if err != nil {
		log.Printf("action=create user, status=failed, err=%s", err)
		return err
	}

	user.ID = uuid.New().String()
	user.Email = input.Email
	user.Password = hashedPassword
	user.School = input.School
	user.Name = input.Name

	result := database.Db.Create(user)
	if result.Error != nil {
		log.Printf("action=create user, status=failed, err=%s", result.Error)
		return result.Error
	}
	return nil
}

func (user *User) Update(input *model.UpdateUser) error {
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
		correct := password.CheckPasswordHash(input.Password.Current, user.Password)
		if !correct {
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
		log.Printf("action=update login user data, status=failed, err=update data must be set")
		return fmt.Errorf("update data must be set")
	}

	result := database.Db.Model(user).Updates(updateData)
	if result.Error != nil {
		log.Printf("action=update login user data, status=failed, err=%s", result.Error)
		return result.Error
	}
	return nil
}

func (user *User) Delete(input model.DeleteUser) (bool, error) {
	correct := password.CheckPasswordHash(input.Password, user.Password)
	if !correct {
		log.Printf("action=delete user, status=failed, err=password is wrong")
		return false, fmt.Errorf("failed to delte user. password is wrong")
	}

	result := database.Db.Select(clause.Associations).Delete(user) // todo 関連レコードも一括削除する。できてるか確認できていない
	if result.Error != nil {
		log.Printf("action=delete user, status=failed, err=%s", result.Error)
		return false, result.Error
	}
	log.Printf("action=delete login user, status=success")
	return true, nil
}

func (user *User) Login(input model.Login) error {
	result := database.Db.Where("email = ?", input.Email).First(user)
	if result.Error != nil {
		log.Printf("action=login, status=failed, err=%s", result.Error)
		return result.Error
	}

	correct := password.CheckPasswordHash(input.Password, user.Password)
	if !correct {
		log.Printf("action=login, status=failed, err=email or password is wrong")
		return fmt.Errorf("failed to login. email or password is wrong")
	}
	return nil
}
