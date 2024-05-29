package users

type UserResponse struct {
	UserId    string `json:"userId"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
