package reqres

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BookCategoryRequest struct {
	Name        string `json:"name"  validate:"required"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Status      bool   `json:"status" `
}

func (request BookCategoryRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
	)
}

type BookCategoryResponse struct {
	CustomGormModel
	Name        string `json:"name"  validate:"required"`
	Description string `json:"description" `
	Icon        string `json:"icon"`
	Status      bool   `json:"status" `
}
