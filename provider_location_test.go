package eygo

import (
	"encoding/json"
	//"fmt"
	"strconv"
	"testing"
)

func TestNewProviderLocationService(t *testing.T) {
	driver := NewMockDriver()
	service := NewProviderLocationService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestProviderLocationService_ForProvider(t *testing.T) {
	provider := &Provider{ID: 1}
	driver := NewMockDriver()
	service := NewProviderLocationService(driver)

	t.Run("when there are matching providerLocations", func(t *testing.T) {
		pl1 := &ProviderLocation{ID: "1"}
		pl2 := &ProviderLocation{ID: "2"}
		pl3 := &ProviderLocation{ID: "3"}

		stubProviderProviderLocations(driver, provider, pl1, pl2, pl3)

		all := service.ForProvider(provider, nil)

		t.Run("it contains all matching providerLocations", func(t *testing.T) {
			providerLocations := []*ProviderLocation{pl1, pl2, pl3}

			if len(all) != len(providerLocations) {
				t.Errorf("Expected %d providerLocations, got %d", len(providerLocations), len(all))
			}

			for _, providerLocation := range providerLocations {
				found := false

				for _, other := range all {
					if providerLocation.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("ProviderLocation %s was not present", providerLocation.ID)
				}
			}
		})

	})

	t.Run("when there are no matching providerLocations", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForProvider(provider, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 providerLocations, got")
			}
		})

	})

}

func TestProviderLocationService_Children(t *testing.T) {
	parent := &ProviderLocation{ID: "mommy"}
	driver := NewMockDriver()
	service := NewProviderLocationService(driver)

	t.Run("when there are matching providerLocations", func(t *testing.T) {
		pl1 := &ProviderLocation{ID: "1"}
		pl2 := &ProviderLocation{ID: "2"}
		pl3 := &ProviderLocation{ID: "3"}

		stubChildren(driver, parent, pl1, pl2, pl3)

		all := service.Children(parent, nil)

		t.Run("it contains all matching providerLocations", func(t *testing.T) {
			providerLocations := []*ProviderLocation{pl1, pl2, pl3}

			if len(all) != len(providerLocations) {
				t.Errorf("Expected %d providerLocations, got %d", len(providerLocations), len(all))
			}

			for _, providerLocation := range providerLocations {
				found := false

				for _, other := range all {
					if providerLocation.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("ProviderLocation %s was not present", providerLocation.ID)
				}
			}
		})

	})

	t.Run("when there are no matching providerLocations", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.Children(parent, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 providerLocations, got")
			}
		})

	})

}

func stubProviderProviderLocations(driver *MockDriver, provider *Provider, providerLocations ...*ProviderLocation) {
	pages := make([][]byte, 0)

	wrapper := struct {
		ProviderLocations []*ProviderLocation `json:"provider_locations,omitempty"`
	}{ProviderLocations: providerLocations}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "providers/"+strconv.Itoa(provider.ID)+"/locations", Response{Pages: pages})
	}
}

func stubChildren(driver *MockDriver, parent *ProviderLocation, children ...*ProviderLocation) {
	pages := make([][]byte, 0)

	wrapper := struct {
		ProviderLocations []*ProviderLocation `json:"provider_locations,omitempty"`
	}{ProviderLocations: children}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "provider-locations/"+parent.ID+"/provider-locations", Response{Pages: pages})
	}
}
