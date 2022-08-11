package jira_hooks

import (
	"errors"
	"fmt"
)

var (
	ErrEventNotFound            = errors.New("event not found")
	ErrEventNotSpecifiedToParse = errors.New("no event specified to parse")
)

type ErrParsingPayload error

func NewErrParsingPayload(err interface{}) ErrParsingPayload {
	return errors.New(fmt.Sprintf("error parsing payload,\n description: %s", err))
}
