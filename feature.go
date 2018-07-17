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
