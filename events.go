package jira_hooks

type Event jsonString

const (
	// Transitions
	StatusTransitionEvent = Event("status_transition")

	// Issues
	IssueCreatedEvent = Event("jira:issue_created")
	IssueDeletedEvent = Event("jira:issue_deleted")
	IssueUpdatedEvent = Event("jira:issue_updated")
	IssueWorkLogEvent = Event("jira:worklog_updated")

	// Work logs
	WorkLogCreatedEvent = Event("worklog_created")
	WorkLogUpdatedEvent = Event("worklog_updated")
	WorkLogDeletedEvent = Event("worklog_deleted")

	// Comments
	CommentCreatedEvent = Event("comment_created")
	CommentUpdatedEvent = Event("comment_updated")
	CommentDeletedEvent = Event("comment_deleted")

	// Links
	LinkCreatedEvent = Event("issuelink_created")
	LinkDeletedEvent = Event("issuelink_deleted")

	// Users
	UserCreatedEvent = Event("user_created")
	UserUpdatedEvent = Event("user_updated")
	UserDeletedEvent = Event("user_deleted")

	// Projects
	ProjectCreatedEvent  = Event("project_created")
	ProjectUpdatedEvent  = Event("project_updated")
	ProjectDeletedEvent  = Event("project_deleted")
	ProjectArchivedEvent = Event("project_archived")
	ProjectRestoredEvent = Event("project_restored")

	// Boards
	BoardCreatedEvent              = Event("board_created")
	BoardUpdatedEvent              = Event("board_updated")
	BoardDeletedEvent              = Event("board_deleted")
	BoardConfigurationChangedEvent = Event("board_configuration_changed")

	// Sprints
	SprintCreatedEvent = Event("sprint_created")
	SprintUpdatedEvent = Event("sprint_updated")
	SprintDeletedEvent = Event("sprint_deleted")
	SprintStartedEvent = Event("sprint_started")
	SprintClosedEvent  = Event("sprint_closed")

	// Versions
	VersionCreatedEvent    = Event("jira:version_created")
	VersionUpdatedEvent    = Event("jira:version_updated")
	VersionDeletedEvent    = Event("jira:version_deleted")
	VersionReleasedEvent   = Event("jira:version_released")
	VersionUnreleasedEvent = Event("jira:version_unreleased")

	// Options
	OptionVotingChangedEvent           = Event("option_voting_changed")
	OptionWatchingChangedEvent         = Event("option_watching_changed")
	OptionSubTasksChangedEvent         = Event("option_subtasks_changed")
	OptionIssueLinksChangedEvent       = Event("option_issuelinks_changed")
	OptionAttachmentsChangedEvent      = Event("option_attachments_changed")
	OptionTimeTrackingChangedEvent     = Event("option_timetracking_changed")
	OptionUnassignedIssuesChangedEvent = Event("option_unassigned_issues_changed")
)
