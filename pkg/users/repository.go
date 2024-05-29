package users

import (
	"context"
	// "errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"file-system/pkg/helpers"
)

type UserRepository interface {
	List() []UserRetrieve
	Retrieve(userId string) UserRetrieve
	Create(user UserCreate) (userRetrieve UserRetrieveForToken)
	Update(userId string, user UserUpdate)
	Delete(userId string)
}

type UserRepositoryImpl struct {
	Db *mongo.Collection
}

func NewUserRepositoryImpl(Db *mongo.Collection) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Create implements UserRepository.
func (u *UserRepositoryImpl) Create(user UserCreate) (userRetrieve UserRetrieveForToken) {

	userModel := UserCreate{
		FirstName: user.FirstName,
		Lastname:  user.Lastname,
		UserName:  user.UserName,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      "user",
	}

	id, err := u.Db.InsertOne(context.TODO(), userModel)

	helpers.ErrorHelper(err)

	filter := bson.D{{Key: "_id", Value: id.InsertedID}}

	err = u.Db.FindOne(context.TODO(), filter).Decode(&userRetrieve)

	helpers.ErrorHelper(err)

	return userRetrieve
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId string) {
	id, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: time.Now()}}}}

	_, err := u.Db.UpdateOne(context.TODO(), filter, update)
	helpers.ErrorHelper(err)
}

// List implements UserRepository.
func (u *UserRepositoryImpl) List() []UserRetrieve {
	var users []UserRetrieve
	cursor, err := u.Db.Find(context.TODO(), bson.D{{Key: "deleted_at", Value: nil}})
	helpers.ErrorHelper(err)

	for cursor.Next(context.TODO()) {
		var user UserRetrieve
		err := cursor.Decode(&user)
		helpers.ErrorHelper(err)
		users = append(users, user)
	}

	return users
}

// Retrieve implements UserRepository.
func (u *UserRepositoryImpl) Retrieve(userId string) UserRetrieve {
	id, _ := primitive.ObjectIDFromHex(userId)

	var user UserRetrieve

	filter := bson.D{{Key: "_id", Value: id}, {Key: "deleted_at", Value: nil}}

	err := u.Db.FindOne(context.TODO(), filter).Decode(&user)

	log.Println(err)
	helpers.ErrorHelper(err)

	return user
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(userId string, user UserUpdate) {
	id, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "first_name", Value: user.FirstName}, {Key: "last_name", Value: user.Lastname}, {Key: "user_name", Value: user.UserName}, {Key: "email", Value: user.Email}, {Key: "phone", Value: user.Phone}, {Key: "updated_at", Value: time.Now()}}}}

	_, err := u.Db.UpdateOne(context.TODO(), filter, update)
	helpers.ErrorHelper(err)
}
