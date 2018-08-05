package eygo

import (
	"encoding/json"
)

// Network is a data structure that models a network on the Engine Yard API
type Network struct {
	ID            string `json:"id,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	DeletedAt     string `json:"deleted_at,omitempty"`
	CIDR          string `json:"cidr,omitempty"`
	Tenancy       string `json:"tenancy,omitempty"`
	ProvisionedID string `json:"provisioned_id,omitempty"`
	ProviderURL   string `json:"provider,omitempty"`
	Location      string `json:"location,omitempty"`
}

// NetworkService is a repository that one can use to create, retrieve, delete,
// and perform other operations on Network records on the API.
type NetworkService struct {
	Driver Driver
}

// NewNetworkService returns a NetworkService configured with the provided Driver.
func NewNetworkService(driver Driver) *NetworkService {
	return &NetworkService{Driver: driver}
}

// All returns an array of Network records that match the given Params.
func (service *NetworkService) All(params Params) []*Network {
	return service.collection("networks", params)
}

// Find returns the Network record identified by the given network id. If there
// are errors in retrieving this information, an error is returned as well.
func (service *NetworkService) Find(id string) (*Network, error) {
	response := service.Driver.Get("networks/"+id, nil)
	if response.Okay() {
		wrapper := struct {
			Network *Network `json:"network,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Network, nil
	}

	return nil, response.Error
}

func (service *NetworkService) collection(path string, params Params) []*Network {
	networks := make([]*Network, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Networks []*Network `json:"networks,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				networks = append(networks, wrapper.Networks...)
			}
		}
	}

	return networks
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
