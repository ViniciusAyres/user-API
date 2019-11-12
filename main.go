package main

import (
	"log"
	"net/http"

	"github.com/ViniciusAyres/user-API/controller"
	"github.com/ViniciusAyres/user-API/model"
)

type userRepositoryFake struct {
	user model.User
}

func (r userRepositoryFake) Save(u model.User) model.User {
	return r.user
}

func main() {
	fakeUser := model.User{
		ID:        "123",
		Name:      "João Roberto",
		Username:  "joao.roberto",
		Password:  "123456",
		Birthdate: "2019-11-12T00:37:44.359Z",
		CreatedAt: "2019-11-12T00:37:44.359Z",
	}

	repo := userRepositoryFake{fakeUser}

	userController := controller.UserController{repo}

	if err := http.ListenAndServe(":5000", userController); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
