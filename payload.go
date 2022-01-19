package jira_hooks

import (
	"github.com/trifonovmixail/jira-hooks/objects"
	"strings"
)

// Transitions

func isTransitionIssueStatusPayload(payload []byte) bool {
	data := string(payload)
	hasTransition := strings.Contains(data, "\"transition\"")
	hasFromStatus := strings.Contains(data, "\"from_status\"")
	hasToStatus := strings.Contains(data, "\"to_status\"")

	return hasTransition && hasFromStatus && hasToStatus
}

type TransitionIssueStatusPayload struct {
	Time       objects.Timestamp         `json:"timestamp"`
	User       *objects.User             `json:"user"`
	Issue      *objects.Issue            `json:"issue"`
	Transition *objects.StatusTransition `json:"transition"`
}

func (s *TransitionIssueStatusPayload) isEmpty() bool {
	hasWorkFlow := s.Transition.WorkFlowID != 0
	hasTransition := s.Transition.TransitionID != 0
	return !(hasWorkFlow && hasTransition)
}

// Issues

type issuePayload struct {
	Time  objects.Timestamp `json:"timestamp"`
	User  *objects.User     `json:"user"`
	Issue *objects.Issue    `json:"issue"`
}

type issueUpdatedPayload struct {
	Time      objects.Timestamp  `json:"timestamp"`
	User      *objects.User      `json:"user"`
	Issue     *objects.Issue     `json:"issue"`
	ChangeLog *objects.ChangeLog `json:"changelog"`
	Comment   *objects.Comment   `json:"comment"`
}

type issueCommentPayload struct {
	Time    objects.Timestamp `json:"timestamp"`
	User    *objects.User     `json:"user"`
	Issue   *objects.Issue    `json:"issue"`
	Comment *objects.Comment  `json:"comment"`
}

type IssueCreatedPayload issuePayload
type IssueDeletedPayload issuePayload

type IssueGenericPayload issueUpdatedPayload
type IssueUpdatedPayload issueUpdatedPayload

type IssueAssignedPayload issueUpdatedPayload

type IssueCommentCreatedPayload issueCommentPayload
type IssueCommentUpdatedPayload issueCommentPayload
type IssueCommentDeletedPayload issuePayload

type IssueWorkLogCreatedPayload issueUpdatedPayload
type IssueWorkLogUpdatedPayload issueUpdatedPayload
type IssueWorkLogDeletedPayload issueUpdatedPayload

type IssueMovedPayload issueUpdatedPayload

type IssueClosedPayload issueUpdatedPayload

// WorkLogs

type workLogPayload struct {
	Time    objects.Timestamp      `json:"timestamp"`
	WorkLog *objects.WorkLogRecord `json:"worklog"`
}

type WorkLogCreatedPayload workLogPayload
type WorkLogUpdatedPayload workLogPayload
type WorkLogDeletedPayload workLogPayload

// Comments

type commentPayload struct {
	Time    objects.Timestamp `json:"timestamp"`
	Comment *objects.Comment  `json:"comment"`
}

type CommentCreatedPayload commentPayload
type CommentUpdatedPayload commentPayload
type CommentDeletedPayload commentPayload

// Links

type linkPayload struct {
	Time objects.Timestamp `json:"timestamp"`
	Link *objects.Link     `json:"issueLink"`
}

type LinkCreatedPayload linkPayload
type LinkDeletedPayload linkPayload

// Users

type userPayload struct {
	Time objects.Timestamp `json:"timestamp"`
	User *objects.User     `json:"user"`
}

type UserCreatedPayload userPayload
type UserUpdatedPayload userPayload
type UserDeletedPayload userPayload

// Projects

type projectPayload struct {
	Time    objects.Timestamp `json:"timestamp"`
	Project *objects.Project  `json:"project"`
}

type ProjectCreatedPayload projectPayload
type ProjectUpdatedPayload projectPayload
type ProjectDeletedPayload projectPayload
type ProjectArchivedPayload projectPayload
type ProjectRestoredPayload projectPayload

// Boards

type BoardConfigurationChangedPayload struct {
	Time          objects.Timestamp           `json:"timestamp"`
	Configuration *objects.BoardConfiguration `json:"configuration"`
}

type boardPayload struct {
	Time  objects.Timestamp `json:"timestamp"`
	Board *objects.Board    `json:"board"`
}

type BoardCreatedPayload boardPayload
type BoardUpdatedPayload boardPayload
type BoardDeletedPayload boardPayload

// Sprints

type sprintPayload struct {
	Time   objects.Timestamp `json:"timestamp"`
	Sprint *objects.Sprint   `json:"sprint"`
}

type SprintCreatedPayload sprintPayload
type SprintUpdatedPayload sprintPayload
type SprintDeletedPayload sprintPayload
type SprintStartedPayload sprintPayload
type SprintClosedPayload sprintPayload

// Versions

type versionPayload struct {
	Time    objects.Timestamp `json:"timestamp"`
	Version *objects.Version  `json:"version"`
}

type VersionCreatedPayload versionPayload
type VersionUpdatedPayload versionPayload
type VersionDeletedPayload versionPayload
type VersionReleasedPayload versionPayload
type VersionUnreleasedPayload versionPayload

// Options

type optionChangedPayload struct {
	Time     objects.Timestamp `json:"timestamp"`
	Property *objects.Property `json:"property"`
}

type OptionTimeTrackingChangedPayload optionChangedPayload
type OptionIssueLinksChangedPayload optionChangedPayload
type OptionSubTasksChangedPayload optionChangedPayload
type OptionAttachmentsChangedPayload optionChangedPayload
type OptionWatchingChangedPayload optionChangedPayload
type OptionVotingChangedPayload optionChangedPayload
type OptionUnassignedIssuesChangedPayload optionChangedPayload
