package eygo

import (
	"encoding/json"
	//"fmt"
	"testing"
)

func TestNewEnvironmentService(t *testing.T) {
	driver := NewMockDriver()
	service := NewEnvironmentService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestEnvironmentService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewEnvironmentService(driver)

	t.Run("when there are matching environments", func(t *testing.T) {
		env1 := &Environment{ID: 1, Name: "Environment 1"}
		env2 := &Environment{ID: 2, Name: "Environment 2"}
		env3 := &Environment{ID: 3, Name: "Environment 3"}

		stubEnvironments(driver, env1, env2, env3)

		all := service.All(nil)

		t.Run("it contains all matching environments", func(t *testing.T) {
			environments := []*Environment{env1, env2, env3}

			if len(all) != len(environments) {
				t.Errorf("Expected %d environments, got %d", len(environments), len(all))
			}

			for _, environment := range environments {
				found := false

				for _, other := range all {
					if environment.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Environment %d was not present", environment.ID)
				}
			}
		})

	})

	t.Run("when there are no matching environments", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 environments, got")
			}
		})

	})

}

func TestEnvironmentService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewEnvironmentService(driver)

	t.Run("when there are matching environments", func(t *testing.T) {
		env1 := &Environment{ID: 1, Name: "Environment 1"}
		env2 := &Environment{ID: 2, Name: "Environment 2"}
		env3 := &Environment{ID: 3, Name: "Environment 3"}

		stubAccountEnvironments(driver, account, env1, env2, env3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching environments", func(t *testing.T) {
			environments := []*Environment{env1, env2, env3}

			if len(all) != len(environments) {
				t.Errorf("Expected %d environments, got %d", len(environments), len(all))
			}

			for _, environment := range environments {
				found := false

				for _, other := range all {
					if environment.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Environment %d was not present", environment.ID)
				}
			}
		})

	})

	t.Run("when there are no matching environments", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 environments, got")
			}
		})

	})

}

func stubEnvironments(driver *MockDriver, environments ...*Environment) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Environments []*Environment `json:"environments,omitempty"`
	}{Environments: environments}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "environments", Response{Pages: pages})
	}
}

func stubAccountEnvironments(driver *MockDriver, account *Account, environments ...*Environment) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Environments []*Environment `json:"environments,omitempty"`
	}{Environments: environments}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/environments", Response{Pages: pages})
	}
}
