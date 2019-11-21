package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ViniciusAyres/user-API/config"
	"github.com/ViniciusAyres/user-API/controller"
	"github.com/ViniciusAyres/user-API/repository"
)

// type userRepositoryFake struct {
// 	user model.User
// }

// func (r userRepositoryFake) Save(u model.User) model.User {
// 	return r.user
// }

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	db, err := config.NewDB(ctx)
	if err != nil {
		log.Fatalf("todo: database configuration failed: %v", err)
	}

	repo := repository.New(db)
	userController := controller.UserController{repo}

	if err := http.ListenAndServe(":5000", userController); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
