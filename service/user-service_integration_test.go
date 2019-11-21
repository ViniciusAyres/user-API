package service

import (
	"testing"

	"github.com/ViniciusAyres/user-API/mocks"
	"github.com/ViniciusAyres/user-API/model"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	t.Run("When creating user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

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
			ID:        primitive.NewObjectID(),
			CreatedAt: "2019-11-12T00:37:44.359Z",
		}

		m := mocks.NewMockUserRepository(ctrl)
		m.EXPECT().Save(fixtureUser).Return(want).Times(1)

		createdUser := CreateUser(fixtureUser, m)
		if createdUser.ID != want.ID {
			t.Errorf("got %s want %s", createdUser.ID, want.ID)
		}

	})
}
