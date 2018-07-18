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

//type FeatureService interface {
//All(url.Values) []*Feature
//ForAccount(*Account, url.Values) []*Feature
//Enable(*Account, *Feature) error
//Disable(*Account, *Feature) error
//}

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

// func (service *FeatureService) Enable(account *Account, feature *Feature) error

// func (service *FeatureService) Disable(account *Account, feature *Feature) error

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
