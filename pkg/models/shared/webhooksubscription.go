package shared

import (
	"time"
)

type WebhookSubscriptionEventsEnum string

const (
	WebhookSubscriptionEventsEnumInviteeCreated               WebhookSubscriptionEventsEnum = "invitee.created"
	WebhookSubscriptionEventsEnumInviteeCanceled              WebhookSubscriptionEventsEnum = "invitee.canceled"
	WebhookSubscriptionEventsEnumRoutingFormSubmissionCreated WebhookSubscriptionEventsEnum = "routing_form_submission.created"
)

type WebhookSubscriptionScopeEnum string

const (
	WebhookSubscriptionScopeEnumUser         WebhookSubscriptionScopeEnum = "user"
	WebhookSubscriptionScopeEnumOrganization WebhookSubscriptionScopeEnum = "organization"
)

type WebhookSubscriptionStateEnum string

const (
	WebhookSubscriptionStateEnumActive   WebhookSubscriptionStateEnum = "active"
	WebhookSubscriptionStateEnumDisabled WebhookSubscriptionStateEnum = "disabled"
)

// WebhookSubscription
// Webhook Subscription Object
type WebhookSubscription struct {
	CallbackURL    string                          `json:"callback_url"`
	CreatedAt      time.Time                       `json:"created_at"`
	Creator        string                          `json:"creator"`
	Events         []WebhookSubscriptionEventsEnum `json:"events"`
	Organization   string                          `json:"organization"`
	RetryStartedAt time.Time                       `json:"retry_started_at"`
	Scope          WebhookSubscriptionScopeEnum    `json:"scope"`
	State          WebhookSubscriptionStateEnum    `json:"state"`
	UpdatedAt      time.Time                       `json:"updated_at"`
	URI            string                          `json:"uri"`
	User           string                          `json:"user"`
}
