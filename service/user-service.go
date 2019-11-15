package service

import (
	"github.com/ViniciusAyres/user-API/model"
	"github.com/ViniciusAyres/user-API/repository"
)

// CreateUser creates user
func CreateUser(u model.User, r repository.UserRepository) model.User {
	return r.Save(u)
}
