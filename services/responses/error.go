package responses

// Error will return a error response
func Error(message string) interface{} {
	return struct {
		Error string `json:"error"`
	}{message}
}
