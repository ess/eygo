package eygo

import (
	"encoding/json"
	//"fmt"
	"testing"
)

func TestNewAddressService(t *testing.T) {
	driver := NewMockDriver()
	service := NewAddressService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestAddressService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewAddressService(driver)

	t.Run("when there are matching addresss", func(t *testing.T) {
		address1 := &Address{ID: 1, IPAddress: "127.0.0.1"}
		address2 := &Address{ID: 2, IPAddress: "127.0.0.2"}
		address3 := &Address{ID: 3, IPAddress: "127.0.0.3"}

		stubAddresss(driver, address1, address2, address3)

		all := service.All(nil)

		t.Run("it contains all matching addresss", func(t *testing.T) {
			addresss := []*Address{address1, address2, address3}

			if len(all) != len(addresss) {
				t.Errorf("Expected %d addresss, got %d", len(addresss), len(all))
			}

			for _, address := range addresss {
				found := false

				for _, other := range all {
					if address.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Address %d was not present", address.ID)
				}
			}
		})

	})

	t.Run("when there are no matching addresss", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 addresss, got")
			}
		})

	})

}

func TestAddressService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewAddressService(driver)

	t.Run("when there are matching addresss", func(t *testing.T) {
		address1 := &Address{ID: 1, IPAddress: "127.0.0.1"}
		address2 := &Address{ID: 2, IPAddress: "127.0.0.2"}
		address3 := &Address{ID: 3, IPAddress: "127.0.0.3"}

		stubAccountAddresss(driver, account, address1, address2, address3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching addresss", func(t *testing.T) {
			addresss := []*Address{address1, address2, address3}

			if len(all) != len(addresss) {
				t.Errorf("Expected %d addresss, got %d", len(addresss), len(all))
			}

			for _, address := range addresss {
				found := false

				for _, other := range all {
					if address.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Address %d was not present", address.ID)
				}
			}
		})

	})

	t.Run("when there are no matching addresss", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 addresss, got")
			}
		})

	})

}

func stubAddresss(driver *MockDriver, addresss ...*Address) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Addresss []*Address `json:"addresses,omitempty"`
	}{Addresss: addresss}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "addresses", Response{Pages: pages})
	}
}

func stubAccountAddresss(driver *MockDriver, account *Account, addresss ...*Address) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Addresss []*Address `json:"addresses,omitempty"`
	}{Addresss: addresss}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/addresses", Response{Pages: pages})
	}
}
