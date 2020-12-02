package users

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	School   string `json:"school"`
	Name     string `json:"name"`
}

