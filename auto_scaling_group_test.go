package eygo

import (
	"encoding/json"
	"testing"
)

func TestNewAutoScalingGroupService(t *testing.T) {
	driver := NewMockDriver()
	service := NewAutoScalingGroupService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestAutoScalingGroupService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewAutoScalingGroupService(driver)

	t.Run("when there are matching autoscalinggroups", func(t *testing.T) {
		autoscalinggroup1 := &AutoScalingGroup{ID: "1"}
		autoscalinggroup2 := &AutoScalingGroup{ID: "2"}
		autoscalinggroup3 := &AutoScalingGroup{ID: "3"}

		stubAutoScalingGroups(driver, autoscalinggroup1, autoscalinggroup2, autoscalinggroup3)

		all := service.All(nil)

		t.Run("it contains all matching autoscalinggroups", func(t *testing.T) {
			autoscalinggroups := []*AutoScalingGroup{autoscalinggroup1, autoscalinggroup2, autoscalinggroup3}

			if len(all) != len(autoscalinggroups) {
				t.Errorf("Expected %d autoscalinggroups, got %d", len(autoscalinggroups), len(all))
			}

			for _, autoscalinggroup := range autoscalinggroups {
				found := false

				for _, other := range all {
					if autoscalinggroup.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("AutoScalingGroup %s was not present", autoscalinggroup.ID)
				}
			}
		})

	})

	t.Run("when there are no matching autoscalinggroups", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 autoscalinggroups, got")
			}
		})

	})

}

func TestAutoScalingGroupService_Find(t *testing.T) {
	driver := NewMockDriver()
	service := NewAutoScalingGroupService(driver)
	autoscalinggroup := &AutoScalingGroup{ID: "1"}
	stubAutoScalingGroup(driver, autoscalinggroup)

	t.Run("for a known autoscalinggroup", func(t *testing.T) {
		result, err := service.Find("1")

		t.Run("it is the requested autoscalinggroup", func(t *testing.T) {
			if result.ID != autoscalinggroup.ID {
				t.Errorf("Expected autoscalinggroup 1, got autoscalinggroup %s", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown autoscalinggroup", func(t *testing.T) {
		result, err := service.Find("2")

		t.Run("it returns no autoscalinggroup", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no autoscalinggroup, got autoscalinggroup %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func stubAutoScalingGroups(driver *MockDriver, autoscalinggroups ...*AutoScalingGroup) {
	pages := make([][]byte, 0)

	wrapper := struct {
		AutoScalingGroups []*AutoScalingGroup `json:"auto_scaling_groups,omitempty"`
	}{AutoScalingGroups: autoscalinggroups}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "auto_scaling_groups", Response{Pages: pages})
	}
}

func stubAutoScalingGroup(driver *MockDriver, autoscalinggroup *AutoScalingGroup) {
	pages := make([][]byte, 0)

	wrapper := struct {
		AutoScalingGroup *AutoScalingGroup `json:"auto_scaling_group,omitempty"`
	}{AutoScalingGroup: autoscalinggroup}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "auto_scaling_groups/"+autoscalinggroup.ID, Response{Pages: pages})
	}
}
