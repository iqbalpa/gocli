package main

import (
	"bufio"
	"fmt"
	"main/model"
	"main/service"
	"os"
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
	service := service.Service{}

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
		switch command {
			case CommandAdd:
				service.Add(&ts, words[2])
			case CommandList:
				res := service.GetList(&ts)
				fmt.Println("List Tasks:\n", *res)
			case CommandDelete:
				service.Delete(&ts, words[2])
			case CommandComplete:
				service.Complete(&ts, words[2])
			default:
				fmt.Println("Program stopped.")
				break LOOP
		}
	}
}
