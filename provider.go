package eygo

//import (
  //"encoding/json"
//)

type Provider struct {
	ID            int          `json:"id,omitempty"`
	ProvisionedID string       `json:"provisioned_id,omitempty"`
	Type          string       `json:"type,omitempty"`
	Credentials   *Credentials `json:"credentials,omitempty"`
	CancelledAt   string       `json:"cancelled_at,omitempty"`
	CreatedAt     string       `json:"created_at,omitempty"`
	UpdatedAt     string       `json:"updated_at,omitempty"`
}

type Credentials struct {
	InstanceAwsSecretID  string `json:"instance_aws_secret_id,omitempty"`
	InstanceAwsSecretKey string `json:"instance_aws_secret_key,omitempty"`
	AwsSecretID          string `json:"aws_secret_id"`
	AwsSecretKey         string `json:"aws_secret_key,omitempty"`
	AwsLogin             string `json:"aws_login,omitempty"`
	AwsPass              string `json:"aws_pass,omitempty"`
	PayerAccountID       string `json:"payer_account_id,omitempty"`
}

//type ProviderService interface {
	//ForAccount(*Account, url.Values) []*Provider
//}

type ProviderService struct {
	Driver Driver
}

func NewProviderService(driver Driver) *ProviderService {
	return &ProviderService{Driver: driver}
}
