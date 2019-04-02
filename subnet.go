package eygo

import (
	"encoding/json"
)

// Subnet is a data structure that models a subnet on the Engine Yard API
type Subnet struct {
	ID            string `json:"id,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	DeletedAt     string `json:"deleted_at,omitempty"`
	CIDR          string `json:"cidr,omitempty"`
	ProvisionedID string `json:"provisioned_id,omitempty"`
	Location      string `json:"location,omitempty"`
	NetworkURL    string `json:"network,omitempty"`
	Primary       bool   `json:"primary,omitempty"`
}

// SubnetService is a repository that one can use to create, retrieve, delete,
// and perform other operations on Subnet records on the API.
type SubnetService struct {
	Driver Driver
}

// NewSubnetService returns a SubnetService configured with the provided Driver.
func NewSubnetService(driver Driver) *SubnetService {
	return &SubnetService{Driver: driver}
}

// All returns an array of Subnet records that match the given Params.
func (service *SubnetService) All(params Params) []*Subnet {
	return service.collection("subnets", params)
}

// ForNetwork returns an array of Subnet records that are both associated
// with the given Network and matching the given Params.
func (service *SubnetService) ForNetwork(network *Network, params Params) []*Subnet {
	return service.collection("networks/"+network.ID+"/subnets", params)
}

// Find returns the Subnet record identified by the given subnet id. If there
// are errors in retrieving this information, an error is returned as well.
func (service *SubnetService) Find(id string) (*Subnet, error) {
	response := service.Driver.Get("subnets/"+id, nil)
	if response.Okay() {
		wrapper := struct {
			Subnet *Subnet `json:"subnet,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Subnet, nil
	}

	return nil, response.Error
}

func (service *SubnetService) collection(path string, params Params) []*Subnet {
	subnets := make([]*Subnet, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Subnets []*Subnet `json:"subnets,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				subnets = append(subnets, wrapper.Subnets...)
			}
		}
	}

	return subnets
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
