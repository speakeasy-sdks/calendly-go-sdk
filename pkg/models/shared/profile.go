package shared

type ProfileTypeEnum string

const (
	ProfileTypeEnumUser ProfileTypeEnum = "User"
	ProfileTypeEnumTeam ProfileTypeEnum = "Team"
)

// Profile
// The publicly visible profile of a User or a Team that's associated with the Event Type (note: some Event Types don't have profiles)
type Profile struct {
	Name  string          `json:"name"`
	Owner string          `json:"owner"`
	Type  ProfileTypeEnum `json:"type"`
}
