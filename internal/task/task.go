package task

import (
	"time"
)

type Status int
type Priority int

const (
	Todo Status = iota + 1
	InProgress
	Done
)

const (
	Low Priority = iota + 1
	Medium
	High
)

type Task struct {
	ID          string
	Title       string
	Description string
	Status      Status
	Priority    Priority
	DueDate     time.Time
	CreatedAt   time.Time
}

func (s Status) String() string {
	return []string{"Todo", "InProgress", "Done"}[s-1]
}

func (p Priority) String() string {
	return []string{"Low", "Medium", "High"}[p-1]
}
