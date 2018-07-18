package eygo

import (
	"encoding/json"
)

// Address is a data structure that models an IP address on the Engine Yard
// API.
type Address struct {
	ID            int    `json:"id,omitempty"`
	ProvisionedID string `json:"provisioned_id,omitempty"`
	IPAddress     string `json:"ip_address,omitempty"`
	Server        string `json:"server,omitempty"`
	Location      string `json:"location,omitempty"`
	Provider      string `json:"provider,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

// AddressService is a repository one can use to retrieve Address records from
// the API.
type AddressService struct {
	Driver Driver
}

// NewAddressService returns an AddressService configured to use the provided
// Driver.
func NewAddressService(driver Driver) *AddressService {
	return &AddressService{Driver: driver}
}

// All returns an array of all Address records that match the given Params.
func (service *AddressService) All(params Params) []*Address {
	return service.collection("addresses", params)
}

// ForAccount returns an array of Addresses that are both associated with the
// given Account and that match the given Params.
func (service *AddressService) ForAccount(account *Account, params Params) []*Address {
	return service.collection("accounts/"+account.ID+"/addresses", params)
}

func (service *AddressService) collection(path string, params Params) []*Address {
	addresses := make([]*Address, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Addresses []*Address `json:"addresses,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				addresses = append(addresses, wrapper.Addresses...)
			}
		}
	}

	return addresses
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
