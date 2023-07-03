package helpers

type ResponseWrapper struct {
	Success bool        `json:"success"`
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
}
