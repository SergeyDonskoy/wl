package wl

import "time"

// TaskComment represents information about a comment on a Task.
type TaskComment struct {
	ID        uint64    `json:"id" yaml:"id"`
	TaskID    uint64    `json:"task_id" yaml:"task_id"`
	Revision  uint64    `json:"revision" yaml:"revision"`
	Text      string    `json:"text" yaml:"text"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}
