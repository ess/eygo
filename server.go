package eygo

import (
	"encoding/json"
)

// Server is a data structure that models a server on the Engine Yard API.
type Server struct {
	ID              int    `json:"id,omitempty"`
	ProvisionedID   string `json:"provisioned_id,omitempty"`
	Role            string `json:"role,omitempty"`
	Dedicated       bool   `json:"dedicated,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	Location        string `json:"location,omitempty"`
	Name            string `json:"name,omitempty"`
	PrivateHostname string `json:"private_hostname,omitempty"`
	PublicHostname  string `json:"public_hostname,omitempty"`
	ReleaseLabel    string `json:"release_label,omitempty"`
	State           string `json:"state,omitempty"`
	EnvironmentURI  string `json:"environment,omitempty"`
	AccountURI      string `json:"account,omitempty"`
	ProviderURI     string `json:"provider,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	DeletedAt       string `json:"deleted_at,omitempty"`
	DeprovisionedAt string `json:"deprovisioned_at,omitempty"`
	ProvisionedAt   string `json:"provisioned_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	Flavor          struct {
		ID string `json:"id"`
	} `json:"flavor,omitempty"`
}

// ServerService is a repository that one can use to create, retrieve, delete,
// and perform other operations on Server records on the API.
type ServerService struct {
	Driver Driver
}

// NewServerService returns a ServerService configured with the provided Driver.
func NewServerService(driver Driver) *ServerService {
	return &ServerService{Driver: driver}
}

// All returns an array of Server records that match the given Params.
func (service *ServerService) All(params Params) []*Server {
	return service.collection("servers", params)
}

// ForAccount returns an array of Server records that are both associated with
// the given Account and matching the given Params.
func (service *ServerService) ForAccount(account *Account, params Params) []*Server {
	return service.collection("accounts/"+account.ID+"/servers", params)
}

// ForEnvironment returns an array of Server records that are both associated
// with the given Environment and matching the given Params.
func (service *ServerService) ForEnvironment(environment *Environment, params Params) []*Server {
	return service.collection(
		fmt.Sprintf("environments/%d/servers", environment.ID),
		params,
	)
}

func (service *ServerService) collection(path string, params Params) []*Server {
	servers := make([]*Server, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Servers []*Server `json:"servers,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				servers = append(servers, wrapper.Servers...)
			}
		}
	}

	return servers
}
