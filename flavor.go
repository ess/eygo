package eygo

import (
	"encoding/json"
)

// Flavor is a data strcture that models an instance flavor on the Engine Yard
// API.
type Flavor struct {
	ID              string `json:"id,omitempty"`
	APIName         string `json:"api_name,omitempty"`
	Description     string `json:"description,omitempty"`
	Dedicated       bool   `json:"dedicated,omitempty"`
	VolumeOptimized bool   `json:"volume_optimized,omitempty"`
	Architecture    int    `json:"architecture,omitempty"`
	Name            string `json:"name,omitempty"`
}

// FlavorService is a repository one can use to retrieve Flavor records from
// the API.
type FlavorService struct {
	Driver Driver
}

// NewFlavorService returns a FlavorService configured with the provided
// Driver.
func NewFlavorService(driver Driver) *FlavorService {
	return &FlavorService{Driver: driver}
}

// ForAccount returns an array of Flavor records that are both associated with
// the provided Account and matches for the provided Params.
func (service *FlavorService) ForAccount(account *Account, params Params) []*Flavor {
	return service.collection("accounts/"+account.ID+"/flavors", params)
}

func (service *FlavorService) collection(path string, params Params) []*Flavor {
	flavors := make([]*Flavor, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Flavors []*Flavor `json:"flavors,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				flavors = append(flavors, wrapper.Flavors...)
			}
		}
	}

	return flavors
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
