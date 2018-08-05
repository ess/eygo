package eygo

import (
	"encoding/json"
	"fmt"
)

// Server is a data structure that models a server on the Engine Yard API.
type Server struct {
	ID               int       `json:"id,omitempty"`
	ProvisionedID    string    `json:"provisioned_id,omitempty"`
	Role             string    `json:"role,omitempty"`
	Dedicated        bool      `json:"dedicated,omitempty"`
	Enabled          bool      `json:"enabled,omitempty"`
	Location         string    `json:"location,omitempty"`
	Name             string    `json:"name,omitempty"`
	PrivateHostname  string    `json:"private_hostname,omitempty"`
	PublicHostname   string    `json:"public_hostname,omitempty"`
	ReleaseLabel     string    `json:"release_label,omitempty"`
	State            string    `json:"state,omitempty"`
	EnvironmentURL   string    `json:"environment,omitempty"`
	AccountURL       string    `json:"account,omitempty"`
	ProviderURL      string    `json:"provider,omitempty"`
	AddressURL       string    `json:"address,omitempty"`
	CreatedAt        string    `json:"created_at,omitempty"`
	DeletedAt        string    `json:"deleted_at,omitempty"`
	DeprovisionedAt  string    `json:"deprovisioned_at,omitempty"`
	ProvisionedAt    string    `json:"provisioned_at,omitempty"`
	DisappearedAt    string    `json:"disappeared_at,omitempty"`
	UpdatedAt        string    `json:"updated_at,omitempty"`
	SSHPort          int       `json:"ssh_port,omitempty"`
	IAMRoleURL       string    `json:"iam_role,omitempty"`
	NetworkURL       string    `json:"network,omitempty"`
	LatestChefLogURL string    `json:"latest_chef_log,omitempty"`
	NoDeploy         bool      `json:"no_deploy,omitempty"`
	Devices          []*Device `json:"devices,omitempty"`
	Flavor           struct {
		ID string `json:"id"`
	} `json:"flavor,omitempty"`
	ChefStatus struct {
		Message   string `json:"message,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
		TimeAgo   string `json:"time_ago,omitempty"`
	} `json:"chef_status,omitempty"`
}

type Device struct {
	Size                int    `json:"size,omitempty"`
	DeleteOnTermination bool   `json:"delete_on_termination,omitempty"`
	Device              string `json:"device,omitempty"`
	VolumeType          string `json:"volume_type,omitempty"`
	Name                string `json:"name,omitempty"`
	NoDevice            bool   `json:"no_device,omitempty"`
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
