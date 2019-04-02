package eygo

import (
	"encoding/json"
	"testing"
)

func TestNewNetworkService(t *testing.T) {
	driver := NewMockDriver()
	service := NewNetworkService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestNetworkService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewNetworkService(driver)

	t.Run("when there are matching networks", func(t *testing.T) {
		network1 := &Network{ID: "Network 1"}
		network2 := &Network{ID: "Network 2"}
		network3 := &Network{ID: "Network 3"}

		stubNetworks(driver, network1, network2, network3)

		all := service.All(nil)

		t.Run("it contains all matching networks", func(t *testing.T) {
			networks := []*Network{network1, network2, network3}

			if len(all) != len(networks) {
				t.Errorf("Expected %d networks, got %d", len(networks), len(all))
			}

			for _, network := range networks {
				found := false

				for _, other := range all {
					if network.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Network %s was not present", network.ID)
				}
			}
		})

	})

	t.Run("when there are no matching networks", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 networks, got")
			}
		})

	})

}

func TestNetworkService_Find(t *testing.T) {
	driver := NewMockDriver()
	service := NewNetworkService(driver)
	network := &Network{ID: "1"}
	stubNetwork(driver, network)

	t.Run("for a known network", func(t *testing.T) {
		result, err := service.Find("1")

		t.Run("it is the networked network", func(t *testing.T) {
			if result.ID != network.ID {
				t.Errorf("Expected network 1, got network %s", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown network", func(t *testing.T) {
		result, err := service.Find("2")

		t.Run("it returns no network", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no network, got network %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func stubNetworks(driver *MockDriver, networks ...*Network) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Networks []*Network `json:"networks,omitempty"`
	}{Networks: networks}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "networks", Response{Pages: pages})
	}
}

func stubNetwork(driver *MockDriver, network *Network) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Network *Network `json:"network,omitempty"`
	}{Network: network}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "networks/"+network.ID, Response{Pages: pages})
	}
}
