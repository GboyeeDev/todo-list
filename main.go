package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tm := TaskOwner{}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("#### Todo List ####")
		fmt.Println("1. Add task")
		fmt.Println("2. Mark Task as completed")
		fmt.Println("3. Remove task")
		fmt.Println("4. Display Tasks")
		fmt.Println("5. Exit")
		fmt.Println("Pick your choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task name: ")
			taskName, _ := reader.ReadString('\n')
			taskName = strings.TrimSpace(taskName)
			tm.AddTask(taskName)
		case "2":
			fmt.Print("Enter task ID to mark as completed: ")
			taskIDInput, _ := reader.ReadString('\n')
			taskIDInput = strings.TrimSpace(taskIDInput)
			taskID, _ := strconv.Atoi(taskIDInput)
			tm.MarkTaskComplete(taskID)
		case "3":
			fmt.Print("Enter task ID to remove: ")
			taskIDInput, _ := reader.ReadString('\n')
			taskIDInput = strings.TrimSpace(taskIDInput)
			taskID, _ := strconv.Atoi(taskIDInput)
			tm.RemoveTask(taskID)
		case "4":
			tm.DisplayTasks()
		case "5":
			fmt.Println("Closing... ")
			os.Exit(0)
		default:
			fmt.Println("Wrong chioce. Please choose again.")
		}
	}
}
