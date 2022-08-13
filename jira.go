package jira_hooks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createHookFromRequest(request *http.Request) (*hook, []byte, error) {
	payload, err := ioutil.ReadAll(request.Body)
	if err != nil || len(payload) == 0 {
		return nil, nil, ErrParsingPayload
	}
	if DebugRequest {
		debugRequest(payload)
	}

	hook := &hook{}
	if err := unmarshalJson(payload, hook); err != nil {
		return nil, nil, ErrParsingPayload
	}
	if hook.isEmpty() && isTransitionIssueStatusPayload(payload) {
		hook.Event = StatusTransitionEvent
	}

	return hook, payload, nil
}

func parsing(hook *hook, payload []byte) (interface{}, error) {
	switch hook.Event {
	case StatusTransitionEvent:
		var pl TransitionIssueStatusPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			if !pl.isEmpty() {
				return pl, nil
			}
		}
		return nil, ErrParsingPayload
	case IssueCreatedEvent:
		var pl IssueCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case IssueUpdatedEvent:
		switch hook.Action {
		case issueGenericAction:
			var pl IssueGenericPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueUpdatedAction:
			var pl IssueUpdatedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueCommentCreatedAction:
			var pl IssueCommentCreatedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueCommentUpdatedAction:
			var pl IssueCommentUpdatedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueCommentDeletedAction:
			var pl IssueCommentDeletedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueAssignedAction:
			var pl IssueAssignedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueMovedAction:
			var pl IssueMovedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueClosedAction:
			var pl IssueClosedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueResolvedAction:
			var pl IssueResolvedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		}
	case IssueDeletedEvent:
		var pl IssueDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case IssueWorkLogEvent:
		switch hook.Action {
		case issueWorkLogCreatedAction:
			var pl IssueWorkLogCreatedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueWorkLogUpdatedAction:
			var pl IssueWorkLogUpdatedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		case issueWorkLogDeletedAction:
			var pl IssueWorkLogDeletedPayload
			if err := unmarshalJson(payload, &pl); err == nil {
				return pl, nil
			}
			return nil, ErrParsingPayload
		}
	case CommentCreatedEvent:
		var pl CommentCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case CommentUpdatedEvent:
		var pl CommentUpdatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case CommentDeletedEvent:
		var pl CommentDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case WorkLogCreatedEvent:
		var pl WorkLogCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case WorkLogUpdatedEvent:
		var pl WorkLogUpdatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case WorkLogDeletedEvent:
		var pl WorkLogDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case LinkCreatedEvent:
		var pl LinkCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case LinkDeletedEvent:
		var pl LinkDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case UserCreatedEvent:
		var pl UserCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case UserUpdatedEvent:
		var pl UserUpdatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case UserDeletedEvent:
		var pl UserDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case ProjectCreatedEvent:
		var pl ProjectCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case ProjectUpdatedEvent:
		var pl ProjectUpdatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case ProjectDeletedEvent:
		var pl ProjectDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case ProjectArchivedEvent:
		var pl ProjectArchivedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case ProjectRestoredEvent:
		var pl ProjectRestoredPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case BoardCreatedEvent:
		var pl BoardCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case BoardUpdatedEvent:
		var pl BoardUpdatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case BoardDeletedEvent:
		var pl BoardDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case BoardConfigurationChangedEvent:
		var pl BoardConfigurationChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case SprintCreatedEvent:
		var pl SprintCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case SprintUpdatedEvent:
		var pl SprintUpdatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case SprintDeletedEvent:
		var pl SprintDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case SprintStartedEvent:
		var pl SprintStartedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case SprintClosedEvent:
		var pl SprintClosedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case VersionCreatedEvent:
		var pl VersionCreatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case VersionUpdatedEvent:
		var pl VersionUpdatedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case VersionDeletedEvent:
		var pl VersionDeletedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case VersionReleasedEvent:
		var pl VersionReleasedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case VersionUnreleasedEvent:
		var pl VersionUnreleasedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case OptionTimeTrackingChangedEvent:
		var pl OptionTimeTrackingChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case OptionIssueLinksChangedEvent:
		var pl OptionIssueLinksChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case OptionSubTasksChangedEvent:
		var pl OptionSubTasksChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case OptionAttachmentsChangedEvent:
		var pl OptionAttachmentsChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case OptionWatchingChangedEvent:
		var pl OptionWatchingChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case OptionVotingChangedEvent:
		var pl OptionVotingChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
	case OptionUnassignedIssuesChangedEvent:
		var pl OptionUnassignedIssuesChangedPayload
		if err := unmarshalJson(payload, &pl); err == nil {
			return pl, nil
		}
		return nil, ErrParsingPayload
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
