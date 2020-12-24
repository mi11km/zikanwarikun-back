package users

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"github.com/mi11km/zikanwarikun-back/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	School   string `json:"school"`
	Name     string `json:"name"`
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) CreateUser(input model.NewUser) (string, error) {
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		log.Printf("action=create user, status=failed, err=%s", err)
		return "", err
	}

	user.ID = uuid.New().String()
	user.Email = input.Email
	user.Password = hashedPassword
	user.School = input.School
	user.Name = input.Name

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

func (user *User) UpdateUser(input model.UpdateUser, ctx context.Context) (string,error) {
	return "", nil // todo 認証処理と切り離すべきか
}

func (user *User) DeleteUser(input model.DeleteUser, ctx context.Context) (bool,error) {
	return false, nil

}

func (user *User) Login(input model.Login) (string, error) {
	result := database.Db.Select("id", "password").Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		log.Printf("action=login, status=failed, err=%s", result.Error)
		return "", result.Error
	}

	correct := CheckPasswordHash(input.Password, user.Password)
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

func (user *User) RefreshToken(input model.RefreshTokenInput) (string, error) {
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
