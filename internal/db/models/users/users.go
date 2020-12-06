package users

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	School   string `json:"school"`
	Name     string `json:"name"`
}

// Create create a user in database from given email and password
func (user *User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO Users(id, email, password) VALUES (?,?,?)")
	if err != nil {
		log.Printf("action=prepare create user statement, err=%s", err)
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("action=generate hashpassword, err=%s", err)
	}
	user.ID = uuid.New().String() // todo uuidのDBへの保存方法の最適化(現在は36文字のVARCHAR)
	_, err = statement.Exec(user.ID, user.Email, hashedPassword)
	if err != nil {
		log.Printf("action=create user, err=%s", err)
	}
	log.Println("action=create user, status=success")
}

// Authenticate check if a user exists in database by given email and password
func (user *User) Authenticate() bool {
	statement, err := database.Db.Prepare("select id, password from Users WHERE email = ?")
	if err != nil {
		log.Printf("action=prepare select password from Users statement, err=%s", err)
	}
	row := statement.QueryRow(user.Email)

	var hashedPassword string
	err = row.Scan(&user.ID, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Printf("action=scan password from row, err=%s", err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

//GetUserById get user info(email, school, name) by given id
func GetUserById(id string) (User, error) {
	statement, err := database.Db.Prepare("select email, school, name from Users WHERE id = ?")
	if err != nil {
		log.Printf("action=prepare select user statement by id, err=%s", err)
	}
	row := statement.QueryRow(id)

	var user User
	err = row.Scan(&user.Email, &user.School, &user.Name)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("action=no rows from Users, err=%s", err)
		}
		return User{}, err  // todo エラー時に空の構造体を返すのでいいかわからない
	}
	user.ID = id
	return user, nil
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
