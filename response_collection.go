package eygo

import (
	"errors"
)

type responseCollection struct {
	responses map[string][]Response
}

func (rc *responseCollection) add(method string, path string, response Response) {
	identifier := rc.identify(method, path)

	rc.setup(identifier)

	rc.responses[identifier] = append(rc.responses[identifier], response)
}

func (rc *responseCollection) remove(method string, path string) {
	identifier := rc.identify(method, path)

	rc.responses[identifier] = nil

	rc.setup(identifier)
}

func (rc *responseCollection) consume(method string, path string) Response {

	rc.setup("")

	identifier := rc.identify(method, path)

	if len(rc.responses[identifier]) == 0 {
		return Response{Error: errors.New("No response")}
	}

	response := rc.responses[identifier][0]
	rc.trim(identifier)

	return response
}

func (rc *responseCollection) trim(identifier string) {
	if len(rc.responses[identifier]) == 1 {
		rc.responses[identifier] = nil
	} else {
		rc.responses[identifier] = rc.responses[identifier][1:]
	}
}

func (rc *responseCollection) identify(method string, path string) string {
	return method + ":" + path
}

func (rc *responseCollection) setup(scope string) {
	if rc.responses == nil {
		rc.responses = make(map[string][]Response)
	}

	if len(scope) > 0 && rc.responses[scope] == nil {
		rc.responses[scope] = make([]Response, 0)
	}
}
