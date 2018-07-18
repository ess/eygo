package eygo

import (
	"encoding/json"
)

// ProviderLocation is a data structure that models a resource location on a
// provider on the Engine Yard API.
//
// At the top level, this is usually an AWS region, and in nested levels it
// is generally an availability zone.
type ProviderLocation struct {
	ID         string  `json:"id,omitempty"`
	LocationID string  `json:"location_id,omitempty"`
	Limits     *Limits `json:"limits,omitempty"`
	CreatedAt  string  `json:"created_at,omitempty"`
	DisabledAt string  `json:"disabled_at,omitempty"`
	UpdatedAt  string  `json:"updated_at,omitempty"`
}

// Limits is a data structure that models the server and address limits for
// a ProviderLocation.
type Limits struct {
	Servers   int `json:"servers,omitempty"`
	Addresses int `json:"addresses,omitempty"`
}

// ProviderLocationService is a repository one can use to retrieve
// and perform operations on ProviderLocation records from the API.
type ProviderLocationService struct {
	Driver Driver
}

// NewProviderLocationService returns a ProviderLocationService configured with
// the provided Driver.
func NewProviderLocationService(driver Driver) *ProviderLocationService {
	return &ProviderLocationService{Driver: driver}
}

// ForProvider returns an array of ProviderLocation records that are both
// associated with the given Provider as well as matching the given Params.
func (service *ProviderLocationService) ForProvider(provider *Provider, params Params) []*ProviderLocation {
	return service.collection(
		fmt.Sprintf("providers/%d/locations", provider.ID),
		params,
	)
}

// Children returns an array of ProviderLocation records that are both
// associated with the given ProviderLocation as well as matching the given
// Params.
func (service *ProviderLocationService) Children(location *ProviderLocation, params Params) []*ProviderLocation {
	return service.collection(
		"provider-locations/"+location.ID+"/provider-locations",
		params,
	)
}

func (service *ProviderLocationService) collection(path string, params Params) []*ProviderLocation {
	locations := make([]*ProviderLocation, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				ProviderLocations []*ProviderLocations `json:"provider_locations,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				locations = append(locations, wrapper.ProviderLocations...)
			}
		}
	}

	return locations
}
