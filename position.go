package wl

// Position contains an ordered list of IDs of Lists, Tasks or Subtasks.
type Position struct {
	ID       uint64   `json:"id" yaml:"id"`
	Values   []uint64 `json:"values" yaml:"values"`
	Revision uint64   `json:"revision" yaml:"revision"`
}
