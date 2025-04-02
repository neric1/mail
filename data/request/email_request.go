package request

type EmailRequestBody struct {
	Recipients []string `json:"recipients"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
	// From string `json:"from"`
}
