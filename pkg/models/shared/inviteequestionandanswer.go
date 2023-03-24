// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// InviteeQuestionAndAnswer - A response to a question on a booking page form
type InviteeQuestionAndAnswer struct {
	// The invitee's answer to the question
	Answer string `json:"answer"`
	// The position of the question in relation to others on the booking form
	Position float64 `json:"position"`
	// A question on the invitee's booking form
	Question string `json:"question"`
}
