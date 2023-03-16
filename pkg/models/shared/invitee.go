package shared

import (
	"time"
)

// InviteeNoShow1
// Provides data pertaining to the associated no show for the Invitee
type InviteeNoShow1 struct {
	CreatedAt time.Time `json:"created_at"`
	URI       string    `json:"uri"`
}

type InviteePaymentCurrencyEnum string

const (
	InviteePaymentCurrencyEnumAud InviteePaymentCurrencyEnum = "AUD"
	InviteePaymentCurrencyEnumCad InviteePaymentCurrencyEnum = "CAD"
	InviteePaymentCurrencyEnumEur InviteePaymentCurrencyEnum = "EUR"
	InviteePaymentCurrencyEnumGbp InviteePaymentCurrencyEnum = "GBP"
	InviteePaymentCurrencyEnumUsd InviteePaymentCurrencyEnum = "USD"
)

type InviteePaymentProviderEnum string

const (
	InviteePaymentProviderEnumStripe InviteePaymentProviderEnum = "stripe"
	InviteePaymentProviderEnumPaypal InviteePaymentProviderEnum = "paypal"
)

// InviteePayment
// Invitee payment
type InviteePayment struct {
	Amount     float32                    `json:"amount"`
	Currency   InviteePaymentCurrencyEnum `json:"currency"`
	ExternalID string                     `json:"external_id"`
	Provider   InviteePaymentProviderEnum `json:"provider"`
	Successful bool                       `json:"successful"`
	Terms      string                     `json:"terms"`
}

// InviteeReconfirmation
// Assuming reconfirmation is enabled for the event type, when reconfirmation is requested this object is present with a `created_at` that reflects when the reconfirmation notification was sent. Once the invitee has reconfirmed the `confirmed_at` attribute will change from `null` to a timestamp that reflects when they took action.
type InviteeReconfirmation struct {
	ConfirmedAt time.Time `json:"confirmed_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type InviteeStatusEnum string

const (
	InviteeStatusEnumActive   InviteeStatusEnum = "active"
	InviteeStatusEnumCanceled InviteeStatusEnum = "canceled"
)

// Invitee
// An individual who has been invited to meet with a Calendly member
type Invitee struct {
	CancelURL             string                     `json:"cancel_url"`
	Cancellation          *Cancellation              `json:"cancellation,omitempty"`
	CreatedAt             time.Time                  `json:"created_at"`
	Email                 string                     `json:"email"`
	Event                 string                     `json:"event"`
	FirstName             string                     `json:"first_name"`
	LastName              string                     `json:"last_name"`
	Name                  string                     `json:"name"`
	NewInvitee            string                     `json:"new_invitee"`
	NoShow                InviteeNoShow1             `json:"no_show"`
	OldInvitee            string                     `json:"old_invitee"`
	Payment               InviteePayment             `json:"payment"`
	QuestionsAndAnswers   []InviteeQuestionAndAnswer `json:"questions_and_answers"`
	Reconfirmation        InviteeReconfirmation      `json:"reconfirmation"`
	RescheduleURL         string                     `json:"reschedule_url"`
	Rescheduled           bool                       `json:"rescheduled"`
	RoutingFormSubmission string                     `json:"routing_form_submission"`
	Status                InviteeStatusEnum          `json:"status"`
	TextReminderNumber    string                     `json:"text_reminder_number"`
	Timezone              string                     `json:"timezone"`
	Tracking              InviteeTracking            `json:"tracking"`
	UpdatedAt             time.Time                  `json:"updated_at"`
	URI                   string                     `json:"uri"`
}
