package objects

type User struct {
	Self         string      `json:"self"`
	Name         string      `json:"name"`
	Key          string      `json:"key"`
	EmailAddress string      `json:"emailAddress"`
	AvatarUrls   *AvatarUrls `json:"avatarUrls"`
	DisplayName  string      `json:"displayName"`
	Active       bool        `json:"active"`
	TimeZone     string      `json:"timeZone"`
}
