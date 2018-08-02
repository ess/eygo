package eygo

import (
	"encoding/json"
	"fmt"
)

// Account is a data structure that models an Engine Yard account.
type Account struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Plan             string `json:"plan,omitempty"`
	SupportPlan      string `json:"support_plan,omitempty"`
	Type             string `json:"type,omitempty"`
	EmergencyContact string `json:"emergency_contact,omitempty"`
	CanceledAt       string `json:"canceled_at,omitempty"`
	CancelledAt      string `json:"cancelled_at,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

// AccountService is a repository one can use to retrieve and save Account
// records on the API.
type AccountService struct {
	Driver Driver
}

// NewAccountService returns an AccountService configured to use the provided
// Driver.
func NewAccountService(driver Driver) *AccountService {
	return &AccountService{Driver: driver}
}

// All returns an array of all Account records that match the given Params.
func (service *AccountService) All(params Params) []*Account {
	return service.collection("accounts", params)
}

// Find returns the Account record identified by the given account id. If there
// are errors in retrieving this information, an error is returned as well.
func (service *AccountService) Find(id string) (*Account, error) {
	response := service.Driver.Get("accounts/"+id, nil)
	if response.Okay() {
		wrapper := struct {
			Account *Account `json:"account,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Account, nil
	}

	return nil, response.Error
}

// ForUser returns an array of Accounts that are both associated with the
// given User and that match the given Params.
func (service *AccountService) ForUser(user *User, params Params) []*Account {
	return service.collection(fmt.Sprintf("users/%s/accounts", user.ID), params)
}

type accountName struct {
	Name string `json:"name,omitempty"`
}

// Rename takes an Account and a string, saving the Account on the upstream API
// with the new name. IF there are any issues along the way, an error is
// returned. Otherwise, the updated Account is returned.
func (service *AccountService) Rename(account *Account, name string) (*Account, error) {
	wrapper := struct {
		Account *accountName `json:"account,omitempty"`
	}{Account: &accountName{Name: name}}

	body, err := json.Marshal(&wrapper)
	if err != nil {
		return nil, err
	}

	return service.update(account, body)
}

type accountEmergencyContact struct {
	EmergencyContact string `json:"emergency_contact,omitempty"`
}

// UpdateEmergencyContact takes an Account and a string, saving the Account on
// the upstream API with the new emergency contact. If there are issues along
// the way, an error is returned. Otherwise, the updated Account is returned.
func (service *AccountService) UpdateEmergencyContact(account *Account, contact string) (*Account, error) {
	wrapper := struct {
		Account *accountEmergencyContact `json:"account,omitempty"`
	}{Account: &accountEmergencyContact{EmergencyContact: contact}}

	body, err := json.Marshal(&wrapper)
	if err != nil {
		return nil, err
	}

	return service.update(account, body)
}

type accountSupportPlan struct {
	SupportPlan string `json:"support_plan,omitempty"`
}

// UpdateSupportPlan takes an Account and a string, saving the Account on
// the upstream API with the new support plan. If there are issues along
// the way, an error is returned. Otherwise, the updated Account is returned.
func (service *AccountService) UpdateSupportPlan(account *Account, plan string) (*Account, error) {
	wrapper := struct {
		Account *accountSupportPlan `json:"account, omitempty"`
	}{&accountSupportPlan{SupportPlan: plan}}

	body, err := json.Marshal(&wrapper)
	if err != nil {
		return nil, err
	}

	return service.update(account, body)
}

func (service *AccountService) update(account *Account, data []byte) (*Account, error) {
	if len(account.ID) == 0 {
		return nil, fmt.Errorf("can't update an account without an ID")
	}

	response := service.Driver.Put("accounts/"+account.ID, nil, data)
	if response.Okay() {
		wrapped := struct {
			Account *Account `json:"account,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapped)
		if err != nil {
			return nil, err
		}

		return wrapped.Account, nil
	}

	return nil, response.Error
}

func (service *AccountService) collection(path string, params Params) []*Account {
	accounts := make([]*Account, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Accounts []*Account `json:"accounts,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				accounts = append(accounts, wrapper.Accounts...)
			}
		}
	}

	return accounts
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
