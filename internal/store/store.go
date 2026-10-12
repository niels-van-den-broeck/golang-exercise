package store

import (
	"crypto/rand"
	application_error "nvdb/taskapp/internal/core"
	"nvdb/taskapp/internal/task"
	"time"
)

type TaskStore struct {
	tasks []task.Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: []task.Task{},
	}
}

func (t *TaskStore) Add(tsk task.Task) (task.Task, error) {
	tsk.ID = rand.Text()
	tsk.CreatedAt = time.Now()

	err := tsk.Validate()

	if err != nil {
		return task.Task{}, err
	}

	t.tasks = append(t.tasks, tsk)

	return tsk, nil
}

func (t *TaskStore) Get(id string) (task.Task, error) {
	for _, task := range t.tasks {
		if task.ID == id {
			return task, application_error.ErrNotFound
		}
	}
	return task.Task{}, nil
}

func (t TaskStore) List() []task.Task {
	return t.tasks
}
