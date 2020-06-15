package objects

type Project struct {
	Self           string      `json:"self"`
	ID             string      `json:"id"`
	Key            string      `json:"key"`
	Name           string      `json:"name"`
	AvatarUrls     *AvatarUrls `json:"avatarUrls"`
	ProjectTypeKey string      `json:"projectTypeKey"`
	ProjectLead    *User       `json:"projectLead"`
	AssigneeType   string      `json:"assigneeType"`
}
