package eygo

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestNewKeyPairService(t *testing.T) {
	driver := NewMockDriver()
	service := NewKeyPairService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestKeyPairService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewKeyPairService(driver)

	t.Run("when there are matching keyPairs", func(t *testing.T) {
		keyPair1 := &KeyPair{ID: 1}
		keyPair2 := &KeyPair{ID: 2}
		keyPair3 := &KeyPair{ID: 3}

		stubKeyPairs(driver, keyPair1, keyPair2, keyPair3)

		all := service.All(nil)

		t.Run("it contains all matching keyPairs", func(t *testing.T) {
			keyPairs := []*KeyPair{keyPair1, keyPair2, keyPair3}

			if len(all) != len(keyPairs) {
				t.Errorf("Expected %d keyPairs, got %d", len(keyPairs), len(all))
			}

			for _, keyPair := range keyPairs {
				found := false

				for _, other := range all {
					if keyPair.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("KeyPair %d was not present", keyPair.ID)
				}
			}
		})

	})

	t.Run("when there are no matching keyPairs", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 keyPairs, got")
			}
		})

	})

}

func TestKeyPairService_ForUser(t *testing.T) {
	user := &User{ID: "1"}
	driver := NewMockDriver()
	service := NewKeyPairService(driver)

	t.Run("when there are matching keyPairs", func(t *testing.T) {
		keyPair1 := &KeyPair{ID: 1}
		keyPair2 := &KeyPair{ID: 2}
		keyPair3 := &KeyPair{ID: 3}

		stubUserKeyPairs(driver, user, keyPair1, keyPair2, keyPair3)

		all := service.ForUser(user, nil)

		t.Run("it contains all matching keyPairs", func(t *testing.T) {
			keyPairs := []*KeyPair{keyPair1, keyPair2, keyPair3}

			if len(all) != len(keyPairs) {
				t.Errorf("Expected %d keyPairs, got %d", len(keyPairs), len(all))
			}

			for _, keyPair := range keyPairs {
				found := false

				for _, other := range all {
					if keyPair.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("KeyPair %d was not present", keyPair.ID)
				}
			}
		})

	})

	t.Run("when there are no matching keyPairs", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForUser(user, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 keyPairs, got")
			}
		})

	})

}

func TestKeyPairService_ForEnvironment(t *testing.T) {
	environment := &Environment{ID: 1, Name: "Environment 1"}
	driver := NewMockDriver()
	service := NewKeyPairService(driver)

	t.Run("when there are matching keyPairs", func(t *testing.T) {
		keyPair1 := &KeyPair{ID: 1}
		keyPair2 := &KeyPair{ID: 2}
		keyPair3 := &KeyPair{ID: 3}

		stubEnvironmentKeyPairs(driver, environment, keyPair1, keyPair2, keyPair3)

		all := service.ForEnvironment(environment, nil)

		t.Run("it contains all matching keyPairs", func(t *testing.T) {
			keyPairs := []*KeyPair{keyPair1, keyPair2, keyPair3}

			if len(all) != len(keyPairs) {
				t.Errorf("Expected %d keyPairs, got %d", len(keyPairs), len(all))
			}

			for _, keyPair := range keyPairs {
				found := false

				for _, other := range all {
					if keyPair.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("KeyPair %d was not present", keyPair.ID)
				}
			}
		})

	})

	t.Run("when there are no matching keyPairs", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForEnvironment(environment, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 keyPairs, got")
			}
		})

	})

}

func TestKeyPairService_ForApplication(t *testing.T) {
	application := &Application{ID: 1, Name: "Application 1"}
	driver := NewMockDriver()
	service := NewKeyPairService(driver)

	t.Run("when there are matching keyPairs", func(t *testing.T) {
		keyPair1 := &KeyPair{ID: 1}
		keyPair2 := &KeyPair{ID: 2}
		keyPair3 := &KeyPair{ID: 3}

		stubApplicationKeyPairs(driver, application, keyPair1, keyPair2, keyPair3)

		all := service.ForApplication(application, nil)

		t.Run("it contains all matching keyPairs", func(t *testing.T) {
			keyPairs := []*KeyPair{keyPair1, keyPair2, keyPair3}

			if len(all) != len(keyPairs) {
				t.Errorf("Expected %d keyPairs, got %d", len(keyPairs), len(all))
			}

			for _, keyPair := range keyPairs {
				found := false

				for _, other := range all {
					if keyPair.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("KeyPair %d was not present", keyPair.ID)
				}
			}
		})

	})

	t.Run("when there are no matching keyPairs", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForApplication(application, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 keyPairs, got")
			}
		})

	})

}

func stubKeyPairs(driver *MockDriver, keyPairs ...*KeyPair) {
	pages := make([][]byte, 0)

	wrapper := struct {
		KeyPairs []*KeyPair `json:"keyPairs,omitempty"`
	}{KeyPairs: keyPairs}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "keypairs", Response{Pages: pages})
	}
}

func stubUserKeyPairs(driver *MockDriver, user *User, keyPairs ...*KeyPair) {
	pages := make([][]byte, 0)

	wrapper := struct {
		KeyPairs []*KeyPair `json:"keypairs,omitempty"`
	}{KeyPairs: keyPairs}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "users/"+user.ID+"/keypairs", Response{Pages: pages})
	}
}

func stubEnvironmentKeyPairs(driver *MockDriver, environment *Environment, keyPairs ...*KeyPair) {
	pages := make([][]byte, 0)

	wrapper := struct {
		KeyPairs []*KeyPair `json:"keypairs,omitempty"`
	}{KeyPairs: keyPairs}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "environments/"+strconv.Itoa(environment.ID)+"/keypairs", Response{Pages: pages})
	}
}

func stubApplicationKeyPairs(driver *MockDriver, application *Application, keyPairs ...*KeyPair) {
	pages := make([][]byte, 0)

	wrapper := struct {
		KeyPairs []*KeyPair `json:"keypairs,omitempty"`
	}{KeyPairs: keyPairs}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "applications/"+strconv.Itoa(application.ID)+"/keypairs", Response{Pages: pages})
	}
}
