package model

import "issue-tracker/utils"

type Status string

const (
	New        string = "new"
	Refinement string = "refinement"
	Ready      string = "ready"
	ToDo       string = "to-do"
	InProgress string = "in-progress"
	InReview   string = "in-review"
	Testing    string = "testing"
	Done       string = "done"
	OnHold     string = "on-hold"
	Blocked    string = "blocked"
	Cancelled  string = "cancelled"
	Rejected   string = "rejected"
)

var statusNames = map[Status]string{
	"new":         "New",
	"refinement":  "Refinement",
	"ready":       "Ready",
	"to-do":       "To Do",
	"in-progress": "In Progress",
	"in-review":   "Ready for Review",
	"testing":     "Ready for Testing",
	"done":        "Done",
	"on-hold":     "On Hold",
	"blocked":     "Blocked",
	"cancelled":   "Cancelled",
	"rejected":    "Rejected",
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
