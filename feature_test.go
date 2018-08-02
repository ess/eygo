package eygo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewFeatureService(t *testing.T) {
	driver := NewMockDriver()
	service := NewFeatureService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestFeatureService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewFeatureService(driver)

	t.Run("when there are matching features", func(t *testing.T) {
		feature1 := &Feature{ID: "Feature 1"}
		feature2 := &Feature{ID: "Feature 2"}
		feature3 := &Feature{ID: "Feature 3"}

		stubFeatures(driver, feature1, feature2, feature3)

		all := service.All(nil)

		t.Run("it contains all matching features", func(t *testing.T) {
			features := []*Feature{feature1, feature2, feature3}

			if len(all) != len(features) {
				t.Errorf("Expected %d features, got %d", len(features), len(all))
			}

			for _, feature := range features {
				found := false

				for _, other := range all {
					if feature.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Feature %s was not present", feature.ID)
				}
			}
		})

	})

	t.Run("when there are no matching features", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 features, got")
			}
		})

	})

}

func TestFeatureService_ForAccount(t *testing.T) {
	account := &Account{ID: "1", Name: "Account 1"}
	driver := NewMockDriver()
	service := NewFeatureService(driver)

	t.Run("when there are matching features", func(t *testing.T) {
		feature1 := &Feature{ID: "Feature 1"}
		feature2 := &Feature{ID: "Feature 2"}
		feature3 := &Feature{ID: "Feature 3"}

		stubAccountFeatures(driver, account, feature1, feature2, feature3)

		all := service.ForAccount(account, nil)

		t.Run("it contains all matching features", func(t *testing.T) {
			features := []*Feature{feature1, feature2, feature3}

			if len(all) != len(features) {
				t.Errorf("Expected %d features, got %d", len(features), len(all))
			}

			for _, feature := range features {
				found := false

				for _, other := range all {
					if feature.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Feature %s was not present", feature.ID)
				}
			}
		})

	})

	t.Run("when there are no matching features", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForAccount(account, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 features, got")
			}
		})

	})

}

func TestFeatureService_Enable(t *testing.T) {
	driver := NewMockDriver()
	service := NewFeatureService(driver)
	account := &Account{ID: "1", Name: "Account 1"}
	feature := &Feature{ID: "Feature1"}

	t.Run("when successful", func(t *testing.T) {
		driver.Reset()

		driver.AddResponse(
			"post",
			"accounts/"+account.ID+"/features/"+feature.ID,
			Response{
				Pages: [][]byte{
					[]byte(`true`),
				},
			},
		)

		err := service.Enable(account, feature)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})

	})

	t.Run("when unsuccessful", func(t *testing.T) {
		driver.Reset()
		driver.AddResponse(
			"post",
			"accounts/"+account.ID+"/features/"+feature.ID,
			Response{Error: fmt.Errorf("Oh no!")},
		)

		err := service.Enable(account, feature)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})

	})
}

func TestFeatureService_Disable(t *testing.T) {
	driver := NewMockDriver()
	service := NewFeatureService(driver)

	t.Run("with an invalid account", func(t *testing.T) {
		account := &Account{}
		feature := &Feature{ID: "feature1"}

		err := service.Disable(account, feature)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})

	t.Run("with an invalid feature", func(t *testing.T) {
		account := &Account{ID: "1"}
		feature := &Feature{}

		err := service.Disable(account, feature)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})

	t.Run("with a valid feature and account", func(t *testing.T) {
		account := &Account{ID: "1"}
		feature := &Feature{ID: "feature1"}

		t.Run("when successful", func(t *testing.T) {
			driver.Reset()
			driver.AddResponse(
				"delete",
				"accounts/"+account.ID+"/features/"+feature.ID,
				Response{
					Pages: [][]byte{
						[]byte(`true`),
					},
				},
			)

			err := service.Disable(account, feature)

			t.Run("it returns no error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected no error")
				}
			})

		})

		t.Run("when unsuccessful", func(t *testing.T) {
			driver.Reset()
			driver.AddResponse(
				"delete",
				"accounts/"+account.ID+"/features/"+feature.ID,
				Response{Error: fmt.Errorf("Oh no!")},
			)

			err := service.Disable(account, feature)

			t.Run("it returns an error", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error")
				}
			})

		})
	})
}

func stubFeatures(driver *MockDriver, features ...*Feature) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Features []*Feature `json:"features,omitempty"`
	}{Features: features}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "features", Response{Pages: pages})
	}
}

func stubAccountFeatures(driver *MockDriver, account *Account, features ...*Feature) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Features []*Feature `json:"features,omitempty"`
	}{Features: features}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID+"/features", Response{Pages: pages})
	}
}
