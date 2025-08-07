package reqres

import (
	"book-app/app/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/guregu/null"
)

type SignUpRequest struct {
	Avatar   string `json:"avatar"`
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Village  string `json:"village"`
	District string `json:"district"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	ZipCode  string `json:"zip_code"`
	RoleID   int    `json:"role_id"`
	Gender   string `json:"gender"`
}

func (request SignUpRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Length(8, 30)),
		validation.Field(&request.Fullname, validation.Required, validation.Length(5, 50)),
		validation.Field(&request.Phone, validation.Length(7, 17).Error("Nomor telepon harus benar")),
		validation.Field(&request.Address, validation.Length(5, 100)),
		validation.Field(&request.Gender, validation.In("m", "f").Error("Gender harus m(Laki) atau f(perempuan)")),
	)
}

type GlobalUserRequest struct {
	Avatar       string `json:"avatar"`
	Fullname     string `json:"fullname" validate:"required"`
	Gender       string `json:"gender"`
	Email        string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Code         string `json:"code"`
	Village      string `json:"village"`
	District     string `json:"district"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Experience   string `json:"experience"`
	ZipCode      string `json:"zip_code"`
	RoleID       int    `json:"role_id"`
	Status       int    `json:"status"`
	AutoVerified bool   `json:"auto_verified"`
}

func (request GlobalUserRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Fullname, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
}

type GlobalUserResponse struct {
	models.CustomGormModel
	Avatar                string               `json:"avatar"`
	Fullname              string               `json:"fullname"`
	Email                 string               `json:"email"`
	Password              string               `json:"-"`
	Phone                 string               `json:"phone"`
	Address               string               `json:"address"`
	Village               string               `json:"village"`
	District              string               `json:"district"`
	City                  string               `json:"city"`
	Gender                string               `json:"gender"`
	Province              string               `json:"province"`
	Country               string               `json:"country"`
	ZipCode               string               `json:"zip_code"`
	Status                int                  `json:"status"`
	Role                  GlobalIDNameResponse `json:"role"`
	Department            GlobalIDNameResponse `json:"department"`
	Branch                GlobalIDNameResponse `json:"branch"`
	Plant                 GlobalIDNameResponse `json:"plant"`
	Company               GlobalIDNameResponse `json:"company"`
	App                   GlobalIDNameResponse `json:"app"`
	EmailVerifiedAt       null.Time            `json:"-"`
	EmployeeID            string               `json:"employee_id,omitempty"`
	Payday                null.Time            `json:"payday,omitempty"`
	LastSalary            float64              `json:"last_salary,omitempty"`
	BackofficeAccess      bool                 `json:"backoffice_access"`
	TotalSaleQuantity     int                  `json:"total_sale_quantity"`
	TotalSalePrice        int                  `json:"total_sale_price"`
	TotalPurchaseQuantity int                  `json:"total_purchase_quantity"`
	TotalPurchasePrice    int                  `json:"total_purchase_price"`
	CompanyRole           GlobalIDNameResponse `json:"company_role"`
}
