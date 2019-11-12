package model

// User ...
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Birthdate string `json:"birthdate"`
	CreatedAt string `json:createdAt`
}
