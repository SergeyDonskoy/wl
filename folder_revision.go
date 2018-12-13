package wl

// FolderRevision contains information about the revision of folder.
type FolderRevision struct {
	ID         uint64 `json:"id" yaml:"id"`
	TypeString string `json:"type" yaml:"type"`
	Revision   uint64 `json:"revision" yaml:"revision"`
}
