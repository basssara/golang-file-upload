package files

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"file-system/pkg/helpers"
)

type FileRepository interface {
	List() []RetrieveFile
	Retrieve(fileId primitive.ObjectID) (RetrieveFile, error)
	Create(file FileCreate)
	Update(fileId primitive.ObjectID, file FileUpdate) error
	Delete(fileId primitive.ObjectID) error
}

type FileRepositoryImpl struct {
	Db *mongo.Collection
}

func NewFileRepositoryImpl(Db *mongo.Collection) FileRepository {
	return &FileRepositoryImpl{Db: Db}
}

// Create implements FileRepository.
func (f *FileRepositoryImpl) Create(file FileCreate) {
	fileModel := FileCreate{
		FileName:  file.FileName,
		FilePath:  file.FilePath,
		CreatedAt: time.Now(),
	}

	_, err := f.Db.InsertOne(context.TODO(), fileModel)

	helpers.ErrorHelper(err)
}

// Delete implements FileRepository.
func (f *FileRepositoryImpl) Delete(fileId primitive.ObjectID) error {
	panic("unimplemented")
}

// List implements FileRepository.
func (f *FileRepositoryImpl) List() []RetrieveFile {
	panic("unimplemented")
}

// Retrieve implements FileRepository.
func (f *FileRepositoryImpl) Retrieve(fileId primitive.ObjectID) (RetrieveFile, error) {
	panic("unimplemented")
}

// Update implements FileRepository.
func (f *FileRepositoryImpl) Update(fileId primitive.ObjectID, file FileUpdate) error {
	panic("unimplemented")
}
