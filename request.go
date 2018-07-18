package eygo

import (
	"encoding/json"
	"fmt"
)

// Request is a data structure that models a long-running request on the
// Engine Yard API.
//
// Notable examples of such requests are starting a server, booting an
// environment, and so on.
type Request struct {
	ID            string `json:"id,omitempty"`
	Type          string `json:"type,omitempty"`
	Successful    bool   `json:"successful,omitempty"`
	Message       string `json:"message,omitempty"`
	RequestStatus string `json:"request_status,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	DeletedAt     string `json:"deleted_at,omitempty"`
	FinishedAt    string `json:"finished_at,omitempty"`
	StartedAt     string `json:"started_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

// RequestService is a repository one can use to retrieve Request records from
// the API.
type RequestService struct {
	Driver Driver
}

// NewRequestService returns a RequestService configured with the given Driver.
func NewRequestService(driver Driver) *RequestService {
	return &RequestService{Driver: driver}
}

// All returns an array of Request records that match the given Params.
func (service *RequestService) All(params Params) []*Request {
	return service.collection("requests", params)
}

// ForAccount returns an array of Request records that are both associated with
// the given Account as well as matching the given Params.
func (service *RequestService) ForAccount(account *Account, params Params) []*Request {
	return service.collection(
		"accounts/"+account.ID+"/requests",
		params,
	)
}

// ForEnvironment returns an array of Request records that are both associated
// with the given Environment as well as matching the given Params.
func (service *RequestService) ForEnvironment(environment *Environment, params Params) []*Request {
	return service.collection(
		fmt.Sprintf("environments/%d/requests", environment.ID),
		params,
	)
}

// ForServer returns an array of Request records that are both associated with
// the given Server as well as matching the given Params.
func (service *RequestService) ForServer(server *Server, params Params) []*Request {
	return service.collection(
		fmt.Sprintf("servers/%d/requests", server.ID),
		params,
	)
}

// Find returns the Request record identified by the given request id. If there
// are errors in retrieving this information, an error is returned as well.
func (service *RequestService) Find(id string) (*Request, error) {
	response := service.Driver.Get("requests/"+id, nil)
	if response.Okay() {
		wrapper := struct {
			Request *Request `json:"request,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Request, nil
	}

	return nil, response.Error
}

func (service *RequestService) collection(path string, params Params) []*Request {
	requests := make([]*Request, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Requests []*Request `json:"requests,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				requests = append(requests, wrapper.Requests...)
			}
		}
	}

	return requests
}
