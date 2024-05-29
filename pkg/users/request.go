package users

type GetUserRequest struct {
	UserId string `json:"userId"`
}

type CreateUserRequest struct {
	FirstName string `validate:"required,min=2,max=50" json:"firstName"`
	Lastname  string `validate:"required,min=2,max=50" json:"lastName"`
	UserName  string `validate:"required,min=5,max=100" json:"username"`
	Password  string `validate:"required,min=5,max=255" json:"password"`
	Email     string `validate:"required,min=5,max=100" json:"email"`
	Phone     string `validate:"required,min=5,max=255" json:"phone"`
}

type UpdateUserRequest struct {
	FirstName string `validate:"omitempty,min=2,max=50" json:"firstName"`
	Lastname  string `validate:"omitempty,min=2,max=50" json:"lastName"`
	UserName  string `validate:"omitempty,min=5,max=100" json:"username"`
	Password  string `validate:"omitempty,min=5,max=255" json:"password"`
	Email     string `validate:"omitempty,min=5,max=100" json:"email"`
	Phone     string `validate:"omitempty,min=5,max=255" json:"phone"`
}

type DeleteUserRequest struct {
	UserId string `json:"userId"`
}
