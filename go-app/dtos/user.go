package dtos

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type UserRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type Login struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
