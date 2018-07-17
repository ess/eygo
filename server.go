package eygo

import (
  "encoding/json"
)

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

//type ServerService interface {
	//All(url.Values) []*Server
	//ForEnvironment(*Environment, url.Values) []*Server
//}

type ServerService struct {
	Driver Driver
}

func NewServerService(driver Driver) *ServerService {
	return &ServerService{Driver: driver}
}

func (service *ServerService) All(params Params) []*Server {
  servers := make([]*Server, 0)
  response := service.Driver.Get("servers", params)

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

