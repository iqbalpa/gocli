package main

import (
	"bufio"
	"fmt"
	"main/model"
	"os"
	"strconv"
	"strings"
)

const (
	CommandAdd = "add"
	CommandList = "list"
	CommandDelete = "delete"
	CommandComplete = "complete"
)

func main() {
	fmt.Println("Program started :)")
	reader := bufio.NewReader(os.Stdin)

	// create task storage model
	ts := model.TaskStorage{}
	ts.Init()

	// add LOOP as label for this infinite loop
	LOOP: for {
    line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		
		// first word must be "task"
		if (words[0] != "task" || len(words) < 2) {
			fmt.Println("Command error, program exited.")
			return
		} 

		command := words[1]
		arg := words[2]
		switch command {
			case CommandAdd:
				add(&ts, arg)
			case CommandList:
				res := getList(&ts)
				fmt.Println("List Tasks:\n", *res)
			case CommandDelete:
				delete(&ts, arg)
			case CommandComplete:
				complete(&ts, arg)
			default:
				fmt.Println("Program stopped.")
				break LOOP
		}
	}
}

func add(ts *model.TaskStorage, s string) {
	// create new task
	t := model.Task{
		Name: s, 
		Id: ts.IdCounter, 
		Completed: false,
	}
	// add task
	ts.AddTask(&t)
}

func getList(ts *model.TaskStorage) (*[]model.Task) {
	res, _ := ts.ListTasks()
	return res
}

func delete(ts *model.TaskStorage, id string) {
	idInt, _ := strconv.Atoi(id)
	err := ts.DeleteTaskById(idInt)
	if err != nil {
		fmt.Println(err)
	}
}

func complete(ts *model.TaskStorage, id string) {
	idInt, _ := strconv.Atoi(id)
	err := ts.CompleteTaskById(idInt)
	if err != nil {
		fmt.Println(err)
	}
}
