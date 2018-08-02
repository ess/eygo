package eygo

import (
	"encoding/json"
	//"fmt"
	"testing"
)

func TestNewUserService(t *testing.T) {
	driver := NewMockDriver()
	service := NewUserService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestUserService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewUserService(driver)

	t.Run("when there are matching users", func(t *testing.T) {
		user1 := &User{ID: "1", Name: "User 1"}
		user2 := &User{ID: "2", Name: "User 2"}
		user3 := &User{ID: "3", Name: "User 3"}

		stubUsers(driver, user1, user2, user3)

		all := service.All(nil)

		t.Run("it contains all matching users", func(t *testing.T) {
			users := []*User{user1, user2, user3}

			if len(all) != len(users) {
				t.Errorf("Expected %d users, got %d", len(users), len(all))
			}

			for _, user := range users {
				found := false

				for _, other := range all {
					if user.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("User %s was not present", user.ID)
				}
			}
		})

	})

	t.Run("when there are no matching users", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 users, got")
			}
		})

	})

}

func TestUserService_Current(t *testing.T) {
	driver := NewMockDriver()
	service := NewUserService(driver)
	user := &User{ID: "1", Name: "User 1"}
	stubCurrent(driver, user)

	t.Run("for a known user", func(t *testing.T) {
		result, err := service.Current()

		t.Run("it is the current user", func(t *testing.T) {
			if result.ID != user.ID {
				t.Errorf("Expected user 1, got user %s", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown user", func(t *testing.T) {
		driver.Reset()
		result, err := service.Current()

		t.Run("it returns no user", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no user, got user %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func stubUsers(driver *MockDriver, users ...*User) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Users []*User `json:"users,omitempty"`
	}{Users: users}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "users", Response{Pages: pages})
	}
}

func stubCurrent(driver *MockDriver, current *User) {
	pages := make([][]byte, 0)

	wrapper := struct {
		User *User `json:"user,omitempty"`
	}{User: current}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "users/current", Response{Pages: pages})
	}
}
