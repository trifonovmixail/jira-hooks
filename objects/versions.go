package objects

type Version struct {
	Self            string `json:"self"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Archived        bool   `json:"archived"`
	Released        bool   `json:"released"`
	Overdue         bool   `json:"overdue"`
	ProjectID       int    `json:"projectId"`
	UserStartDate   string `json:"userStartDate"`
	UserReleaseDate string `json:"userReleaseDate"`
}

type FixVersion Version
type AffectsVersion Version
