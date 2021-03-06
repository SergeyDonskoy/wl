package wl

import "time"

// File contains the information about an uploaded file.
// See also Upload.
type File struct {
	ID             uint64    `json:"id" yaml:"id"`
	URL            string    `json:"url" yaml:"url"`
	TaskID         uint64    `json:"task_id" yaml:"task_id"`
	ListID         uint64    `json:"list_id" yaml:"list_id"`
	UserID         uint64    `json:"user_id" yaml:"user_id"`
	FileName       string    `json:"file_name" yaml:"file_name"`
	ContentType    string    `json:"content_type" yaml:"content_type"`
	FileSize       int       `json:"file_size" yaml:"file_size"`
	LocalCreatedAt time.Time `json:"local_created_at" yaml:"local_created_at"`
	CreatedAt      time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" yaml:"updated_at"`
	Type           string    `json:"type" yaml:"type"`
	Revision       uint64    `json:"revision" yaml:"revision"`
}
