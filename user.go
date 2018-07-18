package eygo

import (
  "encoding/json"
)

// User is a data strcture that models a user on the Engine Yard API.
type User struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	APIToken  string `json:"api_token,omitempty"`
	Verified  bool   `json:"verified,omitempty"`
	Staff     bool   `json:"staff,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// UserService is a repository one can use to retrieve User records from
// the API.
type UserService struct {
	Driver Driver
}

// NewUserService returns a UserService configured with the provided Driver.
func NewUserService(driver Driver) *UserService {
	return &UserService{Driver: driver}
}

// All returns an array of all User records that match the provided Params.
func (service *UserService) All(params Params) []*User {
  users := make([]*User, 0)
  response := service.Driver.Get("users", params)

  if response.Okay() {
    for _, page := range response.Pages {
      wrapper := struct {
        Users []*User `json:"users,omitempty"`
      }{}

      if err := json.Unmarshal(page, &wrapper); err == nil {
        users = append(users, wrapper.Users...)
      }
    }
  }

  return users
}

// Current returns the user that is associated with the current API session.
// If there are issues along the way, an error is returned.
func (service *UserService) Current() (*User, error) {
	response := service.Driver.Get("users/current", nil)
	if !response.Okay() {
		return nil, response.Error
	}

	wrapper := struct {
		User *User `json:"user,omitempty"`
	}{}

  err := json.Unmarshal(response.Pages[0], &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.User, nil
}
