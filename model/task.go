package model

import "fmt"

type TaskStorage struct {
	Tasks []Task
}

type Task struct {
	Name 			string
}

func (ts *TaskStorage) Init() {
	ts.Tasks = []Task{}
}

func (ts *TaskStorage) AddTask(task *Task) (*Task, error) {
	ts.Tasks = append(ts.Tasks, *task)
	fmt.Println("Successfully added new task.", ts.Tasks)
	return task, nil
} 

func (ts *TaskStorage) ListTasks() (*[]Task, error) {
	return &ts.Tasks, nil
}
