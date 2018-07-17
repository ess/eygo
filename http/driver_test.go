package http

import (
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

  "github.com/ess/eygo"
)

func TestDriver_Get(t *testing.T) {
	driver, _ := NewDriver("https://api.engineyard.com", "faketoken")
	data := []byte(`{"sausages" : "gold"}`)

	t.Run(
		"when on the happy path",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"GET",
				"https://api.engineyard.com/sausages",
				httpmock.NewStringResponder(200, string(data)),
			)

			result := driver.Get("sausages", nil)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})

		})

	t.Run(
		"when on the happy path (with params)",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			data = []byte(`{"sausages" : "yup"}`)
			httpmock.RegisterResponder(
				"GET",
				"https://api.engineyard.com/sausages?color=gold&page=1&per_page=100",
				httpmock.NewStringResponder(200, string(data)),
			)

			params := eygo.Params{}
			params.Set("color", "gold")

			result := driver.Get("sausages", params)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})
		})

	t.Run(
		"when a wild API error appears",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"GET",
				"https://api.engineyard.com/ed209",
				httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
			)

			result := driver.Get("ed209", nil)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})

	t.Run(
		"when a problem exists betwen the chair and the keyboard",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"GET",
				"https://api.engineyard.com/404",
				httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
			)

			result := driver.Get("404", nil)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})
}

func TestDriver_Post(t *testing.T) {
	driver, _ := NewDriver("https://api.engineyard.com", "faketoken")

	data := []byte(`{"sausages" : "gold"}`)

	t.Run(
		"when on the happy path",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"POST",
				"https://api.engineyard.com/sausages",
				httpmock.NewStringResponder(200, string(data)),
			)

			result := driver.Post("sausages", nil, data)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})

		})

	t.Run(
		"when on the happy path (with params)",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"POST",
				"https://api.engineyard.com/sausages?color=gold",
				httpmock.NewStringResponder(200, string(data)),
			)

			params := eygo.Params{}
			params.Set("color", "gold")
			result := driver.Post("sausages", params, data)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})
		})

	t.Run(
		"when a wild API error appears",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"POST",
				"https://api.engineyard.com/ed209",
				httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
			)

			result := driver.Post("ed209", nil, data)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})

	t.Run(
		"when a problem exists betwen the chair and the keyboard",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"POST",
				"https://api.engineyard.com/404",
				httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
			)

			result := driver.Post("404", nil, data)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})
}

func TestDriver_Put(t *testing.T) {
	driver, _ := NewDriver("https://api.engineyard.com", "faketoken")

	data := []byte(`{"sausages" : "gold"}`)

	t.Run(
		"when on the happy path",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PUT",
				"https://api.engineyard.com/sausages",
				httpmock.NewStringResponder(200, string(data)),
			)

			result := driver.Put("sausages", nil, data)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})

		})

	t.Run(
		"when on the happy path (with params)",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PUT",
				"https://api.engineyard.com/sausages?color=gold",
				httpmock.NewStringResponder(200, string(data)),
			)

			params := eygo.Params{}
			params.Set("color", "gold")
			result := driver.Put("sausages", params, data)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})
		})

	t.Run(
		"when a wild API error appears",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PUT",
				"https://api.engineyard.com/ed209",
				httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
			)

			result := driver.Put("ed209", nil, data)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})

	t.Run(
		"when a problem exists betwen the chair and the keyboard",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PUT",
				"https://api.engineyard.com/404",
				httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
			)

			result := driver.Put("404", nil, data)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})
}

func TestDriver_Patch(t *testing.T) {
	driver, _ := NewDriver("https://api.engineyard.com", "faketoken")

	data := []byte(`{"sausages" : "gold"}`)

	t.Run(
		"when on the happy path",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PATCH",
				"https://api.engineyard.com/sausages",
				httpmock.NewStringResponder(200, string(data)),
			)

			result := driver.Patch("sausages", nil, data)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})

		})

	t.Run(
		"when on the happy path (with params)",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PATCH",
				"https://api.engineyard.com/sausages?color=gold",
				httpmock.NewStringResponder(200, string(data)),
			)

			params := eygo.Params{}
			params.Set("color", "gold")
			result := driver.Patch("sausages", params, data)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})
		})

	t.Run(
		"when a wild API error appears",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PATCH",
				"https://api.engineyard.com/ed209",
				httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
			)

			result := driver.Patch("ed209", nil, data)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run("it has an error", func(t *testing.T) {
				if result.Okay() {
					t.Errorf("Expected a non-nil error")
				}
			})

		})

	t.Run(
		"when a problem exists betwen the chair and the keyboard",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"PATCH",
				"https://api.engineyard.com/404",
				httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
			)

			result := driver.Patch("404", nil, data)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run("it has an error", func(t *testing.T) {
				if result.Okay() {
					t.Errorf("Expected a non-nil error")
				}
			})

		})
}

func TestDriver_Delete(t *testing.T) {
	driver, _ := NewDriver("https://api.engineyard.com", "faketoken")
	data := []byte(`{"sausages" : "gold"}`)

	t.Run(
		"when on the happy path",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"DELETE",
				"https://api.engineyard.com/sausages",
				httpmock.NewStringResponder(200, string(data)),
			)

			result := driver.Delete("sausages", nil)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})

		})

	t.Run(
		"when on the happy path (with params)",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			data = []byte(`{"sausages" : "yup"}`)

			httpmock.RegisterResponder(
				"DELETE",
				"https://api.engineyard.com/sausages?color=gold",
				httpmock.NewStringResponder(200, string(data)),
			)

			params := eygo.Params{}
			params.Set("color", "gold")
			result := driver.Delete("sausages", params)

			t.Run(
				"it is a success",
				func(t *testing.T) {
					if !result.Okay() {
						t.Errorf("Call was not successful!")
					}
				})

			t.Run(
				"it has the expected value",
				func(t *testing.T) {
					if string(result.Pages[0]) != string(data) {
						t.Errorf("Expected '%s', got '%s'", string(data), string(result.Pages[0]))
					}
				})
		})

	t.Run(
		"when a wild API error appears",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"DELETE",
				"https://api.engineyard.com/ed209",
				httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
			)

			result := driver.Delete("ed209", nil)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})

	t.Run(
		"when a problem exists betwen the chair and the keyboard",
		func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				"DELETE",
				"https://api.engineyard.com/404",
				httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
			)

			result := driver.Delete("404", nil)

			t.Run("it has no result", func(t *testing.T) {
				if result.Pages != nil {
					t.Errorf("Expected a nil result")
				}
			})

			t.Run(
				"it has an error",
				func(t *testing.T) {
					if result.Okay() {
						t.Errorf("Expected a non-nil error")
					}
				})

		})
}
