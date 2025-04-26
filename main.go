package main

import (
	"bufio"
	"fmt"
	"main/model"
	"os"
	"strconv"
	"strings"
)

const COMMAND_ADD = "add"
const COMMAND_LIST = "list"
const COMMAND_DELETE = "delete"
const COMMAND_COMPLETE = "complete"

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
		if words[0] != "task" {
			fmt.Println("Command error, program exited.")
			return
		} 
		command := words[1]
		switch command {
			case COMMAND_ADD:
				add(&ts, words[2])
			case COMMAND_LIST:
				res := getList(&ts)
				fmt.Println("List Tasks:\n", *res)
			case COMMAND_DELETE:
				delete(&ts, words[2])
			case COMMAND_COMPLETE:
				complete(&ts, words[2])
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
