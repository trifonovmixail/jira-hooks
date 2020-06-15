package jira_hooks

import (
	"errors"
)

var (
	ErrEventNotFound            = errors.New("event not found")
	ErrParsingPayload           = errors.New("error parsing payload")
	ErrEventNotSpecifiedToParse = errors.New("no event specified to parse")
)
