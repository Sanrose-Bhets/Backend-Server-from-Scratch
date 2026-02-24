package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKEY: extracts a API ket from
// the headers of the HTTP Request
// Example:
// Autherization: ApiKey {insert api key here}

// SHOULD NOT CAPATILIZE Errors in go, this is called a Linting error(stylistic error)
func GetAPIKey(headers http.Header) (string, error) {

	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found!!!")
	}
	//using the example there will be 2 vals, the string apikey and the apikey itself
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}
	return vals[1], nil
}
