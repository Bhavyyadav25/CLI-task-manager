package repo

import (
	"github.com/Bhavyyadav25/CLI-task-manager/internal/domain"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *domain.Task) domain.Error
	List() ([]domain.Task, domain.Error)
	Update(task *domain.Task) domain.Error
	Delete(id uuid.UUID) domain.Error
}
