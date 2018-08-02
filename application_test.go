package eygo

import (
	"encoding/json"
	//"fmt"
	"strconv"
	"testing"
)

func TestNewApplicationService(t *testing.T) {
	driver := NewMockDriver()
	service := NewApplicationService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestApplicationService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewApplicationService(driver)

	t.Run("when there are matching applications", func(t *testing.T) {
		app1 := &Application{ID: 1, Name: "Application 1"}
		app2 := &Application{ID: 2, Name: "Application 2"}
		app3 := &Application{ID: 3, Name: "Application 3"}

		stubApplications(driver, app1, app2, app3)

		all := service.All(nil)

		t.Run("it contains all matching applications", func(t *testing.T) {
			applications := []*Application{app1, app2, app3}

			if len(all) != len(applications) {
				t.Errorf("Expected %d applications, got %d", len(applications), len(all))
			}

			for _, application := range applications {
				found := false

				for _, other := range all {
					if application.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Application %d was not present", application.ID)
				}
			}
		})

	})

	t.Run("when there are no matching applications", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 applications, got")
			}
		})

	})

}

func TestApplicationService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewApplicationService(driver)

	t.Run("when there are matching applications", func(t *testing.T) {
		app1 := &Application{ID: 1, Name: "Application 1"}
		app2 := &Application{ID: 2, Name: "Application 2"}
		app3 := &Application{ID: 3, Name: "Application 3"}

		stubAccountApplications(driver, account, app1, app2, app3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching applications", func(t *testing.T) {
			applications := []*Application{app1, app2, app3}

			if len(all) != len(applications) {
				t.Errorf("Expected %d applications, got %d", len(applications), len(all))
			}

			for _, application := range applications {
				found := false

				for _, other := range all {
					if application.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Application %d was not present", application.ID)
				}
			}
		})

	})

	t.Run("when there are no matching applications", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 applications, got")
			}
		})

	})

}

func TestApplicationService_ForEnvironment(t *testing.T) {
	environment := &Environment{ID: 1, Name: "Environment 1"}
	driver := NewMockDriver()
	service := NewApplicationService(driver)

	t.Run("when there are matching applications", func(t *testing.T) {
		app1 := &Application{ID: 1, Name: "Application 1"}
		app2 := &Application{ID: 2, Name: "Application 2"}
		app3 := &Application{ID: 3, Name: "Application 3"}

		stubEnvironmentApplications(driver, environment, app1, app2, app3)

		all := service.ForEnvironment(environment, nil)

		t.Run("it contains all matching applications", func(t *testing.T) {
			applications := []*Application{app1, app2, app3}

			if len(all) != len(applications) {
				t.Errorf("Expected %d applications, got %d", len(applications), len(all))
			}

			for _, application := range applications {
				found := false

				for _, other := range all {
					if application.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Application %d was not present", application.ID)
				}
			}
		})

	})

	t.Run("when there are no matching applications", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForEnvironment(environment, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 applications, got")
			}
		})

	})

}

func stubApplications(driver *MockDriver, applications ...*Application) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Applications []*Application `json:"applications,omitempty"`
	}{Applications: applications}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "applications", Response{Pages: pages})
	}
}

func stubAccountApplications(driver *MockDriver, account *Account, applications ...*Application) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Applications []*Application `json:"applications,omitempty"`
	}{Applications: applications}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/applications", Response{Pages: pages})
	}
}

func stubEnvironmentApplications(driver *MockDriver, environment *Environment, applications ...*Application) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Applications []*Application `json:"applications,omitempty"`
	}{Applications: applications}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "environments/"+strconv.Itoa(environment.ID)+"/applications", Response{Pages: pages})
	}
}
