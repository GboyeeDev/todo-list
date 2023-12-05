package main

import (
	"fmt"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

type TaskOwner struct {
	tasks []Task
}

// Adds task
func (tm *TaskOwner) AddTask(name string) {
	taskID := len(tm.tasks) + 1
	newTask := Task{ID: taskID, Name: name, Completed: false}
	tm.tasks = append(tm.tasks, newTask)
}

// Displays tasks
func (tm *TaskOwner) DisplayTasks() {
	fmt.Println("Tasks:")
	for _, task := range tm.tasks {
		completed := "Not Completed"
		if task.Completed {
			completed = "Completed"
		}
		fmt.Printf("%d: %s - %s \n", task.ID, task.Name, completed)
	}
}

// Mark tasks as complete
func (tm *TaskOwner) MarkTaskComplete(taskID int) {
	for index := range tm.tasks {
		if tm.tasks[index].ID == taskID {
			tm.tasks[index].Completed = true
			break
		}
	}
}

// Removes Task
func (tm *TaskOwner) RemoveTask(taskID int) {
	for index, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
			break
		}
	}
}
