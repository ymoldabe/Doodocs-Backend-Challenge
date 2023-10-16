package models

type ArhiveInfo struct {
	FileName    string  `json:"filename"`
	ArchiveSize float64 `json:"archive_size"`
	TotalSize   float64 `json:"total_size"`
	TotalFiles  float64 `json:"total_files"`
	Files       []Files `json:"files"`
}

type Files struct {
	FilePath string  `json:"file_path"`
	Size     float64 `json:"size"`
	Mimetype string  `json:"mimetype"`
}
