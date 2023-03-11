package shared

// InviteeQuestionAndAnswer
// A response to a question on a booking page form
type InviteeQuestionAndAnswer struct {
	Answer   string  `json:"answer"`
	Position float64 `json:"position"`
	Question string  `json:"question"`
}
