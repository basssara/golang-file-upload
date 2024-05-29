package files

import (
	"github.com/go-playground/validator/v10"

	"file-system/pkg/helpers"
)

type FileService interface {
	List() []RetrieveFile
	Retrieve(fileId string) RetrieveFile
	Create(file CreateFileRequest)
	Update(fileId string, file UpdateFileRequest)
	Delete(fileId string)
}

type FileServiceImpl struct {
	fileRepository FileRepository
	Validate       *validator.Validate
}

func NewFileServiceImpl(fileRepository FileRepository, validate *validator.Validate) FileService {
	return &FileServiceImpl{
		fileRepository: fileRepository,
		Validate:       validate,
	}
}

// Create implements FileService.
func (f *FileServiceImpl) Create(file CreateFileRequest) {
	err := f.Validate.Struct(file)
	helpers.ErrorHelper(err)

	fileModel := FileCreate{
		FileName: file.FileName,
		FilePath: file.FilePath,
	}

	f.fileRepository.Create(fileModel)
}

// Delete implements FileService.
func (f *FileServiceImpl) Delete(fileId string) {
	panic("unimplemented")
}

// List implements FileService.
func (f *FileServiceImpl) List() []RetrieveFile {
	panic("unimplemented")
}

// Retrieve implements FileService.
func (f *FileServiceImpl) Retrieve(fileId string) RetrieveFile {
	panic("unimplemented")
}

// Update implements FileService.
func (f *FileServiceImpl) Update(fileId string, file UpdateFileRequest) {
	panic("unimplemented")
}
