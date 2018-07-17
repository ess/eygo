package eygo

import (
  "encoding/json"
)

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

type FeatureService struct {
	Driver Driver
}

func NewFeatureService(driver Driver) *FeatureService {
	return &FeatureService{Driver: driver}
}

func (service *FeatureService) All(params Params) []*Feature {
  features := make([]*Feature, 0)
  response := service.Driver.Get("features", params)

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

