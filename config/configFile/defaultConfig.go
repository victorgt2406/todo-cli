package configFile

func defaultConfig() ConfigFile {
	return ConfigFile{
		LlmProvider: LlmProvider{
			Url:    "",
			Model:  "",
			ApiKey: "",
		},
		Features: Features{
			SmartTask: false,
		},
	}
}
