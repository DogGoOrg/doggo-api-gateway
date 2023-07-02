package helpers

type ResponseWrapper struct {
	Status bool        `json:"status"`
	Error  error       `json:"error"`
	Data   interface{} `json:"data"`
}
