package eygo

import (
	"encoding/json"
	"fmt"
)

// ProviderLocation is a data structure that models a resource location on a
// provider on the Engine Yard API.
//
// At the top level, this is usually an AWS region, and in nested levels it
// is generally an availability zone.
type ProviderLocation struct {
	ID           string  `json:"id,omitempty"`
	LocationID   string  `json:"location_id,omitempty"`
	Limits       *Limits `json:"limits,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	DisabledAt   string  `json:"disabled_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
	LocationName string  `json:"location_name,omitempty"`
	ProviderURL  string  `json:"provider,omitempty"`
	ParentURL    string  `json:"parent,omitempty"`
	Data         struct {
		VPCByDefault bool `json:"vpc_by_default.omitempty"`
	} `json:"data,omitempty"`
}

// Limits is a data structure that models the server and address limits for
// a ProviderLocation.
type Limits struct {
	Servers   int `json:"servers,omitempty"`
	Addresses int `json:"addresses,omitempty"`
}

// ProviderLocationService is a repository one can use to retrieve
// and perform operations on ProviderLocation records from the API.
type ProviderLocationService struct {
	Driver Driver
}

// NewProviderLocationService returns a ProviderLocationService configured with
// the provided Driver.
func NewProviderLocationService(driver Driver) *ProviderLocationService {
	return &ProviderLocationService{Driver: driver}
}

// ForProvider returns an array of ProviderLocation records that are both
// associated with the given Provider as well as matching the given Params.
func (service *ProviderLocationService) ForProvider(provider *Provider, params Params) []*ProviderLocation {
	return service.collection(
		fmt.Sprintf("providers/%d/locations", provider.ID),
		params,
	)
}

// Children returns an array of ProviderLocation records that are both
// associated with the given ProviderLocation as well as matching the given
// Params.
func (service *ProviderLocationService) Children(location *ProviderLocation, params Params) []*ProviderLocation {
	return service.collection(
		"provider-locations/"+location.ID+"/provider-locations",
		params,
	)
}

func (service *ProviderLocationService) collection(path string, params Params) []*ProviderLocation {
	locations := make([]*ProviderLocation, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				ProviderLocations []*ProviderLocation `json:"provider_locations,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				locations = append(locations, wrapper.ProviderLocations...)
			}
		}
	}

	return locations
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
