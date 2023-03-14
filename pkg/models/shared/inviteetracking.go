package shared

// InviteeTracking
// The UTM and Salesforce tracking parameters associated with an Invitee
type InviteeTracking struct {
	SalesforceUUID string `json:"salesforce_uuid"`
	UtmCampaign    string `json:"utm_campaign"`
	UtmContent     string `json:"utm_content"`
	UtmMedium      string `json:"utm_medium"`
	UtmSource      string `json:"utm_source"`
	UtmTerm        string `json:"utm_term"`
}
