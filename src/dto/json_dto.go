package dto

type JSONResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJSONResult(message string, data interface{}) *JSONResult {
	return &JSONResult{Message: message, Data: data}
}
