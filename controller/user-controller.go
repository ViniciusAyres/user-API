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
	Repository repository.User
}

//  ServeHttp ...
func (u UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := service.CreateUser(user, u.Repository)

	json.NewEncoder(w).Encode(createdUser)
}
