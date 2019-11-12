package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusAyres/user-API/model"
	"github.com/ViniciusAyres/user-API/repository"
	"github.com/ViniciusAyres/user-API/service"
)

// UserController ...
type UserController struct {
	repository repository.User
}

// SaveUser ...
func (u UserController) SaveUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := service.CreateUser(user, u.repository)

	json.NewEncoder(w).Encode(createdUser)
}
