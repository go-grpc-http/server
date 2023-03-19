package env

import (
	"fmt"
	"os"
)

// GetValue retrieves the value of environment variable named by the key.
// It returns the value, if not then it will return the default value passed
// as the second parameter.
var GetValue = func(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	// TODO: add logging when passing default value
	return d
}

// requiredFields - list of fields which are mandates
// unrequiredFields - list of fields which are not mandates
// returns a map with rFields values and uFields values
func EnvData(requiredFields, unrequiredFields []string) map[string]string {
	result := map[string]string{}

	for _, eKey := range requiredFields {
		if eVal := os.Getenv(eKey); eVal != "" {
			result[eKey] = eVal
		} else {
			panic(fmt.Sprintf("missing required env: %s", eKey))
		}
	}

	for _, eKey := range unrequiredFields {
		if eVal := os.Getenv(eKey); eVal != "" {
			result[eKey] = eVal
		}
	}

	return result
}
