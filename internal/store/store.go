package store

import (
	"nvdb/taskapp/internal/task"
)

type Store interface {
	Add(tsk task.Task) (task.Task, error)
	Get(id string) (task.Task, error)
	List() []task.Task
	Update(t task.Task) (task.Task, error)
	Delete(id string) error
}
