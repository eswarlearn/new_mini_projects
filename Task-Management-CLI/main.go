package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int
	Name   string
	Status bool
}

var tasks []Task

func main() {

	if len(os.Args) < 1 {
		fmt.Println("enter the task")
	}
	allData, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("Error in json data extractiom", err)
	}

	err = json.Unmarshal(allData, &tasks)
	if err != nil {
		fmt.Println("Error in Unmarshal", err)
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 2 {
			fmt.Println("enter task name")
			return
		}
		taskName := strings.Join(os.Args[2:], " ")
		addTask(taskName)
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 2 {
			fmt.Println("Please enter valif id ")
		}

		taskId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Enter valid number :", err)
			return
		}
		markDone(taskId)

	case "delete":
		if len(os.Args) < 2 {
			fmt.Println("Provide ID fordeleting")
		}
		taskId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Enter valid number :", err)
			return
		}
		deleteTask(taskId)

	}
}

func addTask(name string) {
	fmt.Println(name)
	tskId := 0
	if len(tasks) > 0 {
		tskId = tasks[len(tasks)-1].ID
	}
	tskId++
	tasks = append(tasks, Task{ID: tskId, Name: name})
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error struct to json :", err)
	}
	err = os.WriteFile("task.json", data, 0644)
	if err != nil {
		fmt.Println("Error json file creation :", err)
	}
}

func listTasks() {
	for _, t := range tasks {
		status := " "
		if t.Status {
			status = "X"
		}
		fmt.Printf("[%s] %s\n", status, t.Name)
	}
}
func markDone(id int) {
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = true
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found")
		return
	}
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("data error in mark donr :", err)
	}
	os.WriteFile("task.json", data, 0644)
}
func deleteTask(id int) {
	var tempTask []Task
	for i := range tasks {
		if tasks[i].ID == id {
			continue
		}
		tempTask = append(tempTask, tasks[i])
	}
	data, err := json.Marshal(tempTask)
	if err != nil {
		fmt.Printf("no data found for deleting : %v/n", err)
	}
	tasks = tempTask
	os.WriteFile("task.json", data, 0644)//1 ■ Build CLI project
}

/*🔧 Step 1: Create helper functions
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("task.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("task.json", data, 0644)
}

👉 Now your code is cleaner everywhere else.

🔧 Step 2: Remove global state

Instead of:

var tasks []Task

Do:

tasks, err := loadTasks()
🔧 Step 3: Improve delete (clean + efficient)
func deleteTask(tasks []Task, id int) []Task {
	var result []Task
	for _, t := range tasks {
		if t.ID != id {
			result = append(result, t)
		}
	}
	return result
}
	*****Next Step (Don’t Stay Here)

Now I’d push you slightly harder:

👉 Upgrade this project:

Add:

✅ update command (edit task)
✅ pending → show only incomplete tasks
✅ clear → delete all completed tasks
✅ Sort tasks by ID
💣 Next-Level Challenge (Very Important)

Convert this CLI into:

👉 REST API using Gin

Same logic:

POST /task
GET /tasks
PUT /task/:id
DELETE /task/:id
	*
*/
