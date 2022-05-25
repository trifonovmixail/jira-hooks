package objects

type LinkType struct {
	ID          string `json:"id"`
	Self        string `json:"self"`
	Name        string `json:"name"`
	Inward      string `json:"inward"`
	Outward     string `json:"outward"`
	InwardName  string `json:"inwardName"`
	OutwardName string `json:"outwardName"`
	SubTask     bool   `json:"subTask"`
	System      bool   `json:"system"`
}

type Link struct {
	ID                 string    `json:"id"`
	Self               string    `json:"self"`
	OutwardIssue       *Issue    `json:"outwardIssue"`
	SourceIssueID      int       `json:"sourceIssueId"`
	DestinationIssueID int       `json:"destinationIssueId"`
	Type               *LinkType `json:"issueLinkType"`
	SystemLink         bool      `json:"systemLink"`
}

type LinkPayloadLinkType struct {
	ID int `json:"id"`
	LinkType
}

type LinkPayloadLink struct {
	ID   int                  `json:"id"`
	Type *LinkPayloadLinkType `json:"issueLinkType"`
	Link
}
