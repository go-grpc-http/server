package version

type Data struct {
	Version     string `json:"version"`
	ProjectName string `json:"projectName"`
	ModelName   string `json:"modelName"`
}

type VersionResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       Data   `json:"data"`
}
