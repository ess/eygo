package eygo

import (
	"encoding/json"
	//"fmt"
	"testing"
)

func TestNewFlavorService(t *testing.T) {
	driver := NewMockDriver()
	service := NewFlavorService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestFlavorService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewFlavorService(driver)

	t.Run("when there are matching flavors", func(t *testing.T) {
		flavor1 := &Flavor{ID: "Flavor 1"}
		flavor2 := &Flavor{ID: "Flavor 2"}
		flavor3 := &Flavor{ID: "Flavor 3"}

		stubAccountFlavors(driver, account, flavor1, flavor2, flavor3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching flavors", func(t *testing.T) {
			flavors := []*Flavor{flavor1, flavor2, flavor3}

			if len(all) != len(flavors) {
				t.Errorf("Expected %d flavors, got %d", len(flavors), len(all))
			}

			for _, flavor := range flavors {
				found := false

				for _, other := range all {
					if flavor.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Flavor %s was not present", flavor.ID)
				}
			}
		})

	})

	t.Run("when there are no matching flavors", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 flavors, got")
			}
		})

	})

}

func stubFlavors(driver *MockDriver, flavors ...*Flavor) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Flavors []*Flavor `json:"flavors,omitempty"`
	}{Flavors: flavors}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "flavors", Response{Pages: pages})
	}
}

func stubAccountFlavors(driver *MockDriver, account *Account, flavors ...*Flavor) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Flavors []*Flavor `json:"flavors,omitempty"`
	}{Flavors: flavors}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/flavors", Response{Pages: pages})
	}
}
