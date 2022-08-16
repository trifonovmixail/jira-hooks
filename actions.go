package jira_hooks

type action jsonString

const (
	// Issue update actions
	issueGenericAction = action("issue_generic")
	issueUpdatedAction = action("issue_updated")

	issueAssignedAction = action("issue_assigned")

	issueCommentCreatedAction = action("issue_commented")
	issueCommentUpdatedAction = action("issue_comment_edited")
	issueCommentDeletedAction = action("issue_comment_deleted")

	issueWorkLogCreatedAction = action("issue_work_logged")
	issueWorkLogUpdatedAction = action("issue_worklog_updated")
	issueWorkLogDeletedAction = action("issue_worklog_deleted")

	issueMovedAction = action("issue_moved")

	issueClosedAction = action("issue_closed")

	issueResolvedAction = action("issue_resolved")
)
