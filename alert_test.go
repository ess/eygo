package eygo

import (
	"encoding/json"
	//"fmt"
	"strconv"
	"testing"
)

func TestNewAlertService(t *testing.T) {
	driver := NewMockDriver()
	service := NewAlertService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestAlertService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewAlertService(driver)

	t.Run("when there are matching alerts", func(t *testing.T) {
		alert1 := &Alert{ID: "Alert 1"}
		alert2 := &Alert{ID: "Alert 2"}
		alert3 := &Alert{ID: "Alert 3"}

		stubAlerts(driver, alert1, alert2, alert3)

		all := service.All(nil)

		t.Run("it contains all matching alerts", func(t *testing.T) {
			alerts := []*Alert{alert1, alert2, alert3}

			if len(all) != len(alerts) {
				t.Errorf("Expected %d alerts, got %d", len(alerts), len(all))
			}

			for _, alert := range alerts {
				found := false

				for _, other := range all {
					if alert.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Alert %s was not present", alert.ID)
				}
			}
		})

	})

	t.Run("when there are no matching alerts", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 alerts, got")
			}
		})

	})

}

func TestAlertService_Find(t *testing.T) {
	driver := NewMockDriver()
	service := NewAlertService(driver)
	alert := &Alert{ID: "1"}
	stubAlert(driver, alert)

	t.Run("for a known alert", func(t *testing.T) {
		result, err := service.Find("1")

		t.Run("it is the alerted alert", func(t *testing.T) {
			if result.ID != alert.ID {
				t.Errorf("Expected alert 1, got alert %s", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown alert", func(t *testing.T) {
		result, err := service.Find("2")

		t.Run("it returns no alert", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no alert, got alert %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func TestAlertService_ForEnvironment(t *testing.T) {
	environment := &Environment{ID: 1, Name: "Environment 1"}
	driver := NewMockDriver()
	service := NewAlertService(driver)

	t.Run("when there are matching alerts", func(t *testing.T) {
		alert1 := &Alert{ID: "Alert 1"}
		alert2 := &Alert{ID: "Alert 2"}
		alert3 := &Alert{ID: "Alert 3"}

		stubEnvironmentAlerts(driver, environment, alert1, alert2, alert3)

		all := service.ForEnvironment(environment, nil)

		t.Run("it contains all matching alerts", func(t *testing.T) {
			alerts := []*Alert{alert1, alert2, alert3}

			if len(all) != len(alerts) {
				t.Errorf("Expected %d alerts, got %d", len(alerts), len(all))
			}

			for _, alert := range alerts {
				found := false

				for _, other := range all {
					if alert.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Alert %s was not present", alert.ID)
				}
			}
		})

	})

	t.Run("when there are no matching alerts", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForEnvironment(environment, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 alerts, got")
			}
		})

	})

}

func stubAlerts(driver *MockDriver, alerts ...*Alert) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Alerts []*Alert `json:"alerts,omitempty"`
	}{Alerts: alerts}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "alerts", Response{Pages: pages})
	}
}

func stubEnvironmentAlerts(driver *MockDriver, environment *Environment, alerts ...*Alert) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Alerts []*Alert `json:"alerts,omitempty"`
	}{Alerts: alerts}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "environments/"+strconv.Itoa(environment.ID)+"/alerts", Response{Pages: pages})
	}
}

func stubAlert(driver *MockDriver, alert *Alert) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Alert *Alert `json:"alert,omitempty"`
	}{Alert: alert}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "alerts/"+alert.ID, Response{Pages: pages})
	}
}
