package objects

type ChangeLogItem struct {
	Field      string      `json:"field"`
	FieldType  string      `json:"fieldtype"`
	From       interface{} `json:"from"`
	FromString string      `json:"fromString"`
	To         interface{} `json:"to"`
	ToString   string      `json:"toString"`
}

type ChangeLog struct {
	Id    string           `json:"id"`
	Items []*ChangeLogItem `json:"items"`
}
