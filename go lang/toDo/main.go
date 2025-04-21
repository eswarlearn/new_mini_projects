package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task

const fileName = "tasks.json"

func loadTasks() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		json.Unmarshal(data, &tasks)
	}
}
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "")
	_ = os.WriteFile(fileName, data, 0644)
}

func addTask(title string) {
	tasks = append(tasks, Task{Title: title, Done: false})
	saveTasks()
	fmt.Println("Task added... !")
}

func listTasks() {
	for i, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Title)

	}
}

func markDone(index int) {
	if index >= 0 && index < len(tasks) {
		tasks[index].Done = true
		saveTasks()
		fmt.Println("Task marked as done!")
	} else {
		fmt.Println("Invalid task number.")
	}
}

func deleteTask(index int) {
	if index >= 0 && index < len(tasks) {
		tasks = append(tasks[:index], tasks[index+1:]...)
		saveTasks()
		fmt.Println("Task deleted!")
	} else {
		fmt.Println("Invalid task number.")
	}
}

func main() {
	loadTasks()
	if len(os.Args) < 2 {
		fmt.Println("Usage: add/list/done/delete")
		return
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task title.")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task number.")
			return
		}
		var index int
		fmt.Sscanf(os.Args[2], "%d", &index)
		markDone(index - 1)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task number.")
			return
		}
		var index int
		fmt.Sscanf(os.Args[2], "%d", &index)
		deleteTask(index - 1)
	default:
		fmt.Println("Unknown command.")
	}
}
