package response

type StatusMessage struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Service     string `json:"service"`
	Description string `json:"description"`
	Error       *error `json:"error"`
}

type ResponseMessage struct {
	Status  StatusMessage `json:"status"`
	Payload interface{}   `json:"payload,omitempty"`
}
