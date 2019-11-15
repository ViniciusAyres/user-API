package repository

import "github.com/ViniciusAyres/user-API/model"

//go:generate mockgen -destination=../mocks/user-repository-mock.go -package=mocks github.com/ViniciusAyres/user-API/repository UserRepository

// UserRepository ...
type UserRepository interface {
	Save(u model.User) model.User
}
