package store

import (
	"encoding/json"
	"os"

	"github.com/gookit/slog"

	"github.com/Bhavyyadav25/CLI-task-manager/internal/domain"
)

const (
	Create = iota
	List
	Update
	Delete

	filePath = "user_data.json"
)

func Storage(action int, task *domain.Task) {
	if !jsonFileExists(filePath) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			slog.Error(err)
		}
	}

	switch action {
	case Create:
		// task.ID = uuid.New()
		// task.CreatedAt = time.Now()
		// task.UpdateAt = task.CreatedAt

		jsonData, err := json.MarshalIndent(task, "", "  ")
		if err != nil {
			slog.Error(err)
			return
		}

		addDataToExistingFile(filePath, jsonData)
	case List:
	case Update:
	case Delete:
	}
}

func jsonFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func addDataToExistingFile(path string, data []byte) {
	preData, err := os.ReadFile(path)
	if err != nil {
		slog.Error(err)
		return
	}

	if err := os.WriteFile(path, append(preData, data...), 0644); err != nil {
		slog.Error(err)
	}
}
