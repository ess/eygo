package http

import (
	"net/url"
	"time"

	"github.com/ess/eygo"
)

const (
	// NotFound indicates that the Resource upon which the client is operating
	// was not found.
	NotFound = "404"

	// Forbidden indicates that the authenticated user is not allowed to access
	// a given Resource or API endpoint.
	Forbidden = "403"

	// IllegalOperation is used to express that an attempted operation on the
	// upstream API is not allowed.
	IllegalOperation = "This operation is not allowed in the client."

	perPage  = "100"
	pollTime = 5 * time.Second
)

func paramsToValues(params eygo.Params) url.Values {
	return url.Values(params)
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
