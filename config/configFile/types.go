package configFile

type LlmProvider struct {
	Url    string `json:"url"`
	Model  string `json:"model"`
	ApiKey string `json:"api_key"`
}

type Features struct {
	SmartTask bool `json:"smart_task"`
}

type ConfigFile struct {
	LlmProvider LlmProvider `json:"llm_provider"`
	Features    Features    `json:"features"`
}
