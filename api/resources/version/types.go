package version

type VersionResponseData struct {
	Version     string `json:"version"`
	ProjectName string `json:"projectName"`
	ModelName   string `json:"modelName"`
}
