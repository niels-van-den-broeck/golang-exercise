package store

import (
	"crypto/rand"
	"nvdb/taskapp/internal/core"
	"nvdb/taskapp/internal/task"
	"time"
)

type MemoryStore struct {
	tasks []task.Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: []task.Task{},
	}
}

func (t *MemoryStore) Add(tsk task.Task) (task.Task, error) {
	tsk.ID = rand.Text()
	tsk.CreatedAt = time.Now()

	err := tsk.Validate()

	if err != nil {
		return task.Task{}, err
	}

	t.tasks = append(t.tasks, tsk)

	return tsk, nil
}

func (t *MemoryStore) Get(id string) (task.Task, error) {
	for _, task := range t.tasks {
		if task.ID == id {
			return task, core.ErrNotFound
		}
	}
	return task.Task{}, nil
}

func (t MemoryStore) List() []task.Task {
	return t.tasks
}
