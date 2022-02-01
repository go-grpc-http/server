package resources

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Error      error  `json:"error"`
}
