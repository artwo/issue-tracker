package model

import "issue-tracker/utils"

type Status string

const (
	StatusNew        Status = "new"
	StatusRefinement Status = "refinement"
	StatusReady      Status = "ready"
	StatusTodo       Status = "to-do"
	StatusInProgress Status = "in-progress"
	StatusInReview   Status = "in-review"
	StatusTesting    Status = "testing"
	StatusDone       Status = "done"
	StatusOnHold     Status = "on-hold"
	StatusBlocked    Status = "blocked"
	StatusCancelled  Status = "cancelled"
	StatusRejected   Status = "rejected"
)

var statusNames = map[Status]string{
	StatusNew:        "New",
	StatusRefinement: "Refinement",
	StatusReady:      "Ready",
	StatusTodo:       "To Do",
	StatusInProgress: "In Progress",
	StatusInReview:   "Ready for Review",
	StatusTesting:    "Ready for Testing",
	StatusDone:       "Done",
	StatusOnHold:     "On Hold",
	StatusBlocked:    "Blocked",
	StatusCancelled:  "Cancelled",
	StatusRejected:   "Rejected",
}

var statusNamesKeys []Status

func init() {
	statusNamesKeys = make([]Status, 0, len(statusNames))
	for k := range statusNames {
		statusNamesKeys = append(statusNamesKeys, k)
	}
}

func (status Status) IsValid() bool {
	_, ok := statusNames[status]
	return ok
}

func (status Status) String() string {
	if value, ok := statusNames[status]; ok {
		return value
	}
	return "Unknown"
}

func AllStatusNamesAsString() string {
	return utils.ToString(statusNamesKeys)
}
