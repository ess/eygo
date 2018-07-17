package eygo

//import (
  //"encoding/json"
	//"time"
//)

type Request struct {
	ID            string `json:"id,omitempty"`
	Type          string `json:"type,omitempty"`
	Successful    bool   `json:"successful,omitempty"`
	Message       string `json:"message,omitempty"`
	RequestStatus string `json:"request_status,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	DeletedAt     string `json:"deleted_at,omitempty"`
	FinishedAt    string `json:"finished_at,omitempty"`
	StartedAt     string `json:"started_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

//type RequestService interface {
	//ForEnvironment(*Environment, url.Values) []*Request
	//Refresh(*Request) (*Request, error)
	//Wait(*Request, time.Duration) (*Request, time.Duration, error)
//}

type RequestService struct {
	Driver Driver
}

func NewRequestService(driver Driver) *RequestService {
	return &RequestService{Driver: driver}
}
