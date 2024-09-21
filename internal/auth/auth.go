package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API key from
// the headers of an HTTP request
// Example of header:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization") // Get Authorization header
	if val == "" {
		return "", errors.New("no authentication info found")
		// By convention we should not start errors messages with a capital letter
		// for example we can use it in fmt.Sprintf("Error: %v", err)
	}

	vals := strings.Split(val, " ") // Split header
	if len(vals) != 2 {             // We expect header to consist only of "ApiKey" and {the value of actual apikey}
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
