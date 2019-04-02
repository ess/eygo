package eygo

import (
	"encoding/json"
	//"fmt"
	"testing"
)

func TestNewProviderService(t *testing.T) {
	driver := NewMockDriver()
	service := NewProviderService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestProviderService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewProviderService(driver)

	t.Run("when there are matching providers", func(t *testing.T) {
		provider1 := &Provider{ID: 1, ProvisionedID: "Provider 1"}
		provider2 := &Provider{ID: 2, ProvisionedID: "Provider 2"}
		provider3 := &Provider{ID: 3, ProvisionedID: "Provider 3"}

		stubProviders(driver, provider1, provider2, provider3)

		all := service.All(nil)

		t.Run("it contains all matching providers", func(t *testing.T) {
			providers := []*Provider{provider1, provider2, provider3}

			if len(all) != len(providers) {
				t.Errorf("Expected %d providers, got %d", len(providers), len(all))
			}

			for _, provider := range providers {
				found := false

				for _, other := range all {
					if provider.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Provider %d was not present", provider.ID)
				}
			}
		})

	})

	t.Run("when there are no matching providers", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 providers, got")
			}
		})

	})

}

func TestProviderService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewProviderService(driver)

	t.Run("when there are matching providers", func(t *testing.T) {
		provider1 := &Provider{ID: 1, ProvisionedID: "Provider 1"}
		provider2 := &Provider{ID: 2, ProvisionedID: "Provider 2"}
		provider3 := &Provider{ID: 3, ProvisionedID: "Provider 3"}

		stubAccountProviders(driver, account, provider1, provider2, provider3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching providers", func(t *testing.T) {
			providers := []*Provider{provider1, provider2, provider3}

			if len(all) != len(providers) {
				t.Errorf("Expected %d providers, got %d", len(providers), len(all))
			}

			for _, provider := range providers {
				found := false

				for _, other := range all {
					if provider.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Provider %d was not present", provider.ID)
				}
			}
		})

	})

	t.Run("when there are no matching providers", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 providers, got")
			}
		})

	})

}

func stubProviders(driver *MockDriver, providers ...*Provider) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Providers []*Provider `json:"providers,omitempty"`
	}{Providers: providers}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "providers", Response{Pages: pages})
	}
}

func stubAccountProviders(driver *MockDriver, account *Account, providers ...*Provider) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Providers []*Provider `json:"providers,omitempty"`
	}{Providers: providers}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/providers", Response{Pages: pages})
	}
}
