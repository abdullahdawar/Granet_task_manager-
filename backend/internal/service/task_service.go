package service

import (
	"errors"
	"strings"

	"granet-task-manager/backend/internal/model"
	"granet-task-manager/backend/internal/repository"
)

type TaskService struct {
	repository *repository.TaskRepository
}

func NewTaskService(repository *repository.TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) CreateTask(task *model.Task) error {
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("title is required")
	}

	if strings.TrimSpace(task.DueDate) == "" {
		return errors.New("due date is required")
	}

	if strings.TrimSpace(task.Status) == "" {
		task.Status = "pending"
	}

	return s.repository.Create(task)
}

func (s *TaskService) GetTasks() ([]model.Task, error) {
	return s.repository.GetAll()
}

func (s *TaskService) UpdateTask(task *model.Task) error {
	if task.ID <= 0 {
		return errors.New("invalid task ID")
	}

	if strings.TrimSpace(task.Title) == "" {
		return errors.New("title is required")
	}

	if strings.TrimSpace(task.DueDate) == "" {
		return errors.New("due date is required")
	}

	if strings.TrimSpace(task.Status) == "" {
		return errors.New("status is required")
	}

	return s.repository.Update(task)
}

func (s *TaskService) DeleteTask(id int64) error {
	if id <= 0 {
		return errors.New("invalid task ID")
	}

	return s.repository.Delete(id)
}
