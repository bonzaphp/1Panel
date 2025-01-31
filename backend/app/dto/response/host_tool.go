package response

type HostToolRes struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

type Supervisor struct {
	ConfigPath string `json:"configPath"`
	IncludeDir string `json:"includeDir"`
	LogPath    string `json:"logPath"`
	IsExist    bool   `json:"isExist"`
	Init       bool   `json:"init"`
	Msg        string `json:"msg"`
	Version    string `json:"version"`
	Status     string `json:"status"`
	CtlExist   bool   `json:"ctlExist"`
}

type HostToolConfig struct {
	Content string `json:"content"`
}
