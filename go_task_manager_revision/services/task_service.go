package services

import "github.com/eswar/rev_project_task_manager/models"

type TaskService struct {
	tasks  []models.Task
	nextID uint32
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

func (s *TaskService) CreateTask(
	title, description, priority string,
) models.Task {

	task := models.Task{
		ID:          s.nextID,
		Title:       title,
		Description: description,
		Status:      models.StatusPending,
		Priority:    priority,
	}

	s.tasks = append(s.tasks, task)
	s.nextID++

	return task
}

func (s *TaskService) GetAllTasks() []models.Task {
	return append([]models.Task(nil), s.tasks...)
}

func (s *TaskService) GetTaskByID(id uint32) (models.Task, bool) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, true
		}
	}

	return models.Task{}, false
}

func (s *TaskService) UpdateTaskStatus(
	id uint32, status string,
) bool {

	if status != models.StatusPending &&
		status != models.StatusInProgress &&
		status != models.StatusCompleted {
		return false
	}

	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Status = status
			return true
		}
	}

	return false
}

func (s *TaskService) DeleteTask(id uint32) bool {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}

	return false
}
