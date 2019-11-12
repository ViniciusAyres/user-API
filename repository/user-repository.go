package repository

import "github.com/ViniciusAyres/user-API/model"

// User Respository
type User interface {
	Save(u model.User) model.User
}
