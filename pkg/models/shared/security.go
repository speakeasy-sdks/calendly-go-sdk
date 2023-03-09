package shared

type SchemeOauth2 struct {
	Authorization string `security:"name=Authorization"`
}

type SchemePersonalAccessToken struct {
	Authorization string `security:"name=Authorization"`
}
