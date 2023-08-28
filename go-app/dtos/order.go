package dtos

type Order struct {
	Name      string `json:"name" validate:"required"`
	UserEmail string `json:"user_email" validate:"required"`
}

type GetOrdersParams struct {
	Page        uint64
	DataPerPage uint64
}
