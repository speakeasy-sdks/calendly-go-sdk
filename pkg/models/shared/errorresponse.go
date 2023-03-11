package shared

type ErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// ErrorResponse
// Error Object
type ErrorResponse struct {
	Details []ErrorResponseDetails `json:"details,omitempty"`
	Message string                 `json:"message"`
	Title   string                 `json:"title"`
}
