package logger

type LogInfos struct {
	LogID   string `json:"type"`
	Service string `json:"service"`
	Type    string `json:"type"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
