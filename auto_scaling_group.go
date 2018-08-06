package eygo

import (
	"encoding/json"
)

// AutoScalingGroup is a data structure that models an Engine Yard autoscalinggroup.
type AutoScalingGroup struct {
	ID              string `json:"id,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	DeletedAt       string `json:"deleted_at,omitempty"`
	EnvironmentURL  string `json:"environment,omitempty"`
	MinimumSize     int    `json:"minimum_size,omitempty"`
	MaximumSize     int    `json:"maximum_size,omitempty"`
	DesiredCapacity int    `json:"desired_capacity,omitempty"`
	LocationID      string `json:"location_id,omitempty"`
}

// AutoScalingGroupService is a repository one can use to retrieve and save AutoScalingGroup
// records on the API.
type AutoScalingGroupService struct {
	Driver Driver
}

// NewAutoScalingGroupService returns an AutoScalingGroupService configured to use the provided
// Driver.
func NewAutoScalingGroupService(driver Driver) *AutoScalingGroupService {
	return &AutoScalingGroupService{Driver: driver}
}

// All returns an array of all AutoScalingGroup records that match the given Params.
func (service *AutoScalingGroupService) All(params Params) []*AutoScalingGroup {
	return service.collection("auto_scaling_groups", params)
}

// Find returns the AutoScalingGroup record identified by the given autoscalinggroup id. If there
// are errors in retrieving this information, an error is returned as well.
func (service *AutoScalingGroupService) Find(id string) (*AutoScalingGroup, error) {
	response := service.Driver.Get("auto_scaling_groups/"+id, nil)
	if response.Okay() {
		wrapper := struct {
			AutoScalingGroup *AutoScalingGroup `json:"auto_scaling_group,omitempty"`
		}{}

		err := json.Unmarshal(response.Pages[0], &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.AutoScalingGroup, nil
	}

	return nil, response.Error
}

func (service *AutoScalingGroupService) collection(path string, params Params) []*AutoScalingGroup {
	autoscalinggroups := make([]*AutoScalingGroup, 0)
	response := service.Driver.Get(path, params)

	if response.Okay() {
		for _, page := range response.Pages {
			wrapper := struct {
				AutoScalingGroups []*AutoScalingGroup `json:"auto_scaling_groups,omitempty"`
			}{}

			if err := json.Unmarshal(page, &wrapper); err == nil {
				autoscalinggroups = append(autoscalinggroups, wrapper.AutoScalingGroups...)
			}
		}
	}

	return autoscalinggroups
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
