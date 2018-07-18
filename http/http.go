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
