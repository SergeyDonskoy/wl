package wl

// Membership joins Users and Lists.
type Membership struct {
	ID       uint64 `json:"id" yaml:"id"`
	UserID   uint64 `json:"user_id" yaml:"user_id"`
	ListID   uint64 `json:"list_id" yaml:"list_id"`
	State    string `json:"state" yaml:"state"`
	Owner    bool   `json:"owner" yaml:"owner"`
	Muted    bool   `json:"muted" yaml:"muted"`
	Revision uint64 `json:"revision" yaml:"revision"`
}
