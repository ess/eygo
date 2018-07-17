package eygo

//import (
  //"encoding/json"
//)

type Flavor struct {
	ID              string `json:"id,omitempty"`
	APIName         string `json:"api_name,omitempty"`
	Description     string `json:"description,omitempty"`
	Dedicated       bool   `json:"dedicated,omitempty"`
	VolumeOptimized bool   `json:"volume_optimized,omitempty"`
	Architecture    int    `json:"architecture,omitempty"`
	Name            string `json:"name,omitempty"`
}

//type FlavorService interface {
	//ForAccount(*Account, url.Values) []*Feature
//}

type FlavorService struct {
	Driver Driver
}

func NewFlavorService(driver Driver) *FlavorService {
	return &FlavorService{Driver: driver}
}

