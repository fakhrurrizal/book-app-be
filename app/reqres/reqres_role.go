package reqres

import (
	"book-app/app/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GlobalRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (request GlobalRoleRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
	)
}

type GlobalRoleResponse struct {
	models.CustomGormModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type GlobalRoleUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
