package wl

// ListTaskCount contains information about the number and type of tasks
// a List contains.
type ListTaskCount struct {
	ID               uint64 `json:"id" yaml:"id"`
	CompletedCount   uint64 `json:"completed_count" yaml:"completed_count"`
	UncompletedCount uint64 `json:"uncompleted_count" yaml:"uncompleted_count"`
	TypeString       string `json:"type" yaml:"type"`
}
