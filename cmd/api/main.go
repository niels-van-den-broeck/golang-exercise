package main

import (
	"errors"
	"fmt"
	"nvdb/taskapp/internal/store"
	"nvdb/taskapp/internal/task"
)

func main() {
	fmt.Println("task app starting")
	taskStore := store.MemoryStore{}

	myTask, err := taskStore.Add(task.Task{})

	var validationErr *task.ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("A validation error occurred (%s)", err)
		return
	}

	fmt.Printf("hello %s", myTask.Description)
}
