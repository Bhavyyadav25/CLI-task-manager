package usecase

import (
	"errors"
	"time"

	"github.com/Bhavyyadav25/CLI-task-manager/internal/domain"
	"github.com/Bhavyyadav25/CLI-task-manager/internal/repo"
	"github.com/google/uuid"
)

type TaskService struct {
	repo repo.TaskRepository
}

func NewTaskService(r repo.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) AddTask(desc string) (domain.Task, error) {
	var task domain.Task
	task.ID = uuid.New()
	task.Description = desc
	task.Done = false
	task.CreatedAt = time.Now()
	task.UpdateAt = task.CreatedAt

	err := s.repo.Create(&task)
	if err.Message != "" {
		return domain.Task{}, errors.New(err.Message)
	}

	return task, nil
}

func (s *TaskService) ListTasks() ([]domain.Task, error) {

	tasks, err := s.repo.List()
	if err.Message != "" {
		return []domain.Task{}, errors.New(err.Message)
	}

	return tasks, nil
}

func (s *TaskService) MarkDone(id string) (domain.Task, error) {
	task, err := s.repo.Update(&domain.Task{ID: uuid.MustParse(id), Done: true})
	if err.Message != "" {
		return domain.Task{}, errors.New(err.Message)
	}

	return task, nil
}
