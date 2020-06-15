package jira_hooks

type hook struct {
	Event  Event  `json:"webhookEvent"`
	Action action `json:"issue_event_type_name"`
}

func (h *hook) isEmpty() bool {
	return h.Event == "" && h.Action == ""
}
