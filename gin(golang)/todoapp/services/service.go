package services

import (
	"codesignal.com/todoapp/models"
)

var todos = []models.Todo{
	{ID: 1, Title: "Learn Go", Completed: false},
	{ID: 2, Title: "Master Gin", Completed: true},
	{ID: 3, Title: "Create Gin App", Completed: false},
}

func FilterTodo(completed *bool) []models.Todo { //why []
	if completed == nil {
		return todos
	}

	var filtered []models.Todo //why []
	for _, todo := range todos {
		if todo.Completed == *completed {
			filtered = append(filtered, todo)
		}
	}

	return filtered

}

func FilterTodoById(id int) (models.Todo, bool) { //why no []
	for _, todo := range todos {
		if todo.ID == id {
			return todo, true //what is this true?
		}
	}

	return models.Todo{}, false // why {}
}

func ResetTodos() {
	todos = []models.Todo{}
}

func AddTodo(newTodo models.Todo) (models.Todo, error) {
	newTodo.ID = len(todos) + 1
	newTodo.Completed = false
	todos = append(todos, newTodo)
	return newTodo, nil
}
