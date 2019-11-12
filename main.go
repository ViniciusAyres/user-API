package service

// User ...
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Birthdate string `json:"birthdate"`
	CreatedAt string `json:createdAt`
}

// UserRepositorySaver ...
type UserRepositorySaver interface {
	Save(u User) User
}

// CreateUser creates user
func CreateUser(u User, repository UserRepositorySaver) User {
	return repository.Save(u)
}
