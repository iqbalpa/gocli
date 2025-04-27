package service

import (
	"fmt"
	"main/model"
	"strconv"
)

type Service struct {}

func (service *Service) Add(ts *model.TaskStorage, s string) {
	// create new task
	t := model.Task{
		Name: s, 
		Id: ts.IdCounter, 
		Completed: false,
	}
	// add task
	ts.AddTask(&t)
}

func (service *Service) GetList(ts *model.TaskStorage) (*[]model.Task) {
	res, _ := ts.ListTasks()
	return res
}

func (service *Service) Delete(ts *model.TaskStorage, id string) {
	idInt, _ := strconv.Atoi(id)
	err := ts.DeleteTaskById(idInt)
	if err != nil {
		fmt.Println(err)
	}
}

func (service *Service) Complete(ts *model.TaskStorage, id string) {
	idInt, _ := strconv.Atoi(id)
	err := ts.CompleteTaskById(idInt)
	if err != nil {
		fmt.Println(err)
	}
}
