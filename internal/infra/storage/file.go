package storage

type FileInformation struct {
	Name string // base filename
	Size int64  // file size in bytes
	Type string // MIME type
	Data []byte // file content
}
