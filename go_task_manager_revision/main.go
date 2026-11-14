package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eswar/rev_project_task_manager/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	taskService := services.NewTaskService()

	for {

		fmt.Println("your task list :", taskService)
		fmt.Println("\n===== TASK MANAGER =====")
		fmt.Println("1. Create Task")
		fmt.Println("2. Get All Tasks")
		fmt.Println("3. Get Task By ID")
		fmt.Println("4. Update Task Status")
		fmt.Println("5. Delete Task")
		fmt.Println("6. Exit")

		fmt.Print("Choose option: ")
		scanner.Scan()

		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Description: ")
			scanner.Scan()
			description := scanner.Text()

			fmt.Print("Priority: ")
			scanner.Scan()
			priority := scanner.Text()

			task := taskService.CreateTask(
				title, description, priority,
			)

			fmt.Println("Created:", task)

		case 2:
			tasks := taskService.GetAllTasks()

			for _, task := range tasks {
				fmt.Printf("%+v\n", task)
			}

			// case 3:
			// 	fmt.Print("Enter task ID: ")
			// 	scanner.Scan()

			// 	id, err := strconv.ParseUint(scanner.Text(), 10, 32)
			// 	if err != nil {
			// 		fmt.Println("Invalid ID")
			// 		continue
			// 	}

			// 	task, found := taskService.GetTaskByID(uint32(id))
			// 	if found != nil{
			// 		fmt.Println("Task not found")
			// 		continue
			// 	}

			// 	fmt.Printf("%+v\n", task)
		case 3:
			fmt.Print("Enter task ID: ")
			scanner.Scan()

			id, err := strconv.ParseUint(scanner.Text(), 10, 32)
			if err != nil {
				fmt.Println("Invalid ID:", err)
				continue
			}

			task, err := taskService.GetTaskByID(uint32(id))
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("%+v\n", task)

		case 4:
			fmt.Print("Enter task ID: ")
			scanner.Scan()
			id, err := strconv.ParseUint(scanner.Text(), 10, 32)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			fmt.Print("Enter status: ")
			scanner.Scan()
			status := strings.TrimSpace(scanner.Text())

			if taskService.UpdateTaskStatus(uint32(id), status) {
				fmt.Println("Status updated")
			} else {
				fmt.Println("Invalid status or task ID")
			}

		case 5:
			fmt.Print("Enter task ID: ")
			scanner.Scan()

			id, err := strconv.ParseUint(scanner.Text(), 10, 32)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			if taskService.DeleteTask(uint32(id)) {
				fmt.Println("Task deleted")
			} else {
				fmt.Println("Task not found")
			}

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
