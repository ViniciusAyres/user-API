package service

import (
	"testing"

	"github.com/ViniciusAyres/user-API/model"
)

type stubRepository struct {
	user model.User
}

func (r stubRepository) Save(u model.User) model.User {
	return r.user
}

func TestCreateUser(t *testing.T) {
	t.Run("When creating user", func(t *testing.T) {
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

		r := stubRepository{want}

		createdUser := CreateUser(fixtureUser, r)
		if createdUser.ID != want.ID {
			t.Errorf("got %s want %s", createdUser.ID, want.ID)
		}
	})
}
