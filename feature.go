package eygo

import (
	"encoding/json"
)

// Feature is a data structure that models a platform feature on the Engine
// Yard API.
type Feature struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// FeatureService is a repository one can use to retrieve, enable, and disable
// Feature records on the API.
type FeatureService struct {
	Driver Driver
}

// NewFeatureService returns a FeatureService configured with the provided
// Driver.
func NewFeatureService(driver Driver) *FeatureService {
	return &FeatureService{Driver: driver}
}

// All returns an array of Features that matches the given Params.
func (service *FeatureService) All(params Params) []*Feature {
	return service.collection("features", params)
}

// ForAccount returns an array of Features that are both associated with the
// given Account and matching the given Params.
func (service *FeatureService) ForAccount(account *Account, params Params) []*Feature {
	return service.collection("accounts/"+account.ID+"/features", params)
}

// Enable turns the given Feature on for the Account in question.
func (service *FeatureService) Enable(account *Account, feature *Feature) error {
	params := Params{}

	response := service.Driver.Post(
		"accounts/"+account.ID+"/features/"+feature.ID,
		params,
		nil,
	)

	if !response.Okay() {
		return response.Error
	}

	return nil
}

// Disable turns the given feature off for the Account in question.
func (service *FeatureService) Disable(account *Account, feature *Feature) error {
	if feature == nil || len(feature.ID) == 0 {
		return fmt.Errorf("No valid feature given")
	}

	if account == nil || len(account.ID) == 0 {
		return fmt.Errorf("No valid account given")
	}

	response := service.Driver.Delete("accounts/"+account.ID+"/features/"+feature.ID, Params{})

	if !response.Okay() {
		return response.Error
	}

	return nil
}

func (service *FeatureService) collection(path string, params Params) []*Feature {
	features := make([]*Feature, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Features []*Feature `json:"features,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				features = append(features, wrapper.Features...)
			}
		}
	}

	return features
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
