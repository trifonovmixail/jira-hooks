package objects

type Comment struct {
	ID           string `json:"id"`
	Self         string `json:"self"`
	Name         string `json:"name"`
	Author       *User  `json:"author"`
	Body         string `json:"body"`
	UpdateAuthor *User  `json:"updateAuthor"`
	Updated      string `json:"updated"`
	Created      string `json:"created"`
}

type Comments struct {
	MaxResults int        `json:"maxResults"`
	Total      int        `json:"total"`
	StartAt    int        `json:"startAt"`
	Comments   []*Comment `json:"comments"`
}
