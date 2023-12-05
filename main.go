package main

func main() {
	tm := TaskOwner{}

	// Adding tasks
	tm.AddTaskName("Task 1")
	tm.AddTaskName("Task 2")
	tm.AddTaskName("Task 3")

	// Displaying tasks
	tm.DisplayTasks()

	// Marking a task as completed
	tm.MarkTaskComplete(2)

	// Displaying tasks when one is completed
	tm.DisplayTasks()

	// Removing a task
	tm.RemoveTask(1)

	// Displaying tasks after deleting
	tm.DisplayTasks()

}
