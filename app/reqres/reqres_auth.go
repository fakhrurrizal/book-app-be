package reqres

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (request SignInRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
}

type EmailRequest struct {
	Email string `json:"email" validate:"required"`
}

func (request EmailRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
	)
}

type TokenRequest struct {
	Token string `json:"token" validate:"required"`
}

func (request TokenRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Token, validation.Required),
	)
}

type GlobalUserAuthResponse struct {
	ID            int                  `json:"id"`
	EncodedID     string               `json:"encoded_id"`
	Fullname      string               `json:"fullname"`
	Email         string               `json:"email"`
	Avatar        string               `json:"avatar"`
	Phone         string               `json:"phone"`
	Address       string               `json:"address"`
	Village       string               `json:"village"`
	District      string               `json:"district"`
	City          string               `json:"city"`
	Province      string               `json:"province"`
	Country       string               `json:"country"`
	ZipCode       string               `json:"zip_code"`
	IdAnggota     string               `json:"id_anggota"`
	Status        int                  `json:"status"`
	Gender        string               `json:"gender"`
	Role          GlobalIDNameResponse `json:"role"`
	Branch        GlobalIDNameResponse `json:"branch,omitempty"`
	EmailVerified bool                 `json:"email_verified"`
}

type GlobalUserAuthCompanyResponse struct {
	ID              int                                       `json:"id"`
	EncodedID       string                                    `json:"encoded_id,omitempty"`
	Name            string                                    `json:"name"`
	BusinessType    GlobalUserAuthCompanyBusinessTypeResponse `json:"business_type"`
	IsStockDisabled bool                                      `json:"is_stock_disabled"`
}

type GlobalUserAuthCompanyBusinessTypeResponse struct {
	ID        int    `json:"id"`
	EncodedID string `json:"encoded_id,omitempty"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsService bool   `json:"is_service"`
}
