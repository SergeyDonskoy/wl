package wl

// Root contains information about the root of the object hierarchy.
type Root struct {
	ID       uint64 `json:"id" yaml:"id"`
	Revision uint64 `json:"revision" yaml:"revision"`
	UserID   uint64 `json:"user_id" yaml:"user_id"`
}
