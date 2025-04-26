package model

import (
	"fmt"
)

type TaskStorage struct {
	Tasks []Task
	IdCounter int
}

type Task struct {
	Id				int
	Name 			string
}

func (ts *TaskStorage) Init() {
	ts.Tasks = []Task{}
	ts.IdCounter = 0
}

func (ts *TaskStorage) AddTask(task *Task) (*Task, error) {
	ts.Tasks = append(ts.Tasks, *task)
	ts.IdCounter += 1
	fmt.Println("Successfully added new task.", ts.Tasks)
	return task, nil
} 

func (ts *TaskStorage) ListTasks() (*[]Task, error) {
	return &ts.Tasks, nil
}

func (ts *TaskStorage) DeleteTaskById(id int) error {
	deleteIdx := -1
	for i,v := range ts.Tasks {
		if (v.Id == id) {
			deleteIdx = i
			ts.Tasks = append(ts.Tasks[:deleteIdx], ts.Tasks[deleteIdx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task with id %d not found", id)
}