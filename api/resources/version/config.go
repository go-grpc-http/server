package version

type VersionResponseData struct {
	Version     string `json:"version"`
	ProjectName string `json:"projectName"`
	ModelName   string `json:"modelName"`
}

type Config struct {
	projectName string
	modelName   string
	version     string
}

func New(pName, mName, version string) *Config {
	return &Config{
		projectName: pName,
		modelName:   mName,
		version:     version,
	}
}
