package models

type Task struct {
	Description string     `json="description"`
	Status      TaskStatus `json="status"`
}

type TaskStatus int

const (
	WAITING TaskStatus = iota + 1
	DOING
	DONE
)

func NewTask(description string, status TaskStatus) *Task {
	return &Task{
		description,
		status,
	}
}
