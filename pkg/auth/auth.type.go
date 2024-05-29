package auth

type User struct {
	Email    string `validate:"required,min=5,max=100" json:"email"`
	Password string `validate:"required,min=5,max=255" json:"password"`
}
