package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ess/debuggable"

	"github.com/ess/eygo"
)

// Driver is an object that knows specifically how to interact with the
// Engine Yard API at the HTTP level
type Driver struct {
	raw     *http.Client
	baseURL url.URL
	token   string
}

// NewDriver takes a base URL for an Engine Yard API and a token, returning a
// Driver that can be used to interact with the API in question.
func NewDriver(baseURL string, token string) (*Driver, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	d := &Driver{
		&http.Client{Timeout: 20 * time.Second},
		*url,
		token,
	}

	return d, nil
}

// Get performs a GET operation for the given path and params against the
// upstream API. it returns a byte array and an error.
func (driver *Driver) Get(path string, params eygo.Params) eygo.Response {
	if params == nil {
		params = make(eygo.Params)
	}

	params.Set("page", "1")
	params.Set("per_page", perPage)
	return driver.makeRequest("GET", path, paramsToValues(params), nil)
}

// Post performs a POST operation for the given path, params, and data against
// the upstream API. it returns a byte array and an error.
func (driver *Driver) Post(path string, params eygo.Params, data []byte) eygo.Response {
	return driver.makeRequest("POST", path, paramsToValues(params), data)
}

// Put performs a PUT operation for the given path, params, and data against
// the upstream API. It returns a byte array and an error.
func (driver *Driver) Put(path string, params eygo.Params, data []byte) eygo.Response {
	return driver.makeRequest("PUT", path, paramsToValues(params), data)
}

// Patch performs a PATCH operation for the given path, params, and data against
// the upstream API. it returns a byte array and an error.
func (driver *Driver) Patch(path string, params eygo.Params, data []byte) eygo.Response {
	return driver.makeRequest("PATCH", path, paramsToValues(params), data)
}

// Delete performs a DELETE operation for the given path and params against the
// upstream API. it returns a byte array and an error.
func (driver *Driver) Delete(path string, params eygo.Params) eygo.Response {
	return driver.makeRequest("DELETE", path, paramsToValues(params), nil)
}

//func (driver *Driver) multiRequest(verb string, path string, params url.Values, data []byte) eygo.Response {}

func (driver *Driver) rawRequest(verb string, path string, params url.Values, data []byte) (*http.Response, []byte, error) {

	request, err := driver.newRequest(verb, path, params, data)
	if err != nil {
		return nil, nil, err
	}

	response, err := driver.raw.Do(request)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()

	if debuggable.Enabled() {
		fmt.Println("[DEBUG] Method:", verb)
		fmt.Println("[DEBUG] Status code:", response.StatusCode)
	}

	if response.StatusCode > 299 {
		return nil, nil,
			fmt.Errorf(
				"The upstream API returned the following status: %d",
				response.StatusCode,
			)
	}

	return response, body, nil
}

func (driver *Driver) makeRequest(verb string, path string, params url.Values, data []byte) eygo.Response {

	pages := make([][]byte, 0)

	response, page, err := driver.rawRequest(verb, path, params, data)
	if err != nil {
		return eygo.Response{nil, err}
	}

	pages = append(pages, page)

	totalPages := driver.pageCount(response.Header.Get("X-Total-Count"))
	currentPage := 1

	for currentPage < totalPages {
		time.Sleep(1 * time.Second)
		currentPage = currentPage + 1
		params.Set("page", strconv.Itoa(currentPage))
		params.Set("per_page", perPage)

		if _, p, e := driver.rawRequest(verb, path, params, data); e == nil {
			pages = append(pages, p)
		}
	}

	if debuggable.Enabled() {
		readable := make([]string, 0)
		for _, page := range pages {
			readable = append(readable, string(page))
		}

		fmt.Println("[DEBUG] Body:", readable)
	}

	return eygo.Response{pages, nil}
}

func (driver *Driver) pageCount(total string) int {
	if len(total) == 0 {
		return 1
	}

	records, err := strconv.Atoi(total)
	if err != nil {
		return 1
	}

	max, _ := strconv.Atoi(perPage)

	pages := records / max
	if records%max > 0 {
		pages = pages + 1
	}

	return pages
}

func (driver *Driver) newRequest(verb string, path string, params url.Values, data []byte) (*http.Request, error) {
	request, err := http.NewRequest(
		verb,
		driver.constructRequestURL(path, params),
		bytes.NewReader(data),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("X-EY-TOKEN", driver.token)
	request.Header.Add("Accept", "application/vnd.engineyard.v3+json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "eygo/0.1.0 (https://github.com/engineyard/eygo)")

	return request, nil
}

func (driver *Driver) constructRequestURL(path string, params url.Values) string {

	pathParts := []string{driver.baseURL.Path, path}

	requestURL := url.URL{
		Scheme:   driver.baseURL.Scheme,
		Host:     driver.baseURL.Host,
		Path:     strings.Join(pathParts, "/"),
		RawQuery: params.Encode(),
	}

	result := requestURL.String()

	if debuggable.Enabled() {
		fmt.Println("[DEBUG] Request URL:", result)
	}

	return result
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
