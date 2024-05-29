package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRetrieveForToken struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	UserName  string             `bson:"user_name, omitempty"`
	Role      string             `bson:"role, omitempty"`
	CreatedAt time.Time          `bson:"created_at, omitempty"`
}

type UserRetrieve struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	FirstName string             `bson:"first_name, omitempty"`
	Lastname  string             `bson:"last_name, omitempty"`
	UserName  string             `bson:"user_name, omitempty"`
	Password  string             `bson:"password, omitempty"`
	Email     string             `bson:"email, omitempty"`
	Phone     string             `bson:"phone, omitempty"`
}

type UserCreate struct {
	FirstName string     `bson:"first_name, omitempty"`
	Lastname  string     `bson:"last_name, omitempty"`
	UserName  string     `bson:"user_name, omitempty"`
	Password  string     `bson:"password, omitempty"`
	Email     string     `bson:"email, omitempty"`
	Phone     string     `bson:"phone, omitempty"`
	Role      string     `bson:"role, omitempty"`
	CreatedAt time.Time  `bson:"created_at, omitempty"`
	UpdatedAt *time.Time `bson:"updated_at, omitempty"`
	DeletedAt *time.Time `bson:"deleted_at, omitempty"`
}

type UserUpdate struct {
	FirstName string     `bson:"first_name, omitempty"`
	Lastname  string     `bson:"last_name, omitempty"`
	UserName  string     `bson:"user_name, omitempty"`
	Password  string     `bson:"password, omitempty"`
	Email     string     `bson:"email, omitempty"`
	Phone     string     `bson:"phone, omitempty"`
	CreatedAt *time.Time `bson:"created_at, omitempty"`
	UpdatedAt time.Time  `bson:"updated_at, omitempty"`
	DeletedAt *time.Time `bson:"deleted_at, omitempty"`
}

type UserDelete struct {
	CreatedAt *time.Time `bson:"created_at, omitempty"`
	UpdatedAt *time.Time `bson:"updated_at, omitempty"`
	DeletedAt time.Time  `bson:"deleted_at, omitempty"`
}
