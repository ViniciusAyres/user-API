package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ViniciusAyres/user-API/model"
)

type stubRepository struct {
	user model.User
}

func (r stubRepository) Save(u model.User) model.User {
	return r.user
}

func TestUserController(t *testing.T) {
	fixtureUser := model.User{
		Name:      "João Roberto",
		Username:  "joao.roberto",
		Password:  "123456",
		Birthdate: "2019-11-12T00:37:44.359Z",
	}

	want := model.User{
		Name:      fixtureUser.Name,
		Username:  fixtureUser.Username,
		Password:  fixtureUser.Password,
		Birthdate: fixtureUser.Birthdate,
		ID:        "1",
		CreatedAt: "2019-11-12T00:37:44.359Z",
	}

	requestByte, _ := json.Marshal(fixtureUser)
	reader := bytes.NewReader(requestByte)

	req, _ := http.NewRequest(http.MethodPost, "/user", reader)
	w := httptest.NewRecorder()

	r := stubRepository{want}
	userController := UserController{r}

	userController.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 201 {
		t.Fatalf("should return 201, got: %d", resp.StatusCode)
	}

	var bodyParsed model.User
	json.NewDecoder(resp.Body).Decode(&bodyParsed)

	if bodyParsed != want {
		t.Fatalf("error! got: [%s], want: [%s]", bodyParsed.ID, want.ID)
	}

}
