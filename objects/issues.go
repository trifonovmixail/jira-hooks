package objects

type IssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	SubTask     bool   `json:"subtask"`
	AvatarID    int    `json:"avatarId"`
}

type Resolution struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Priority struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

type Watches struct {
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}

type Component struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Status struct {
	Self        string `json:"self"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type TimeTracking struct {
	RemainingEstimate        string `json:"remainingEstimate"`
	TimeSpent                string `json:"timeSpent"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds"`
}

type SubTaskFields struct {
	Type     *IssueType `json:"issuetype"`
	Priority *Priority  `json:"priority"`
	Summary  string     `json:"summary"`
	Status   *Status    `json:"status"`
}

type SubTask struct {
	ID     string         `json:"id"`
	Key    string         `json:"key"`
	Self   string         `json:"self"`
	Fields *SubTaskFields `json:"fields"`
}

type Attachment struct {
	Self      string `json:"self"`
	ID        string `json:"id"`
	Filename  string `json:"filename"`
	Author    *User  `json:"author"`
	Created   string `json:"created"`
	Size      int    `json:"size"`
	MimeType  string `json:"mimeType"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail"`
}

type IssueFields struct {
	Type                          *IssueType        `json:"issuetype"`
	Project                       *Project          `json:"project"`
	Resolution                    *Resolution       `json:"resolution"`
	Priority                      *Priority         `json:"priority"`
	ResolutionDate                Time              `json:"resolutiondate"`
	Created                       Time              `json:"created"`
	DueDate                       Date              `json:"duedate"`
	Watches                       *Watches          `json:"watches"`
	Assignee                      *User             `json:"assignee"`
	Updated                       Time              `json:"updated"`
	Description                   string            `json:"description"`
	Summary                       string            `json:"summary"`
	Creator                       *User             `json:"creator"`
	Reporter                      *User             `json:"reporter"`
	Components                    []*Component      `json:"components"`
	Status                        *Status           `json:"status"`
	Progress                      *Progress         `json:"progress"`
	AggregateProgress             *Progress         `json:"aggregateprogress"`
	TimeTracking                  *TimeTracking     `json:"timetracking"`
	TimeSpent                     int               `json:"timespent"`
	TimeEstimate                  int               `json:"timeestimate"`
	TimeOriginalEstimate          int               `json:"timeoriginalestimate"`
	Worklog                       *WorkLog          `json:"worklog"`
	IssueLinks                    []*Link           `json:"issuelinks"`
	Comments                      *Comments         `json:"comment"`
	FixVersions                   []*FixVersion     `json:"fixVersions"`
	AffectsVersions               []*AffectsVersion `json:"versions"`
	Labels                        []string          `json:"labels"`
	SubTasks                      []*SubTask        `json:"subtasks"`
	Attachments                   []*Attachment     `json:"attachment"`
	AggregateTimeOriginalEstimate int               `json:"aggregatetimeoriginalestimate"`
	AggregateTimeSpent            int               `json:"aggregatetimespent"`
	AggregateTimeEstimate         int               `json:"aggregatetimeestimate"`
}

type Issue struct {
	ID     string       `json:"id"`
	Self   string       `json:"self"`
	Key    string       `json:"key"`
	Fields *IssueFields `json:"fields"`
}
