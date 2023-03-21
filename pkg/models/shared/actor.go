package shared

// ActorGroup
// User group information about the actor
type ActorGroup struct {
	Name *string `json:"name,omitempty"`
	Role *string `json:"role,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type ActorOrganization struct {
	Role *string `json:"role,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

// Actor
// The Calendly actor that took the action creating the activity log entry
//
// Specific actors:
//
// <details>
// <summary>Calendly System</summary>
//
// Used when an action is performed by the Calendly system and not triggered directly by a user interaction.
//
// Example:
// ```json
//
//	{
//	    "display_name": "Calendly System",
//	    "type": "System"
//	}
//
// ```
//
// </details>
//
// <br />
//
// <details>
// <summary>Calendly Support</summary>
// Used when an action is performed by Calendly support.
//
// Example:
// ```json
//
//	{
//	  "display_name": "Calendly Support",
//	  "organization": {
//	    "uri": "https://api.calendly.com/organizations/AAAAAAAAAAAAAAAA"
//	  },
//	  "type": "User",
//	  "uri": "https://api.calendly.com/users/AAAAAAAAAAAAAAAA"
//	}
//
// ```
// </details>
type Actor struct {
	AlternativeIdentifier *string            `json:"alternative_identifier,omitempty"`
	DisplayName           *string            `json:"display_name,omitempty"`
	Group                 *ActorGroup        `json:"group,omitempty"`
	Organization          *ActorOrganization `json:"organization,omitempty"`
	Type                  string             `json:"type"`
	URI                   *string            `json:"uri,omitempty"`
}
