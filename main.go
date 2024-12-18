package main

import "fmt"

//task struct
type Task struct {
	ID    int
	Title string
	Done  bool
}

//task-manager struct
type TaskManager struct {
	tasks  []Task
	nextId int
}

//add task
func (tm *TaskManager) AddTask(title string) {
	task := Task{
		ID:    tm.nextId,
		Title: title,
		Done:  false,
	}
	tm.tasks = append(tm.tasks, task)
	tm.nextId++
	fmt.Printf("Task Added: %s (ID: %d)\n", title, task.ID)
}

//view task
func (tm *TaskManager) ViewTasks(){
	fmt.Println("All the tasks are:")
	for _, task := range tm.tasks{
		status := "Pending"
		if task.Done{
			status = "Completed"
		}
		fmt.Printf("-[%d] %s (%v)\n",task.ID, task.Title, status )
	}
	fmt.Println()
}
//deleting task
func (tm *TaskManager) DeleteTask(id int){
	for i, task := range tm.tasks{
		if task.ID == id{
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			fmt.Printf("Task with Id %d is deleted.\n", id)
			return
		}
	}
	fmt.Printf("task id %d is not found !", id)
}

func main() {
	manager := &TaskManager{}

	manager.AddTask("learn go basic")
	manager.AddTask("build a simple project")
	manager.AddTask("post what u learnt")

	manager.ViewTasks()

	manager.DeleteTask(1)

	manager.ViewTasks()

}
