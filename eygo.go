// Package eygo provides Drivers, Services, and entity definitions for
// interacting with version 3 of the Engine Yard Core API programmatically.
package eygo

// Driver is an interface that defines the minimal API to perform low-level
// operations on the upstream Engine Yard REST API.
type Driver interface {
	Get(string, Params) Response
	Post(string, Params, []byte) Response
	Put(string, Params, []byte) Response
	Patch(string, Params, []byte) Response
	Delete(string, Params) Response
}

// Response is a data structure that describes the payload of every Driver
// method.
type Response struct {
	Pages [][]byte
	Error error
}

// Okay returns false if the response contains an error, and true otherwise.
func (response Response) Okay() bool {
	if response.Error == nil {
		return true
	}
	return false
}

// Params is a type that describes filtering options available in all Driver
// methods.
type Params map[string][]string

// Set takes two strings: a key and a value. The value is recorded within the
// receiver, indexed by the key.
func (params Params) Set(key string, value string) {
	params[key] = []string{value}
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
