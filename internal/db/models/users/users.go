package users

import (
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

func (user *User) UpdateUser(input *model.UpdateUser, u User) (string, error) {
	user.Copy(u)

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
		correct := CheckPasswordHash(*input.CurrentPassword, user.Password)
		if !correct {
			log.Printf("action=update user, status=failed, err=currentPassword is wrong")
			return "", fmt.Errorf("failed to update password. currentPassword is wrong")
		}

		hashedPassword, err := HashPassword(*input.Password)
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

func (user *User) DeleteUser(input model.DeleteUser, u User) (bool, error) {
	user.Copy(u)

	correct := CheckPasswordHash(input.Password, user.Password)
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

func (user *User) RefreshToken(token string) (string, error) {
	id, err := jwt.ParseToken(token)
	if err != nil {
		log.Printf("action=refresh token, status=failed, err=failed to parse token")
		return "", fmt.Errorf("failed to parse given token")
	}
	refreshToken, err := jwt.GenerateToken(id)
	if err != nil {
		log.Printf("action=refresh token, status=failed, err=failed to generate token")
		return "", err
	}
	log.Printf("action=refresh token, status=success")
	return refreshToken, nil
}

func (user *User) Copy(u User) {
	user.ID = u.ID
	user.Email = u.Email
	user.Password = u.Password
	user.Name = u.Name
	user.School = u.School
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
