package eygo

import (
	"encoding/json"
	"fmt"
)

// Snapshot is a data structure that models a snapshot on the Engine Yard API.
type Snapshot struct {
	ID             int    `json:"id,omitempty"`
	State          string `json:"state,omitempty"`
	Progress       int    `json:"progress,omitempty"`
	Size           int    `json:"size,omitempty"`
	Snaplocked     bool   `json:"snaplocked,omitempty"`
	Grade          string `json:"grade,omitempty"`
	EnvironmentURL string `json:"environment,omitempty"`
	ServerURL      string `json:"server,omitempty"`
	ProviderURL    string `json:"provider,omitempty"`
	AccountURL     string `json:"account,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	DeletedAt      string `json:"deleted_at,omitempty"`
	Region         struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"region,omitempty"`
}

// SnapshotService is a repository that one can use to create, retrieve, delete,
// and perform other operations on Snapshot records on the API.
type SnapshotService struct {
	Driver Driver
}

// NewSnapshotService returns a SnapshotService configured with the provided Driver.
func NewSnapshotService(driver Driver) *SnapshotService {
	return &SnapshotService{Driver: driver}
}

// ForEnvironment returns an array of Snapshot records that are both associated
// with the given Environment and matching the given Params.
func (service *SnapshotService) ForEnvironment(environment *Environment, params Params) []*Snapshot {
	return service.collection(
		fmt.Sprintf("environments/%d/snapshots", environment.ID),
		params,
	)
}

// ForServer returns an array of Snapshot records that are both associated with
// the given Server as well as matching the given Params.
func (service *SnapshotService) ForServer(server *Server, params Params) []*Snapshot {
	return service.collection(
		fmt.Sprintf("servers/%d/snapshots", server.ID),
		params,
	)
}

// Find returns the Snapshot record identified by the given snapshot id. If there
// are errors in retrieving this information, an error is returned as well.
func (service *SnapshotService) Find(id string) (*Snapshot, error) {
	response := service.Driver.Get("snapshots/"+id, nil)
	if response.Okay() {
		wrapper := struct {
			Snapshot *Snapshot `json:"snapshot,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Snapshot, nil
	}

	return nil, response.Error
}

func (service *SnapshotService) collection(path string, params Params) []*Snapshot {
	snapshots := make([]*Snapshot, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Snapshots []*Snapshot `json:"snapshots,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				snapshots = append(snapshots, wrapper.Snapshots...)
			}
		}
	}

	return snapshots
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
