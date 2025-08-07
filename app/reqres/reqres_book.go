package reqres

import (
	"book-app/app/models"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type BookRequest struct {
	Title           string `json:"title" validate:"required"`
	Author          string `json:"author"`
	Publisher       string `json:"publisher"`
	BookCode        string `json:"book_code"`
	Image           string `json:"image"`
	PublicationYear int    `json:"publication_year"`
	Language        string `json:"language"`
	Description     string `json:"description"`
	NumberOfPages   int    `json:"number_of_pages"`
	CategoryID      int    `json:"category_id" validate:"required"`
	Status          bool   `json:"status"`
}

type CustomGormModel struct {
	ID        uint            `gorm:"primary_key" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type GlobalIDNameResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Image string `json:"image,omitempty"`
}

func (request BookRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Title, validation.Required),
	)
}

type BookResponse struct {
	models.CustomGormModel
	Title           string               `json:"title"`
	Author          string               `json:"author"`
	BookCode        string               `json:"book_code"`
	Publisher       string               `json:"publisher"`
	Image           string               `json:"image"`
	PublicationYear int                  `json:"publication_year"`
	Language        string               `json:"language"`
	Description     string               `json:"description"`
	NumberOfPages   int                  `json:"number_of_pages"`
	Status          bool                 `json:"status"`
	Category        GlobalIDNameResponse `json:"category" `
}
