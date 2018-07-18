package eygo

import (
	"encoding/json"
	"fmt"
)

// Environment is a data structure that models an Engine Yard environment.
type Environment struct {
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	DatabaseStack      string `json:"database_stack,omitempty"`
	DeployMethod       string `json:"deploy_method,omitempty"`
	FrameworkEnv       string `json:"framework_env,omitempty"`
	InternalPrivateKey string `json:"internal_private_key,omitempty"`
	InternalPublicKey  string `json:"internal_public_key,omitempty"`
	Language           string `json:"language,omitempty"`
	Region             string `json:"region,omitempty"`
	ReleaseLabel       string `json:"release_label,omitempty"`
	ServiceLevel       string `json:"service_level,omitempty"`
	ServicePlan        string `json:"service_plan,omitempty"`
	StackName          string `json:"stack_name,omitempty"`
	UserName           string `json:"username,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	DeletedAt          string `json:"deleted_at,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
}

//type EnvironmentService interface {
//All(url.Values) []*Environment
//ForAccount(*Account, url.Values) []*Environment
//Create(*Account, *Environment) (*Environment, error)
//Destroy(*Environment) (*Request, error)
//}

// EnvironmentService is a repository one can use to create, retrieve, update,
// delete, and otherwise operate on Environment records on the API.
type EnvironmentService struct {
	Driver Driver
}

// NewEnvironmentService returns an EnvironmentService configured to use the
// provided Driver.
func NewEnvironmentService(driver Driver) *EnvironmentService {
	return &EnvironmentService{Driver: driver}
}

// All returns an array of all Environment records that match the given Params.
func (service *EnvironmentService) All(params Params) []*Environment {
	return service.collection("environments", params)
}

// ForAccount returns an array of Environments that are both associated with the
// given Account and that match the given Params.
func (service *EnvironmentService) ForAccount(account *Account, params Params) []*Environment {
	return service.collection(
		fmt.Sprintf("accounts/%s/environments", account.ID),
		params,
	)
}

func (service *EnvironmentService) collection(path string, params Params) []*Environment {
	environments := make([]*Environment, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Environments []*Environment `json:"environments,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				environments = append(environments, wrapper.Environments...)
			}
		}
	}

	return environments

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
