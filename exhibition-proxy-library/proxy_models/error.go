package proxy_models

type Error struct{
	ErrorMessage string `json:"error_message"`
	StatusCode int `json:"status_code"`
}