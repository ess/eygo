package eygo

import (
	"encoding/json"
	"strconv"
)

// KeyPair is a data structure that models a keyPair on the Engine Yard API
type KeyPair struct {
	ID             int    `json:"id,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"deleted_at,omitempty"`
	Fingerprint    string `json:"fingerprint,omitempty"`
	Name           string `json:"name,omitempty"`
	PublicKey      string `json:"public_key,omitempty"`
	UserURL        string `json:"user,omitempty"`
	ApplicationURL string `json:"application,omitempty"`
}

// KeyPairService is a repository that one can use to create, retrieve, delete,
// and perform other operations on KeyPair records on the API.
type KeyPairService struct {
	Driver Driver
}

// NewKeyPairService returns a KeyPairService configured with the provided Driver.
func NewKeyPairService(driver Driver) *KeyPairService {
	return &KeyPairService{Driver: driver}
}

// All returns an array of KeyPair records that match the given Params.
func (service *KeyPairService) All(params Params) []*KeyPair {
	return service.collection("keypairs", params)
}

// ForUser returns an array of KeyPair records that are both associated
// with the given User and matching the given Params.
func (service *KeyPairService) ForUser(user *User, params Params) []*KeyPair {
	return service.collection("users/"+user.ID+"/keypairs", params)
}

// ForEnvironment returns an array of KeyPair records that are both associated
// with the given Environment and matching the given Params.
func (service *KeyPairService) ForEnvironment(environment *Environment, params Params) []*KeyPair {
	return service.collection("environments/"+strconv.Itoa(environment.ID)+"/keypairs", params)
}

// ForApplication returns an array of KeyPair records that are both associated
// with the given Application and matching the given Params.
func (service *KeyPairService) ForApplication(application *Application, params Params) []*KeyPair {
	return service.collection("applications/"+strconv.Itoa(application.ID)+"/keypairs", params)
}

func (service *KeyPairService) collection(path string, params Params) []*KeyPair {
	keyPairs := make([]*KeyPair, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				KeyPairs []*KeyPair `json:"keyPairs,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				keyPairs = append(keyPairs, wrapper.KeyPairs...)
			}
		}
	}

	return keyPairs
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
