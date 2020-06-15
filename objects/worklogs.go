package objects

type WorkLogRecord struct {
	Self             string `json:"self"`
	Author           *User  `json:"author"`
	UpdateAuthor     *User  `json:"updateAuthor"`
	Comment          string `json:"comment"`
	Created          Time   `json:"created"`
	Updated          Time   `json:"updated"`
	Started          Time   `json:"started"`
	TimeSpent        string `json:"timeSpent"`
	TimeSpentSeconds int    `json:"timeSpentSeconds"`
	ID               string `json:"id"`
	IssueID          string `json:"issueId"`
}

type WorkLog struct {
	StartAt    int              `json:"startAt"`
	MaxResults int              `json:"maxResults"`
	Total      int              `json:"total"`
	WorkLogs   []*WorkLogRecord `json:"worklogs"`
}
