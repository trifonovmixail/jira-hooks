package jira_hooks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createHookFromRequest(request *http.Request) (*hook, []byte, error) {
	errorParsingPayLoad := ErrParsingPayload

	payload, err := ioutil.ReadAll(request.Body)
	if err != nil || len(payload) == 0 {
		errorParsingPayLoad.ParsingError = err
		return nil, nil, errorParsingPayLoad
	}
	if DebugRequest {
		debugRequest(payload)
	}

	hook := &hook{}
	if err := unmarshalJson(payload, hook); err != nil {
		errorParsingPayLoad.ParsingError = err
		return nil, nil, errorParsingPayLoad
	}
	if hook.isEmpty() && isTransitionIssueStatusPayload(payload) {
		hook.Event = StatusTransitionEvent
	}

	return hook, payload, nil
}

func parsing(hook *hook, payload []byte) (interface{}, error) {
	var err error
	errorParsingPayLoad := ErrParsingPayload
	switch hook.Event {
	case StatusTransitionEvent:
		var pl TransitionIssueStatusPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			if !pl.isEmpty() {
				return pl, nil
			}
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case IssueCreatedEvent:
		var pl IssueCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case IssueUpdatedEvent:
		switch hook.Action {
		case issueGenericAction:
			var pl IssueGenericPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueUpdatedAction:
			var pl IssueUpdatedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueCommentCreatedAction:
			var pl IssueCommentCreatedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueCommentUpdatedAction:
			var pl IssueCommentUpdatedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueCommentDeletedAction:
			var pl IssueCommentDeletedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueAssignedAction:
			var pl IssueAssignedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueMovedAction:
			var pl IssueMovedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueClosedAction:
			var pl IssueClosedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		}
	case IssueDeletedEvent:
		var pl IssueDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case IssueWorkLogEvent:
		switch hook.Action {
		case issueWorkLogCreatedAction:
			var pl IssueWorkLogCreatedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueWorkLogUpdatedAction:
			var pl IssueWorkLogUpdatedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		case issueWorkLogDeletedAction:
			var pl IssueWorkLogDeletedPayload
			if err = unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			errorParsingPayLoad.ParsingError = err
			return nil, errorParsingPayLoad
		}
	case CommentCreatedEvent:
		var pl CommentCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case CommentUpdatedEvent:
		var pl CommentUpdatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case CommentDeletedEvent:
		var pl CommentDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case WorkLogCreatedEvent:
		var pl WorkLogCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case WorkLogUpdatedEvent:
		var pl WorkLogUpdatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case WorkLogDeletedEvent:
		var pl WorkLogDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case LinkCreatedEvent:
		var pl LinkCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case LinkDeletedEvent:
		var pl LinkDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case UserCreatedEvent:
		var pl UserCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case UserUpdatedEvent:
		var pl UserUpdatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case UserDeletedEvent:
		var pl UserDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case ProjectCreatedEvent:
		var pl ProjectCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case ProjectUpdatedEvent:
		var pl ProjectUpdatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case ProjectDeletedEvent:
		var pl ProjectDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case ProjectArchivedEvent:
		var pl ProjectArchivedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case ProjectRestoredEvent:
		var pl ProjectRestoredPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case BoardCreatedEvent:
		var pl BoardCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case BoardUpdatedEvent:
		var pl BoardUpdatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case BoardDeletedEvent:
		var pl BoardDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case BoardConfigurationChangedEvent:
		var pl BoardConfigurationChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case SprintCreatedEvent:
		var pl SprintCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case SprintUpdatedEvent:
		var pl SprintUpdatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case SprintDeletedEvent:
		var pl SprintDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case SprintStartedEvent:
		var pl SprintStartedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case SprintClosedEvent:
		var pl SprintClosedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case VersionCreatedEvent:
		var pl VersionCreatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case VersionUpdatedEvent:
		var pl VersionUpdatedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case VersionDeletedEvent:
		var pl VersionDeletedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case VersionReleasedEvent:
		var pl VersionReleasedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case VersionUnreleasedEvent:
		var pl VersionUnreleasedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case OptionTimeTrackingChangedEvent:
		var pl OptionTimeTrackingChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case OptionIssueLinksChangedEvent:
		var pl OptionIssueLinksChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case OptionSubTasksChangedEvent:
		var pl OptionSubTasksChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case OptionAttachmentsChangedEvent:
		var pl OptionAttachmentsChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case OptionWatchingChangedEvent:
		var pl OptionWatchingChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case OptionVotingChangedEvent:
		var pl OptionVotingChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	case OptionUnassignedIssuesChangedEvent:
		var pl OptionUnassignedIssuesChangedPayload
		if err = unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		errorParsingPayLoad.ParsingError = err
		return nil, errorParsingPayLoad
	}

	return nil, errors.New(
		fmt.Sprintf("unknown event '%s' with action: '%s'", hook.Event, hook.Action),
	)
}

func Parse(request *http.Request, events ...Event) (interface{}, error) {
	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}

	hook, payload, err := createHookFromRequest(request)
	if err != nil {
		return nil, err
	}

	found := false
	for _, event := range events {
		if event == hook.Event {
			found = true
			break
		}
	}
	if !found {
		return nil, ErrEventNotFound
	}

	return parsing(hook, payload)
}
