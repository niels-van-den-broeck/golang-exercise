package store

import (
	"crypto/rand"
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

func (t *TaskStore) Add(task task.Task) task.Task {
	task.ID = rand.Text()
	task.CreatedAt = time.Now()

	t.tasks = append(t.tasks, task)

	return task
}

func (t *TaskStore) Get(id string) (task.Task, bool) {
	for _, task := range t.tasks {
		if task.ID == id {
			return task, true
		}
	}
	return task.Task{}, false
}

func (t TaskStore) List() []task.Task {
	return t.tasks
}
