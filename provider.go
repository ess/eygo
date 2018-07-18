package eygo

import (
  "encoding/json"
)

// Provider is a data structure that models an infrastructure provider on the
// Engine Yard API.
type Provider struct {
	ID            int          `json:"id,omitempty"`
	ProvisionedID string       `json:"provisioned_id,omitempty"`
	Type          string       `json:"type,omitempty"`
	Credentials   *Credentials `json:"credentials,omitempty"`
	CancelledAt   string       `json:"cancelled_at,omitempty"`
	CreatedAt     string       `json:"created_at,omitempty"`
	UpdatedAt     string       `json:"updated_at,omitempty"`
}

// Credentials is a data structure that models the credentials used to
// interact with the service modeled by a Provider.
type Credentials struct {
	InstanceAwsSecretID  string `json:"instance_aws_secret_id,omitempty"`
	InstanceAwsSecretKey string `json:"instance_aws_secret_key,omitempty"`
	AwsSecretID          string `json:"aws_secret_id"`
	AwsSecretKey         string `json:"aws_secret_key,omitempty"`
	AwsLogin             string `json:"aws_login,omitempty"`
	AwsPass              string `json:"aws_pass,omitempty"`
	PayerAccountID       string `json:"payer_account_id,omitempty"`
}

// ProviderService is a repository one can use to retrieve Provider records
// from the API.
type ProviderService struct {
	Driver Driver
}

// NewProviderService returns a ProviderService configured with the provided
// Driver.
func NewProviderService(driver Driver) *ProviderService {
	return &ProviderService{Driver: driver}
}

// All returns an array of all Providers that match the provided Params.
func (service *ProviderService) All(params Params) []*Provider {
  return service.collection("providers", params)
}

// ForAccount returns an array of Provider records that are both associated
// with the provided Account and matches for the provided Params.
func (service *ProviderService) ForAccount(account *Account, params Params) []*Provider {
  return service.collection("accounts/" + account.ID + "/providers", params)
}

func (service *ProviderService) collection(path string, params Params) []*Provider {
	providers := make([]*Provider, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Providers []*Provider `json:"providers,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				providers = append(providers, wrapper.Providers...)
			}
		}
	}

	return providers
}
