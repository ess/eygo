package eygo

import (
	"encoding/json"
)

// Addon is a data structure that models a platform addon on the Engine
// Yard API.
type Addon struct {
	ID     int               `json:"id,omitempty"`
	Name   string            `json:"name,omitempty"`
	SSOURL string            `json:"sso_url,omitempty"`
	Vars   map[string]string `json:"vars,omitempty"`
}

// AddonService is a repository one can use to retrieve, enable, and disable
// Addon records on the API.
type AddonService struct {
	Driver Driver
}

// NewAddonService returns a AddonService configured with the provided
// Driver.
func NewAddonService(driver Driver) *AddonService {
	return &AddonService{Driver: driver}
}

// ForAccount returns an array of Addons that are both associated with the
// given Account and matching the given Params.
func (service *AddonService) ForAccount(account *Account, params Params) []*Addon {
	return service.collection("accounts/"+account.ID+"/addons", params)
}

func (service *AddonService) collection(path string, params Params) []*Addon {
	addons := make([]*Addon, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Addons []*Addon `json:"addons,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				addons = append(addons, wrapper.Addons...)
			}
		}
	}

	return addons
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
