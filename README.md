# Task Management System

A simple **Task Management System** built in Go. This program helps you manage tasks with basic operations like adding, viewing, and deleting tasks. It is beginner-friendly and focuses on clean and straightforward implementation.

---

## Features

1. **Add Tasks**: Add tasks with a title to your task list.
2. **View Tasks**: Display all tasks with their ID, title, and status (Pending/Completed).
3. **Delete Tasks**: Remove a task from the list using its unique ID.

---

## Code Overview

### Task Struct
The `Task` struct represents a single task with the following fields:
- `ID` (int): A unique identifier for the task.
- `Title` (string): The description of the task.
- `Done` (bool): Indicates whether the task is completed.

### TaskManager Struct
The `TaskManager` struct manages a list of tasks and contains:
- `tasks` ([]Task): A slice to store tasks.
- `nextID` (int): A counter to generate unique IDs for new tasks.

### Methods

#### `AddTask`
Adds a new task to the task manager.
```go
func (tm *TaskManager) AddTask(title string)
```

#### `ViewTasks`
Displays all tasks with their details.
```go
func (tm *TaskManager) ViewTasks()
```

#### `DeleteTask`
Removes a task by its unique ID.
```go
func (tm *TaskManager) DeleteTask(id int)
```

---

## How to Run

1. **Install Go**: Ensure Go is installed on your system. Download it from [Go's official website](https://golang.org/dl/).

2. **Clone the Code**: Copy the code into a `.go` file, for example, `task_manager.go`.

3. **Run the Program**:
```bash
go run task_manager.go
```

---

## Example Output

```
Task added: Learn Go basics (ID: 0)
Task added: Build a simple project (ID: 1)

Task List:
- [0] Learn Go basics (Pending)
- [1] Build a simple project (Pending)

Task with ID 1 deleted.

Task List:
- [0] Learn Go basics (Pending)
```

---

## Why This Project?

- **Beginner-Friendly**: Designed for those new to Go.
- **Practical**: Demonstrates core Go concepts like structs, slices, and methods.
- **Customizable**: Easy to extend with additional features like marking tasks as completed or saving data to a file.

---

Happy Coding! 🚀

