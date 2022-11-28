// point of this file is to unmarshal json data
// from the body of the http request payload

package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseRequestBody parses the body of an incoming http post request to convert it from json to golang for db storage
func ParseRequestBody(request *http.Request, marshalled_json interface{}) {
	if request_body, err := io.ReadAll(request.Body); err == nil {
		// marshalled_json is the destination data structure to store the converted json
		if err := json.Unmarshal([]byte(request_body), marshalled_json); err != nil {
			return
		}
	}
}
