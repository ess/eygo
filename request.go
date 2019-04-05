package eygo

import (
	"encoding/json"
	"fmt"

	"github.com/ess/debuggable"
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
	Stage         string `json:"stage,omitempty"`
	AccountURL    string `json:"account,omitempty"`
	CallbackURL   string `json:"callback_url,omitempty"`
	ResourceURL   string `json:"resource,omitempty"`
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
			} else {
				if debuggable.Enabled() {
					fmt.Println("[DEBUG] Couldn't unmarshal the request:", err)
				}
			}
		}
	}

	return requests
}

/*
Copyright 2018 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
