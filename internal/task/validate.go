package task

import (
	"fmt"
	"strings"
	"time"
)

// ValidationError reports which field failed validation and why.
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Msg)
}

func (t Task) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return &ValidationError{Field: "title", Msg: "must not be empty"}
	}
	if t.DueDate.Before(time.Now()) {
		return &ValidationError{Field: "due_date", Msg: "must not be in the past"}
	}
	return nil
}
