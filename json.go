package jira_hooks

import (
	"encoding/json"
	"log"
)

type jsonString string

func (j *jsonString) UnmarshalJSON(b []byte) error {
	*j = jsonString(b)
	return nil
}

func unmarshalJson(b []byte, i interface{}) error {
	err := json.Unmarshal(b, i)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	return nil
}
