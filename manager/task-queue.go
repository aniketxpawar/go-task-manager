package manager

import (
	"sync"
	"time"
)

type Task struct {
	ID          int
	Payload     string
	SubmittedAt time.Time
}

type TaskQueue struct{
	tasks []Task
	mu sync.Mutex
}

func NewTaskQueue() *TaskQueue{
	return &TaskQueue{
		tasks: make([]Task, 0),
	}
}

func (t *TaskQueue) Enqueue(task Task){
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tasks = append(t.tasks, task)
}

func (t *TaskQueue) Dequeue() *Task{
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.tasks) == 0{
		return nil
	}
	task := t.tasks[0]
	t.tasks = t.tasks[1:]
	return &task
}