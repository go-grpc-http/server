package configs

import "os"

// GetValue retrieves the value of environment variable named by the key.
// It returns the value, if not then it will return the defult value passed
// as the second parameter.
func GetValue(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	// TODO: add logging when passing default value
	return d
}
