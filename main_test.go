package service

import "testing"

type repository struct {
}

func (r repository) Save(u User) User {
	u.ID = "1"
	return u
}

func TestCreateUser(t *testing.T) {
	t.Run("When creating user", func(t *testing.T) {
		fixtureUser := User{
			Name:      "João Roberto",
			Username:  "joao.roberto",
			Password:  "123456",
			Birthdate: "2019-11-12T00:37:44.359Z",
		}

		want := User{
			ID:        "1",
			Name:      "João Roberto",
			Username:  "joao.roberto",
			Password:  "123456",
			Birthdate: "2019-11-12T00:37:44.359Z",
			CreatedAt: "2019-11-12T00:37:44.359Z",
		}

		r := repository{}

		createdUser := CreateUser(fixtureUser, r) // mocka o banco ai otario

		if createdUser.ID != want.ID {
			t.Errorf("got %s want %s", createdUser.ID, want.ID)
		}
	})
}

// func TestPOSTUser(t *testing.T) {
// 	t.Run("returns Pepper's score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodPost, "/users", nil)
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		got := response.Header()
// 		want := "20"

// 		if got != want {
// 			t.Errorf("got '%s', want '%s'", got, want)
// 		}
// 	})
// }
