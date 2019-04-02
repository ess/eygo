package eygo

import (
	"encoding/json"
	"fmt"
)

// Alert is a data structure that models an alert on the Engine Yard API
type Alert struct {
	ID             string `json:"id,omitempty"`
	Acknowledged   bool   `json:"acknowledged,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	Description    string `json:"description,omitempty"`
	ExternalID     string `json:"external_id,omitempty"`
	StartedAt      string `json:"started_at,omitempty"`
	FinishedAt     string `json:"finished_at,omitempty"`
	Ignored        bool   `json:"ignored,omitempty"`
	Message        string `json:"message,omitempty"`
	Name           string `json:"name,omitempty"`
	ResourceURL    string `json:"resource,omitempty"`
	Severity       string `json:"severity,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	Type           string `json:"type,omitempty"`
	AccountURL     string `json:"account,omitempty"`
	EnvironmentURL string `json:"environment,omitempty"`
}

// AlertService is a repository that one can use to create, retrieve, delete,
// and perform other operations on Alert records on the API.
type AlertService struct {
	Driver Driver
}

// NewAlertService returns a AlertService configured with the provided Driver.
func NewAlertService(driver Driver) *AlertService {
	return &AlertService{Driver: driver}
}

// All returns an array of Alert records that match the given Params.
func (service *AlertService) All(params Params) []*Alert {
	return service.collection("alerts", params)
}

// ForEnvironment returns an array of Alert records that are both associated
// with the given Environment and matching the given Params.
func (service *AlertService) ForEnvironment(environment *Environment, params Params) []*Alert {
	return service.collection(
		fmt.Sprintf("environments/%d/alerts", environment.ID),
		params,
	)
}

// Find returns the Alert record identified by the given alert id. If there
// are errors in retrieving this information, an error is returned as well.
func (service *AlertService) Find(id string) (*Alert, error) {
	response := service.Driver.Get("alerts/"+id, nil)
	if response.Okay() {
		wrapper := struct {
			Alert *Alert `json:"alert,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Alert, nil
	}

	return nil, response.Error
}

func (service *AlertService) collection(path string, params Params) []*Alert {
	alerts := make([]*Alert, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				Alerts []*Alert `json:"alerts,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				alerts = append(alerts, wrapper.Alerts...)
			}
		}
	}

	return alerts
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
