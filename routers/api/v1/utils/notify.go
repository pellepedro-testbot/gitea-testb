// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package utils

import (
	"strings"

	activities_model "gitea.dev/models/activities"
)

// SubjectTypeToSource converts a list of subject-type strings to notification sources.
func SubjectTypeToSource(value []string) (result []activities_model.NotificationSource) {
	for _, v := range value {
		switch strings.ToLower(v) {
		case "issue":
			result = append(result, activities_model.NotificationSourceIssue)
		case "pull":
			result = append(result, activities_model.NotificationSourcePullRequest)
		case "commit":
			result = append(result, activities_model.NotificationSourceCommit)
		case "repository":
			result = append(result, activities_model.NotificationSourceRepository)
		}
	}
	return result
}
