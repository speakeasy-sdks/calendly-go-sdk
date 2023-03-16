package shared

type SchemeOauth2 struct {
	Authorization string `security:"name=Authorization"`
}

type SchemePersonalAccessToken struct {
	Authorization string `security:"name=Authorization"`
}

type Security struct {
	Oauth2              *SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}
