package eygo

//import (
  //"encoding/json"
//)

type ProviderLocation struct {
	ID         string  `json:"id,omitempty"`
	LocationID string  `json:"location_id,omitempty"`
	Limits     *Limits `json:"limits,omitempty"`
	CreatedAt  string  `json:"created_at,omitempty"`
	DisabledAt string  `json:"disabled_at,omitempty"`
	UpdatedAt  string  `json:"updated_at,omitempty"`
}

type Limits struct {
	Servers   int `json:"servers,omitempty"`
	Addresses int `json:"addresses,omitempty"`
}

//type ProviderLocationService interface {
	//ForProvider(*Provider, url.Values) []*ProviderLocation
	//Children(*ProviderLocation, url.Values) []*ProviderLocation
//}

type ProviderLocationService struct {
	Driver Driver
}

func NewProviderLocationService(driver Driver) *ProviderLocationService {
	return &ProviderLocationService{Driver: driver}
}
