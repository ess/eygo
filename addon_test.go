package eygo

import (
	"encoding/json"
	"testing"
)

func TestNewAddonService(t *testing.T) {
	driver := NewMockDriver()
	service := NewAddonService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestAddonService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewAddonService(driver)

	t.Run("when there are matching addons", func(t *testing.T) {
		addon1 := &Addon{ID: 1}
		addon2 := &Addon{ID: 2}
		addon3 := &Addon{ID: 3}

		stubAccountAddons(driver, account, addon1, addon2, addon3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching addons", func(t *testing.T) {
			addons := []*Addon{addon1, addon2, addon3}

			if len(all) != len(addons) {
				t.Errorf("Expected %d addons, got %d", len(addons), len(all))
			}

			for _, addon := range addons {
				found := false

				for _, other := range all {
					if addon.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Addon %d was not present", addon.ID)
				}
			}
		})

	})

	t.Run("when there are no matching addons", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 addons, got")
			}
		})

	})

}

func stubAddons(driver *MockDriver, addons ...*Addon) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Addons []*Addon `json:"addons,omitempty"`
	}{Addons: addons}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "addons", Response{Pages: pages})
	}
}

func stubAccountAddons(driver *MockDriver, account *Account, addons ...*Addon) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Addons []*Addon `json:"addons,omitempty"`
	}{Addons: addons}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/addons", Response{Pages: pages})
	}
}
