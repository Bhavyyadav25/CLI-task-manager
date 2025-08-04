package store

import (
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gookit/slog"

	"github.com/Bhavyyadav25/CLI-task-manager/internal/domain"
)

type FileRepo struct {
	filePath string
}

func NewFileRepo(filePath string) *FileRepo {
	return &FileRepo{
		filePath: filePath,
	}
}

func (f *FileRepo) Create(task *domain.Task) domain.Error {
	tasks, err := f.Load()
	if err.Code != "" {
		return err
	}

	for _, t := range tasks {
		if t.ID == task.ID || t.Description == task.Description {
			return domain.Error{Message: "Task already exists"}
		}
	}

	tasks = append(tasks, *task)
	return f.save(tasks)
}

func (f *FileRepo) List() ([]domain.Task, domain.Error) {
	return f.Load()
}

func (f *FileRepo) Update(task *domain.Task) (domain.Task, domain.Error) {
	tasks, err := f.Load()
	if err.Message != "" {
		return domain.Task{}, err
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == task.ID {
			if tasks[i].Done {
				return domain.Task{}, domain.Error{Message: "task already done"}
			}
			tasks[i].Done = true
			tasks[i].UpdateAt = time.Now()
			task = &tasks[i]
			updated = true
			break
		}
	}

	if !updated {
		return domain.Task{}, domain.Error{Message: "task not found"}
	}

	err = f.save(tasks)
	if err.Message != "" {
		return domain.Task{}, err
	}

	return *task, domain.Error{}
}

func (f *FileRepo) Delete(id uuid.UUID) domain.Error {
	tasks, err := f.Load()
	if err.Message != "" {
		return err
	}

	filtered := make([]domain.Task, 0, len(tasks))
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		filtered = append(filtered, t)
	}
	if !found {
		return domain.Error{Message: "task not found"}
	}

	return f.save(filtered)
}

func (f *FileRepo) Load() ([]domain.Task, domain.Error) {
	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		slog.Warn("file not found, creating a new one")
		if _, err := os.Create(f.filePath); err != nil {
			slog.Error(err)
			return []domain.Task{}, domain.Error{Message: err.Error()}
		}
	}

	taskByte, err := os.ReadFile(f.filePath)
	if err != nil {
		slog.Error(err)
		return []domain.Task{}, domain.Error{Message: err.Error()}
	}

	var tasks []domain.Task
	if len(taskByte) == 0 {
		return tasks, domain.Error{}
	}

	if err := json.Unmarshal(taskByte, &tasks); err != nil {
		slog.Error(err)
		return []domain.Task{}, domain.Error{Message: err.Error()}
	}

	return tasks, domain.Error{}
}

func (f *FileRepo) save(tasks []domain.Task) domain.Error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		slog.Error(err)
		return domain.Error{Message: err.Error()}
	}

	if err := os.WriteFile(f.filePath, data, 0644); err != nil {
		slog.Error(err)
		return domain.Error{Message: err.Error()}
	}

	return domain.Error{}
}
