package wl

// Upload contains information about uploads.
// Uploads represent uploaded files.
type Upload struct {
	ID     uint64 `json:"id" yaml:"id"`
	UserID uint64 `json:"user_id" yaml:"user_id"`
	State  string `json:"state" yaml:"state"`
}
