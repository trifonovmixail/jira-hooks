package objects

type BoardConfigurationFilter struct {
	ID   string `json:"id"`
	Self string `json:"self"`
}

type BoardConfigurationColumnStatus struct {
	ID   string `json:"id"`
	Self string `json:"self"`
}

type BoardConfigurationColumn struct {
	Name   string                           `json:"name"`
	Status []BoardConfigurationColumnStatus `json:"statuses"`
}

type BoardConfigurationColumnConfig struct {
	Columns        []BoardConfigurationColumn `json:"columns"`
	ConstraintType string                     `json:"constraintType"`
}

type BoardConfigurationRanking struct {
	RankCustomFieldID int `json:"rankCustomFieldId"`
}

type BoardConfigurationEstimationField struct {
	FieldID     string `json:"fieldId"`
	DisplayName string `json:"displayName"`
}

type BoardConfigurationEstimation struct {
	Type string `json:"type"`
}

type BoardConfiguration struct {
	ID           int                            `json:"id"`
	Name         string                         `json:"name"`
	Self         string                         `json:"self"`
	Type         string                         `json:"type"`
	Filter       BoardConfigurationFilter       `json:"filter"`
	Ranking      BoardConfigurationRanking      `json:"ranking"`
	Estimation   BoardConfigurationEstimation   `json:"estimation"`
	ColumnConfig BoardConfigurationColumnConfig `json:"columnConfig"`
}

type Board struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Self string `json:"self"`
	Type string `json:"type"`
}
