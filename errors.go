package jira_hooks

import (
	"errors"
	"fmt"
)

type ErrorParsingPayload struct {
	Message      string
	ParsingError error
}

func (e ErrorParsingPayload) Error() string {
	return fmt.Sprintf("%s,\n description: %s", e.Message, e.ParsingError)
}

var (
	ErrEventNotFound            = errors.New("event not found")
	ErrParsingPayload           = ErrorParsingPayload{Message: "error parsing payload"}
	ErrEventNotSpecifiedToParse = errors.New("no event specified to parse")
)
