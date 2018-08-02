package eygo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewAccountService(t *testing.T) {
	driver := NewMockDriver()
	service := NewAccountService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestAccountService_All(t *testing.T) {
	driver := NewMockDriver()
	service := NewAccountService(driver)

	t.Run("when there are matching accounts", func(t *testing.T) {
		account1 := &Account{ID: "1", Name: "Account 1"}
		account2 := &Account{ID: "2", Name: "Account 2"}
		account3 := &Account{ID: "3", Name: "Account 3"}

		stubAccounts(driver, account1, account2, account3)

		all := service.All(nil)

		t.Run("it contains all matching accounts", func(t *testing.T) {
			accounts := []*Account{account1, account2, account3}

			if len(all) != len(accounts) {
				t.Errorf("Expected %d accounts, got %d", len(accounts), len(all))
			}

			for _, account := range accounts {
				found := false

				for _, other := range all {
					if account.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Account %s was not present", account.ID)
				}
			}
		})

	})

	t.Run("when there are no matching accounts", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.All(nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 accounts, got")
			}
		})

	})

}

func TestAccountService_Find(t *testing.T) {
	driver := NewMockDriver()
	service := NewAccountService(driver)
	account := &Account{ID: "1", Name: "Account 1"}
	stubAccount(driver, account)

	t.Run("for a known account", func(t *testing.T) {
		result, err := service.Find("1")

		t.Run("it is the requested account", func(t *testing.T) {
			if result.ID != account.ID {
				t.Errorf("Expected account 1, got account %s", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown account", func(t *testing.T) {
		result, err := service.Find("2")

		t.Run("it returns no account", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no account, got account %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func TestAccountService_ForUser(t *testing.T) {
	user := &User{ID: "1", Email: "user@example.com"}
	driver := NewMockDriver()
	service := NewAccountService(driver)

	t.Run("when there are matching accounts", func(t *testing.T) {
		account1 := &Account{ID: "1", Name: "Account 1"}
		account2 := &Account{ID: "2", Name: "Account 2"}
		account3 := &Account{ID: "3", Name: "Account 3"}

		stubUserAccounts(driver, user, account1, account2, account3)

		all := service.ForUser(user, nil)

		t.Run("it contains all matching accounts", func(t *testing.T) {
			accounts := []*Account{account1, account2, account3}

			if len(all) != len(accounts) {
				t.Errorf("Expected %d accounts, got %d", len(accounts), len(all))
			}

			for _, account := range accounts {
				found := false

				for _, other := range all {
					if account.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Account %s was not present", account.ID)
				}
			}
		})

	})

	t.Run("when there are no matching accounts", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForUser(user, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 accounts, got")
			}
		})

	})

}

func TestAccountService_Rename(t *testing.T) {
	driver := NewMockDriver()
	service := NewAccountService(driver)

	original := &Account{ID: "1", Name: "Account 1"}
	newName := "Account 1 New"

	t.Run("when the update is successful", func(t *testing.T) {
		driver.Reset()

		driver.AddResponse(
			"put",
			"accounts/"+original.ID,
			Response{
				Pages: [][]byte{
					[]byte(
						fmt.Sprintf(
							`{"account": {"id": "1", "name": "%s"}}`,
							newName,
						),
					),
				},
			},
		)

		result, err := service.Rename(original, newName)

		t.Run("it returns the renamed account", func(t *testing.T) {
			if result == nil {
				t.Errorf("Expected an account")
			}

			if result.Name != newName {
				t.Errorf("Expected the account to be renamed")
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("when the update fails", func(t *testing.T) {
		driver.Reset()

		driver.AddResponse(
			"put",
			"accounts/"+original.ID,
			Response{Error: fmt.Errorf("Oh no!")},
		)

		result, err := service.Rename(original, newName)

		t.Run("it returns no account", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no account, got %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func TestAccountService_UpdateEmergencyContact(t *testing.T) {
	driver := NewMockDriver()
	service := NewAccountService(driver)

	original := &Account{ID: "1", EmergencyContact: "Larry"}
	newContact := "Joe"

	t.Run("when the update is successful", func(t *testing.T) {
		driver.Reset()

		driver.AddResponse(
			"put",
			"accounts/"+original.ID,
			Response{
				Pages: [][]byte{
					[]byte(
						fmt.Sprintf(
							`{"account": {"id": "1", "emergency_contact": "%s"}}`,
							newContact,
						),
					),
				},
			},
		)

		result, err := service.UpdateEmergencyContact(original, newContact)

		t.Run("it returns the updated account", func(t *testing.T) {
			if result == nil {
				t.Errorf("Expected an account")
			}

			if result.EmergencyContact != newContact {
				t.Errorf("Expected the account to have a new emergency contact")
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("when the update fails", func(t *testing.T) {
		driver.Reset()

		driver.AddResponse(
			"put",
			"accounts/"+original.ID,
			Response{Error: fmt.Errorf("Oh no!")},
		)

		result, err := service.UpdateEmergencyContact(original, newContact)

		t.Run("it returns no account", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no account, got %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func TestAccountService_UpdateSupportPlan(t *testing.T) {
	driver := NewMockDriver()
	service := NewAccountService(driver)

	original := &Account{ID: "1", SupportPlan: "standard"}
	newPlan := "premium"

	t.Run("when the update is successful", func(t *testing.T) {
		driver.Reset()

		driver.AddResponse(
			"put",
			"accounts/"+original.ID,
			Response{
				Pages: [][]byte{
					[]byte(
						fmt.Sprintf(
							`{"account": {"id": "1", "support_plan": "%s"}}`,
							newPlan,
						),
					),
				},
			},
		)

		result, err := service.UpdateSupportPlan(original, newPlan)

		t.Run("it returns the updated account", func(t *testing.T) {
			if result == nil {
				t.Errorf("Expected an account")
			}

			if result.SupportPlan != newPlan {
				t.Errorf("Expected the account to have a new emergency contact")
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("when the update fails", func(t *testing.T) {
		driver.Reset()

		driver.AddResponse(
			"put",
			"accounts/"+original.ID,
			Response{Error: fmt.Errorf("Oh no!")},
		)

		result, err := service.UpdateSupportPlan(original, newPlan)

		t.Run("it returns no account", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no account, got %s", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func stubAccounts(driver *MockDriver, accounts ...*Account) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Accounts []*Account `json:"accounts,omitempty"`
	}{Accounts: accounts}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts", Response{Pages: pages})
	}
}

func stubUserAccounts(driver *MockDriver, user *User, accounts ...*Account) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Accounts []*Account `json:"accounts,omitempty"`
	}{Accounts: accounts}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "users/"+user.ID+"/accounts", Response{Pages: pages})
	}
}

func stubAccount(driver *MockDriver, account *Account) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Account *Account `json:"account,omitempty"`
	}{Account: account}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "accounts/"+account.ID, Response{Pages: pages})
	}
}
