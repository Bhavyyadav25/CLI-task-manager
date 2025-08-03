package domain

import (
	"time"

	"github.com/google/uuid"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	ErrEmptyDescription = &Error{
		Code:    "empty_description",
		Message: "description cannot be empty",
	}
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

func (task *Task) ValidateTask() Error {
	if task.Description == "" {
		return *ErrEmptyDescription
	}
	return Error{}
}
