package users

import (
	"github.com/go-playground/validator/v10"

	"file-system/pkg/helpers"
)

type UserService interface {
	List() []UserResponse
	Retrieve(userId string) UserResponse
	Create(user CreateUserRequest) (userRetrieve UserRetrieveForToken)
	Update(userId string, user UpdateUserRequest)
	Delete(userId string)
}

type UserServiceImpl struct {
	userRepository UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		Validate:       validate,
	}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user CreateUserRequest) (userRetrieve UserRetrieveForToken) {
	err := u.Validate.Struct(user)
	helpers.ErrorHelper(err)

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		helpers.ErrorHelper(err)
	}

	userModel := UserCreate{
		FirstName: user.FirstName,
		Lastname:  user.Lastname,
		UserName:  user.UserName,
		Password:  hashedPassword,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	result := u.userRepository.Create(userModel)

	return result
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(userId string) {
	u.userRepository.Delete(userId)
}

// List implements UserService.
func (u *UserServiceImpl) List() []UserResponse {
	users := u.userRepository.List()

	var usersResult []UserResponse
	for _, user := range users {
		userResult := UserResponse{
			UserId:    user.ID.String(),
			FirstName: user.FirstName,
			Lastname:  user.Lastname,
			UserName:  user.UserName,
			Email:     user.Email,
			Phone:     user.Phone,
		}
		usersResult = append(usersResult, userResult)
	}
	return usersResult
}

// Retrieve implements UserService.
func (u *UserServiceImpl) Retrieve(userId string) UserResponse {
	user := u.userRepository.Retrieve(userId)
	return UserResponse{
		UserId:    user.ID.String(),
		FirstName: user.FirstName,
		Lastname:  user.Lastname,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}

// Update implements UserService.
func (u *UserServiceImpl) Update(userId string, user UpdateUserRequest) {
	err := u.Validate.Struct(user)
	helpers.ErrorHelper(err)

	userModel := UserUpdate{
		FirstName: user.FirstName,
		Lastname:  user.Lastname,
		UserName:  user.UserName,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	u.userRepository.Update(userId, userModel)
}
