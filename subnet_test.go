package eygo

import (
	"encoding/json"
	"testing"
)

func TestNewSubnetService(t *testing.T) {
	driver := NewMockDriver()
	service := NewSubnetService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestSubnetService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewSubnetService(driver)

	t.Run("when there are matching subnets", func(t *testing.T) {
		subnet1 := &Subnet{ID: "Subnet 1"}
		subnet2 := &Subnet{ID: "Subnet 2"}
		subnet3 := &Subnet{ID: "Subnet 3"}

		stubSubnets(driver, subnet1, subnet2, subnet3)

		all := service.All(nil)

		t.Run("it contains all matching subnets", func(t *testing.T) {
			subnets := []*Subnet{subnet1, subnet2, subnet3}

			if len(all) != len(subnets) {
				t.Errorf("Expected %d subnets, got %d", len(subnets), len(all))
			}

			for _, subnet := range subnets {
				found := false

				for _, other := range all {
					if subnet.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Subnet %s was not present", subnet.ID)
				}
			}
		})

	})

	t.Run("when there are no matching subnets", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 subnets, got")
			}
		})

	})

}

func TestSubnetService_Find(t *testing.T) {
	driver := NewMockDriver()
	service := NewSubnetService(driver)
	subnet := &Subnet{ID: "1"}
	stubSubnet(driver, subnet)

	t.Run("for a known subnet", func(t *testing.T) {
		result, err := service.Find("1")

		t.Run("it is the subneted subnet", func(t *testing.T) {
			if result.ID != subnet.ID {
				t.Errorf("Expected subnet 1, got subnet %s", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown subnet", func(t *testing.T) {
		result, err := service.Find("2")

		t.Run("it returns no subnet", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no subnet, got subnet %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func TestSubnetService_ForNetwork(t *testing.T) {
	network := &Network{ID: "1"}
	driver := NewMockDriver()
	service := NewSubnetService(driver)

	t.Run("when there are matching subnets", func(t *testing.T) {
		subnet1 := &Subnet{ID: "Subnet 1"}
		subnet2 := &Subnet{ID: "Subnet 2"}
		subnet3 := &Subnet{ID: "Subnet 3"}

		stubNetworkSubnets(driver, network, subnet1, subnet2, subnet3)

		all := service.ForNetwork(network, nil)

		t.Run("it contains all matching subnets", func(t *testing.T) {
			subnets := []*Subnet{subnet1, subnet2, subnet3}

			if len(all) != len(subnets) {
				t.Errorf("Expected %d subnets, got %d", len(subnets), len(all))
			}

			for _, subnet := range subnets {
				found := false

				for _, other := range all {
					if subnet.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Subnet %s was not present", subnet.ID)
				}
			}
		})

	})

	t.Run("when there are no matching subnets", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForNetwork(network, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 subnets, got")
			}
		})

	})

}

func stubSubnets(driver *MockDriver, subnets ...*Subnet) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Subnets []*Subnet `json:"subnets,omitempty"`
	}{Subnets: subnets}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "subnets", Response{Pages: pages})
	}
}

func stubNetworkSubnets(driver *MockDriver, network *Network, subnets ...*Subnet) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Subnets []*Subnet `json:"subnets,omitempty"`
	}{Subnets: subnets}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "networks/"+network.ID+"/subnets", Response{Pages: pages})
	}
}

func stubSubnet(driver *MockDriver, subnet *Subnet) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Subnet *Subnet `json:"subnet,omitempty"`
	}{Subnet: subnet}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "subnets/"+subnet.ID, Response{Pages: pages})
	}
}
