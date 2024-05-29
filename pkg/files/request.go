package files

type RetrieveFileRequest struct {
	FileId string `json:"fileId"`
}

type CreateFileRequest struct {
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
}

type UpdateFileRequest struct {
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
}

type DeleteFileRequest struct {
	FileId string `json:"fileId"`
}
