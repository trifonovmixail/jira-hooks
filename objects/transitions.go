package objects

type StatusTransition struct {
	WorkFlowID     int    `json:"workflowId"`
	WorkFlowName   string `json:"workflowName"`
	TransitionID   int    `json:"transitionId"`
	TransitionName string `json:"transitionName"`
	FromStatus     string `json:"from_status"`
	ToStatus       string `json:"to_status"`
}
