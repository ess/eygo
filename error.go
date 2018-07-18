package eygo

// Error is a specific error type for errors returned from the Engine Yard
// API.
type Error struct {
	ErrorString string
}

// Error is the full error string from the API for a given operation.
func (err *Error) Error() string {
	return err.ErrorString
}

// NewError instantiates a Error.
func NewError(errorString string) *Error {
	return &Error{
		ErrorString: errorString,
	}
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
