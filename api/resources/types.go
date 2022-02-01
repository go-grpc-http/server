package resources

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Error      struct {
		Message string `json:"message"`
	} `json:"error"`
}
