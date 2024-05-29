package files

import (
	"time"
)

type FileCreate struct {
	FileName  string     `bson:"file_name, omitempty"`
	FilePath  string     `bson:"file_path, omitempty"`
	CreatedAt time.Time  `bson:"created_at, omitempty"`
	UpdatedAt *time.Time `bson:"updated_at, omitempty"`
	DeletedAt *time.Time `bson:"deleted_at, omitempty"`
}

type FileUpdate struct {
	FileName  string     `bson:"file_name, omitempty"`
	FilePath  string     `bson:"file_path, omitempty"`
	CreatedAt *time.Time `bson:"created_at, omitempty"`
	UpdatedAt time.Time  `bson:"updated_at, omitempty"`
	DeletedAt *time.Time `bson:"deleted_at, omitempty"`
}

type FileDelete struct {
	CreatedAt *time.Time `bson:"created_at, omitempty"`
	UpdatedAt *time.Time `bson:"updated_at, omitempty"`
	DeletedAt time.Time  `bson:"deleted_at, omitempty"`
}

type RetrieveFile struct {
	FileName string `bson:"file_name, omitempty"`
	FilePath string `bson:"file_path, omitempty"`
}
