// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// InviteeNoShow1 - Provides data pertaining to the associated no show for the Invitee
type InviteeNoShow1 struct {
	// The moment when the no show was created
	CreatedAt time.Time `json:"created_at"`
	// Canonical reference (unique identifier) for the no show
	URI string `json:"uri"`
}

// InviteePaymentCurrencyEnum - The currency format that the payment is in.
type InviteePaymentCurrencyEnum string

const (
	InviteePaymentCurrencyEnumAud InviteePaymentCurrencyEnum = "AUD"
	InviteePaymentCurrencyEnumCad InviteePaymentCurrencyEnum = "CAD"
	InviteePaymentCurrencyEnumEur InviteePaymentCurrencyEnum = "EUR"
	InviteePaymentCurrencyEnumGbp InviteePaymentCurrencyEnum = "GBP"
	InviteePaymentCurrencyEnumUsd InviteePaymentCurrencyEnum = "USD"
)

func (e *InviteePaymentCurrencyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "AUD":
		fallthrough
	case "CAD":
		fallthrough
	case "EUR":
		fallthrough
	case "GBP":
		fallthrough
	case "USD":
		*e = InviteePaymentCurrencyEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteePaymentCurrencyEnum: %s", s)
	}
}

// InviteePaymentProviderEnum - Payment provider
type InviteePaymentProviderEnum string

const (
	InviteePaymentProviderEnumStripe InviteePaymentProviderEnum = "stripe"
	InviteePaymentProviderEnumPaypal InviteePaymentProviderEnum = "paypal"
)

func (e *InviteePaymentProviderEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "stripe":
		fallthrough
	case "paypal":
		*e = InviteePaymentProviderEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteePaymentProviderEnum: %s", s)
	}
}

// InviteePayment - Invitee payment
type InviteePayment struct {
	// The amount of the payment
	Amount float32 `json:"amount"`
	// The currency format that the payment is in.
	Currency InviteePaymentCurrencyEnum `json:"currency"`
	// Unique identifier for the payment
	ExternalID string `json:"external_id"`
	// Payment provider
	Provider InviteePaymentProviderEnum `json:"provider"`
	// Indicates whether the payment was successfully processed
	Successful bool `json:"successful"`
	// Terms of the payment
	Terms string `json:"terms"`
}

// InviteeReconfirmation - Assuming reconfirmation is enabled for the event type, when reconfirmation is requested this object is present with a `created_at` that reflects when the reconfirmation notification was sent. Once the invitee has reconfirmed the `confirmed_at` attribute will change from `null` to a timestamp that reflects when they took action.
type InviteeReconfirmation struct {
	// When the Invitee confirmed their attendance.
	ConfirmedAt time.Time `json:"confirmed_at"`
	// When the reconfirmation was created.
	CreatedAt time.Time `json:"created_at"`
}

// InviteeStatusEnum - Indicates if the invitee is "active" or "canceled"
type InviteeStatusEnum string

const (
	InviteeStatusEnumActive   InviteeStatusEnum = "active"
	InviteeStatusEnumCanceled InviteeStatusEnum = "canceled"
)

func (e *InviteeStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "active":
		fallthrough
	case "canceled":
		*e = InviteeStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteeStatusEnum: %s", s)
	}
}

// Invitee - An individual who has been invited to meet with a Calendly member
type Invitee struct {
	// Link to cancelling the event for the invitee
	CancelURL string `json:"cancel_url"`
	// Provides data pertaining to the cancellation of the Event
	Cancellation *Cancellation `json:"cancellation,omitempty"`
	// The moment when the event was created (e.g. "2020-01-02T03:04:05.678123Z")
	CreatedAt time.Time `json:"created_at"`
	// The invitee’s email address
	Email string `json:"email"`
	// A reference to the event
	Event string `json:"event"`
	// The first name of the invitee who booked the event when the event type is configured to use separate fields for first name and last name. Null when event type is configured to use a single field for name.
	FirstName string `json:"first_name"`
	// The last name of the invitee who booked the event when the event type is configured to use separate fields for first name and last name. Null when event type is configured to use a single field for name.
	LastName string `json:"last_name"`
	// The invitee’s name (in human-readable format)
	Name string `json:"name"`
	// Link to new invitee, after reschedule
	NewInvitee string `json:"new_invitee"`
	// Provides data pertaining to the associated no show for the Invitee
	NoShow InviteeNoShow1 `json:"no_show"`
	// Reference to old Invitee instance that got rescheduled
	OldInvitee string `json:"old_invitee"`
	// Invitee payment
	Payment InviteePayment `json:"payment"`
	// A collection of the invitee's responses to questions on the event booking confirmation form
	QuestionsAndAnswers []InviteeQuestionAndAnswer `json:"questions_and_answers"`
	// Assuming reconfirmation is enabled for the event type, when reconfirmation is requested this object is present with a `created_at` that reflects when the reconfirmation notification was sent. Once the invitee has reconfirmed the `confirmed_at` attribute will change from `null` to a timestamp that reflects when they took action.
	Reconfirmation InviteeReconfirmation `json:"reconfirmation"`
	// Link to rescheduling the event for the invitee
	RescheduleURL string `json:"reschedule_url"`
	// Indicates if this invitee has rescheduled. If `true`, a reference to the new Invitee instance is provided in the `new_invitee` field.
	Rescheduled bool `json:"rescheduled"`
	// Reference to a routing form submission that redirected the invitee to a booking page.
	RoutingFormSubmission string `json:"routing_form_submission"`
	// Indicates if the invitee is "active" or "canceled"
	Status InviteeStatusEnum `json:"status"`
	// The phone number to use when sending text (SMS) reminders
	TextReminderNumber string `json:"text_reminder_number"`
	// Time zone to use when displaying time to the invitee
	Timezone string `json:"timezone"`
	// The UTM and Salesforce tracking parameters associated with an Invitee
	Tracking InviteeTracking `json:"tracking"`
	// The moment when the event was last updated (e.g. "2020-01-02T03:04:05.678123Z")
	UpdatedAt time.Time `json:"updated_at"`
	// Canonical reference (unique identifier) for the invitee
	URI string `json:"uri"`
}
