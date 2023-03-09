package shared

import (
	"time"
)

type OrganizationMembershipRoleEnum string

const (
	OrganizationMembershipRoleEnumUser  OrganizationMembershipRoleEnum = "user"
	OrganizationMembershipRoleEnumAdmin OrganizationMembershipRoleEnum = "admin"
	OrganizationMembershipRoleEnumOwner OrganizationMembershipRoleEnum = "owner"
)

// OrganizationMembershipUser
// Information about the user.
type OrganizationMembershipUser struct {
	AvatarURL     string    `json:"avatar_url"`
	CreatedAt     time.Time `json:"created_at"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	SchedulingURL string    `json:"scheduling_url"`
	Slug          string    `json:"slug"`
	Timezone      string    `json:"timezone"`
	UpdatedAt     time.Time `json:"updated_at"`
	URI           string    `json:"uri"`
}

type OrganizationMembership struct {
	CreatedAt    time.Time                      `json:"created_at"`
	Organization string                         `json:"organization"`
	Role         OrganizationMembershipRoleEnum `json:"role"`
	UpdatedAt    time.Time                      `json:"updated_at"`
	URI          string                         `json:"uri"`
	User         OrganizationMembershipUser     `json:"user"`
}
