package eygo

import (
	"encoding/json"
	//"fmt"
	"strconv"
	"testing"
)

func TestNewServerService(t *testing.T) {
	driver := NewMockDriver()
	service := NewServerService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestServerService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewServerService(driver)

	t.Run("when there are matching servers", func(t *testing.T) {
		server1 := &Server{ID: 1, Name: "Server 1"}
		server2 := &Server{ID: 2, Name: "Server 2"}
		server3 := &Server{ID: 3, Name: "Server 3"}

		stubServers(driver, server1, server2, server3)

		all := service.All(nil)

		t.Run("it contains all matching servers", func(t *testing.T) {
			servers := []*Server{server1, server2, server3}

			if len(all) != len(servers) {
				t.Errorf("Expected %d servers, got %d", len(servers), len(all))
			}

			for _, server := range servers {
				found := false

				for _, other := range all {
					if server.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Server %d was not present", server.ID)
				}
			}
		})

	})

	t.Run("when there are no matching servers", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 servers, got")
			}
		})

	})

}

func TestServerService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewServerService(driver)

	t.Run("when there are matching servers", func(t *testing.T) {
		server1 := &Server{ID: 1, Name: "Server 1"}
		server2 := &Server{ID: 2, Name: "Server 2"}
		server3 := &Server{ID: 3, Name: "Server 3"}

		stubAccountServers(driver, account, server1, server2, server3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching servers", func(t *testing.T) {
			servers := []*Server{server1, server2, server3}

			if len(all) != len(servers) {
				t.Errorf("Expected %d servers, got %d", len(servers), len(all))
			}

			for _, server := range servers {
				found := false

				for _, other := range all {
					if server.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Server %d was not present", server.ID)
				}
			}
		})

	})

	t.Run("when there are no matching servers", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 servers, got")
			}
		})

	})

}

func TestServerService_ForEnvironment(t *testing.T) {
	environment := &Environment{ID: 1, Name: "Environment 1"}
	driver := NewMockDriver()
	service := NewServerService(driver)

	t.Run("when there are matching servers", func(t *testing.T) {
		server1 := &Server{ID: 1, Name: "Server 1"}
		server2 := &Server{ID: 2, Name: "Server 2"}
		server3 := &Server{ID: 3, Name: "Server 3"}

		stubEnvironmentServers(driver, environment, server1, server2, server3)

		all := service.ForEnvironment(environment, nil)

		t.Run("it contains all matching servers", func(t *testing.T) {
			servers := []*Server{server1, server2, server3}

			if len(all) != len(servers) {
				t.Errorf("Expected %d servers, got %d", len(servers), len(all))
			}

			for _, server := range servers {
				found := false

				for _, other := range all {
					if server.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Server %d was not present", server.ID)
				}
			}
		})

	})

	t.Run("when there are no matching servers", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForEnvironment(environment, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 servers, got")
			}
		})

	})

}

func stubServers(driver *MockDriver, servers ...*Server) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Servers []*Server `json:"servers,omitempty"`
	}{Servers: servers}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "servers", Response{Pages: pages})
	}
}

func stubAccountServers(driver *MockDriver, account *Account, servers ...*Server) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Servers []*Server `json:"servers,omitempty"`
	}{Servers: servers}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/servers", Response{Pages: pages})
	}
}

func stubEnvironmentServers(driver *MockDriver, environment *Environment, servers ...*Server) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Servers []*Server `json:"servers,omitempty"`
	}{Servers: servers}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "environments/"+strconv.Itoa(environment.ID)+"/servers", Response{Pages: pages})
	}
}
