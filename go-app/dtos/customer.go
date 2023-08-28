package dtos

type Customer struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

type GetCustomersParams struct {
	Page        uint64
	DataPerPage uint64
}
