package gh

import "github.com/rsteube/carapace"

// ActionWorkflowEvents completes events triggering a workflow
//
//	create
//	delete
func ActionWorkflowEvents() carapace.Action {
	return carapace.ActionValues(
		"branch_protection_rule",
		"check_run",
		"check_suite",
		"create",
		"delete",
		"deployment",
		"deployment_status",
		"discussion",
		"discussion_comment",
		"fork",
		"gollum",
		"issue_comment",
		"issues",
		"label",
		"merge_group",
		"milestone",
		"page_build",
		"project",
		"project_card",
		"project_column",
		"public",
		"pull_request",
		"pull_request_review",
		"pull_request_review_comment",
		"pull_request_target",
		"push",
		"registry_package",
		"release",
		"repository_dispatch",
		"schedule",
		"status",
		"watch",
		"workflow_call",
		"workflow_dispatch",
		"workflow_run",
	)
}

// ActionWorkflowFields completes workflow fields.
func ActionWorkflowFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"name",
		"path",
		"state",
	)
}
