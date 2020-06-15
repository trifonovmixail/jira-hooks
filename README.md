# Install

```bash
go get github.com/trifonovmixail/jira-hooks
```

# Simple example

```go
import (
    hooks "github.com/trifonovmixail/jira-hooks"
    "net/http"
)

func manyHooks(request *http.Request) int {
    hook, err := hooks.Parse(
        request,
        hooks.IssueCreatedEvent,
        hooks.StatusTransitionEvent,
    )
    if err != nil {
        return http.StatusBadRequest
    }

    switch hook.(type) {
    case hooks.CreateIssuePayload:
        payload := hook.(hooks.CreateIssuePayload)
    case hooks.TransitionIssueStatusPayload:
        payload := hook.(hooks.TransitionIssueStatusPayload)
    }

    return http.StatusBadRequest
}

func oneHook(request *http.Request) int {
    hook, err := hooks.Parse(request, hooks.StatusTransitionEvent)
    if err != nil {
        return http.StatusBadRequest
    }

    payload, ok := hook.(hooks.TransitionIssueStatusPayload)
    if !ok {
        return http.StatusBadRequest
    }

    return http.StatusOK
}
```

# Events

## Transitions

### StatusTransitionEvent
Payloads:
* TransitionIssueStatusPayload


## Issues

### IssueCreatedEvent
Payloads:
* CreateIssuePayload

### IssueCreatedEvent
Payloads:
* IssueGenericPayload
* IssueUpdatedPayload
* IssueCommentCreatedPayload
* IssueCommentUpdatedPayload
* IssueCommentDeletedPayload
* IssueAssignedPayload

### IssueDeletedEvent
Payloads:
* DeleteIssuePayload

### IssueWorkLogEvent
Payloads:
* IssueWorkLogCreatedPayload
* IssueWorkLogUpdatedPayload
* IssueWorkLogDeletedPayload


## Comments

### CommentCreatedEvent
Payloads:
* CommentCreatedPayload

### CommentUpdatedEvent
Payloads:
* CommentUpdatedPayload

### CommentDeletedEvent
Payloads:
* CommentDeletedPayload


## Work logs

### WorkLogCreatedEvent
Payloads:
* WorkLogCreatedPayload

### WorkLogUpdatedEvent
Payloads:
* WorkLogUpdatedPayload

### WorkLogDeletedEvent
Payloads:
* WorkLogDeletedPayload


## Links

### LinkCreatedEvent
Payloads:
* LinkCreatedPayload

### LinkDeletedEvent
Payloads:
* LinkDeletedPayload


## Users

### UserCreatedEvent
Payloads:
* UserCreatedPayload

### UserUpdatedEvent
Payloads:
* UserUpdatedPayload

### UserDeletedEvent
Payloads:
* UserDeletedPayload


## Projects

### ProjectCreatedEvent
Payloads:
* ProjectCreatedPayload

### ProjectUpdatedEvent
Payloads:
* ProjectUpdatedPayload

### ProjectDeletedEvent
Payloads:
* ProjectDeletedPayload

### ProjectArchivedEvent
Payloads:
* ProjectArchivedPayload

### ProjectRestoredEvent
Payloads:
* ProjectRestoredPayload


## Boards

### BoardCreatedEvent
Payloads:
* BoardCreatedPayload

### BoardUpdatedEvent
Payloads:
* BoardUpdatedPayload

### BoardDeletedEvent
Payloads:
* BoardDeletedPayload

### BoardConfigurationChangedEvent
Payloads:
* BoardConfigurationChangedPayload


## Sprints

### SprintCreatedEvent
Payloads:
* SprintCreatedPayload

### SprintUpdatedEvent
Payloads:
* SprintUpdatedPayload

### SprintDeletedEvent
Payloads:
* SprintDeletedPayload

### SprintStartedEvent
Payloads:
* SprintStartedPayload

### SprintClosedEvent
Payloads:
* SprintClosedPayload


## Versions

### VersionCreatedEvent
Payloads:
* VersionCreatedPayload

### VersionUpdatedEvent
Payloads:
* VersionUpdatedPayload

### VersionDeletedEvent
Payloads:
* VersionDeletedPayload

### VersionReleasedEvent
Payloads:
* VersionReleasedPayload

### VersionUnreleasedEvent
Payloads:
* VersionUnreleasedPayload


## System options

### OptionTimeTrackingChangedEvent
Payloads:
* OptionTimeTrackingChangedPayload

### OptionIssueLinksChangedEvent
Payloads:
* OptionIssueLinksChangedPayload

### OptionSubTasksChangedEvent
Payloads:
* OptionSubTasksChangedPayload

### OptionAttachmentsChangedEvent
Payloads:
* OptionAttachmentsChangedPayload

### OptionWatchingChangedEvent
Payloads:
* OptionWatchingChangedPayload

### OptionVotingChangedEvent
Payloads:
* OptionVotingChangedPayload

### OptionUnassignedIssuesChangedEvent
Payloads:
* OptionUnassignedIssuesChangedPayload
