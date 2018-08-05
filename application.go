package eygo

import (
	"encoding/json"
	"fmt"
)

// Application is a data structure that models an application on the Engine Yard
// API.
type Application struct {
	ID         int    `json:"id,omitempty"`
	Language   string `json:"language,omitempty"`
	Name       string `json:"name,omitempty"`
	Repository string `json:"repository,omitempty"`
	Type       string `json:"type,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	AccountURL string `json:"account,omitempty"`
}

// ApplicationService is a repository one can use to create, retrieve, delete,
// and otherwise operate on Application records on the API.
type ApplicationService struct {
	Driver Driver
}

// NewApplicationService returns an AddressService configured to use the
// provided Driver.
func NewApplicationService(driver Driver) *ApplicationService {
	return &ApplicationService{Driver: driver}
}

// All returns an array of all Application records that match the given Params.
func (service *ApplicationService) All(params Params) []*Application {
	return service.collection("applications", params)
}

// ForAccount returns an array of Applications that are both associated with the
// given Account and that match the given Params.
func (service *ApplicationService) ForAccount(account *Account, params Params) []*Application {
	return service.collection("accounts/"+account.ID+"/applications", params)
}

// ForEnvironment returns an array of Applications that are both associated
// with the given Account and that match the given Params.
func (service *ApplicationService) ForEnvironment(environment *Environment, params Params) []*Application {
	return service.collection(
		fmt.Sprintf("environments/%d/applications", environment.ID),
		params,
	)
}

func (service *ApplicationService) collection(path string, params Params) []*Application {
	applications := make([]*Application, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Applications []*Application `json:"applications,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				applications = append(applications, wrapper.Applications...)
			}
		}
	}

	return applications
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
