package eygo

import (
	"net/url"
)

type MockDriver struct {
	requests  map[string][]string
	responses *responseCollection
}

func NewMockDriver() *MockDriver {
	return &MockDriver{}
}

func (driver *MockDriver) Get(path string, params Params) Response {
	return driver.handle("get", path+driver.processParams(params))
}

func (driver *MockDriver) Post(path string, params Params, data []byte) Response {
	return driver.handle("post", path+driver.processParams(params))
}

func (driver *MockDriver) Put(path string, params Params, data []byte) Response {
	return driver.handle("put", path+driver.processParams(params))
}

func (driver *MockDriver) Patch(path string, params Params, data []byte) Response {
	return driver.handle("patch", path+driver.processParams(params))
}

func (driver *MockDriver) Delete(path string, params Params) Response {
	return driver.handle("delete", path+driver.processParams(params))
}

func (driver *MockDriver) Reset() {
	driver.requests = nil
	driver.responses = nil
	driver.setup()
}

func (driver *MockDriver) Requests(method string) []string {
	var requests []string

	driver.setup()

	requests = append(requests, driver.requests[method]...)

	return requests
}

func (driver *MockDriver) AddResponse(method string, path string, response Response) {
	driver.setup()

	driver.responses.add(method, path, response)
}

func (driver *MockDriver) handle(method string, path string) Response {
	driver.setup()

	driver.requests[method] = append(driver.requests[method], path)

	return driver.responses.consume(method, path)
}

func (driver *MockDriver) processParams(params Params) string {
	if len(params) > 0 {
		return "?" + url.Values(params).Encode()
	}

	return ""
}

func (driver *MockDriver) setup() {
	if driver.responses == nil {
		driver.responses = &responseCollection{}
	}

	if driver.requests == nil {
		driver.requests = make(map[string][]string)
	}
}
