package eygo

import (
	"encoding/json"
	//"fmt"
	"strconv"
	"testing"
)

func TestNewRequestService(t *testing.T) {
	driver := NewMockDriver()
	service := NewRequestService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestRequestService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewRequestService(driver)

	t.Run("when there are matching requests", func(t *testing.T) {
		request1 := &Request{ID: "Request 1"}
		request2 := &Request{ID: "Request 2"}
		request3 := &Request{ID: "Request 3"}

		stubRequests(driver, request1, request2, request3)

		all := service.All(nil)

		t.Run("it contains all matching requests", func(t *testing.T) {
			requests := []*Request{request1, request2, request3}

			if len(all) != len(requests) {
				t.Errorf("Expected %d requests, got %d", len(requests), len(all))
			}

			for _, request := range requests {
				found := false

				for _, other := range all {
					if request.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Request %s was not present", request.ID)
				}
			}
		})

	})

	t.Run("when there are no matching requests", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 requests, got")
			}
		})

	})

}

func TestRequestService_Find(t *testing.T) {
	driver := NewMockDriver()
	service := NewRequestService(driver)
	request := &Request{ID: "1"}
	stubRequest(driver, request)

	t.Run("for a known request", func(t *testing.T) {
		result, err := service.Find("1")

		t.Run("it is the requested request", func(t *testing.T) {
			if result.ID != request.ID {
				t.Errorf("Expected request 1, got request %s", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown request", func(t *testing.T) {
		result, err := service.Find("2")

		t.Run("it returns no request", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no request, got request %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func TestRequestService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewRequestService(driver)

	t.Run("when there are matching requests", func(t *testing.T) {
		request1 := &Request{ID: "Request 1"}
		request2 := &Request{ID: "Request 2"}
		request3 := &Request{ID: "Request 3"}

		stubAccountRequests(driver, account, request1, request2, request3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching requests", func(t *testing.T) {
			requests := []*Request{request1, request2, request3}

			if len(all) != len(requests) {
				t.Errorf("Expected %d requests, got %d", len(requests), len(all))
			}

			for _, request := range requests {
				found := false

				for _, other := range all {
					if request.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Request %s was not present", request.ID)
				}
			}
		})

	})

	t.Run("when there are no matching requests", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 requests, got")
			}
		})

	})

}

func TestRequestService_ForEnvironment(t *testing.T) {
	environment := &Environment{ID: 1, Name: "Environment 1"}
	driver := NewMockDriver()
	service := NewRequestService(driver)

	t.Run("when there are matching requests", func(t *testing.T) {
		request1 := &Request{ID: "Request 1"}
		request2 := &Request{ID: "Request 2"}
		request3 := &Request{ID: "Request 3"}

		stubEnvironmentRequests(driver, environment, request1, request2, request3)

		all := service.ForEnvironment(environment, nil)

		t.Run("it contains all matching requests", func(t *testing.T) {
			requests := []*Request{request1, request2, request3}

			if len(all) != len(requests) {
				t.Errorf("Expected %d requests, got %d", len(requests), len(all))
			}

			for _, request := range requests {
				found := false

				for _, other := range all {
					if request.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Request %s was not present", request.ID)
				}
			}
		})

	})

	t.Run("when there are no matching requests", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForEnvironment(environment, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 requests, got")
			}
		})

	})

}

func TestRequestService_ForServer(t *testing.T) {
	server := &Server{ID: 1}
	driver := NewMockDriver()
	service := NewRequestService(driver)

	t.Run("when there are matching requests", func(t *testing.T) {
		request1 := &Request{ID: "Request 1"}
		request2 := &Request{ID: "Request 2"}
		request3 := &Request{ID: "Request 3"}

		stubServerRequests(driver, server, request1, request2, request3)

		all := service.ForServer(server, nil)

		t.Run("it contains all matching requests", func(t *testing.T) {
			requests := []*Request{request1, request2, request3}

			if len(all) != len(requests) {
				t.Errorf("Expected %d requests, got %d", len(requests), len(all))
			}

			for _, request := range requests {
				found := false

				for _, other := range all {
					if request.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Request %s was not present", request.ID)
				}
			}
		})

	})

	t.Run("when there are no matching requests", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForServer(server, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 requests, got")
			}
		})

	})

}

func stubRequests(driver *MockDriver, requests ...*Request) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Requests []*Request `json:"requests,omitempty"`
	}{Requests: requests}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "requests", Response{Pages: pages})
	}
}

func stubAccountRequests(driver *MockDriver, account *Account, requests ...*Request) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Requests []*Request `json:"requests,omitempty"`
	}{Requests: requests}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/requests", Response{Pages: pages})
	}
}

func stubEnvironmentRequests(driver *MockDriver, environment *Environment, requests ...*Request) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Requests []*Request `json:"requests,omitempty"`
	}{Requests: requests}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "environments/"+strconv.Itoa(environment.ID)+"/requests", Response{Pages: pages})
	}
}

func stubServerRequests(driver *MockDriver, server *Server, requests ...*Request) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Requests []*Request `json:"requests,omitempty"`
	}{Requests: requests}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "servers/"+strconv.Itoa(server.ID)+"/requests", Response{Pages: pages})
	}
}

func stubRequest(driver *MockDriver, request *Request) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Request *Request `json:"request,omitempty"`
	}{Request: request}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "requests/"+request.ID, Response{Pages: pages})
	}
}
