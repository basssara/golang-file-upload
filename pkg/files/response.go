package files

type FileResponse struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
}
