package controllers

import (
	"book-app/app/models"
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// SignUp godoc
// @Summary SignUp
// @Description SignUp
// @Tags Auth
// @Accept json
// @Produce json
// @Param signup body reqres.SignUpRequest true "SignUp user"
// @Success 200
// @Router /v1/auth/signup [post]
func SignUp(c echo.Context) error {

	var request reqres.SignUpRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	utils.StripTagsFromStruct(&request)

	if err := request.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	_, err := repository.GetUserByEmail(strings.ToLower(request.Email))
	if err == nil {
		return c.JSON(http.StatusBadRequest, utils.Respond(http.StatusBadRequest, "bad request", "email sudah terdaftar"))
	}

	today := time.Now().Format("20060102")
	lastID, err := repository.GetLastUserID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Respond(http.StatusInternalServerError, "internal error", "gagal mengambil id terakhir"))
	}
	newID := fmt.Sprintf("%s%02d", today, lastID+1)

	inputUser := reqres.GlobalUserRequest{
		Fullname:  request.Fullname,
		Email:     request.Email,
		Password:  request.Password,
		Phone:     request.Phone,
		Address:   request.Address,
		Gender:    request.Gender,
		Avatar:    request.Avatar,
		ZipCode:   request.ZipCode,
		IdAnggota: newID,
		Village:   request.Village,
		District:  request.District,
		City:      request.City,
		Province:  request.Province,
		Country:   request.Country,
		RoleID:    request.RoleID,
	}

	_, err = repository.CreateUser(1, &inputUser, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestError([]map[string]interface{}{
			{
				"field": "Email",
				"error": err.Error(),
			},
		}))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "Pendaftaran anda berhasil",
	})

}

// SignIn godoc
// @Summary SignIn
// @Description SignIn
// @Tags Auth
// @Accept json
// @Param x-csrf-token header string false "csrf token"
// @Produce json
// @Param signin body reqres.SignInRequest true "SignIn user"
// @Success 200
// @Router /v1/auth/signin [post]
// @Security ApiKeyAuth
func SignIn(c echo.Context) error {

	var req reqres.SignInRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	user, accessToken, err := repository.SignIn(req.Email, req.Password)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status": 400,
			"error":  err.Error(),
		})
	}

	userData, _ := repository.GetUserByIDPlain(int(user.ID))

	var userResponse reqres.GlobalUserAuthResponse
	userResponse.ID = int(userData.ID)
	userResponse.Fullname = userData.Fullname
	userResponse.Email = userData.Email
	userResponse.Phone = userData.Phone
	userResponse.Address = userData.Address
	userResponse.Gender = userData.Gender
	userResponse.Avatar = userData.Avatar
	userResponse.ZipCode = userData.ZipCode
	userResponse.Village = userData.Village
	userResponse.District = userData.District
	userResponse.City = userData.City
	userResponse.Province = userData.Province
	userResponse.Country = userData.Country
	userResponse.IdAnggota = userData.IdAnggota
	userResponse.Role.ID = userData.RoleID

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"user":         userResponse,
			"access_token": accessToken,
			"expiration":   time.Now().Add(time.Hour * 72).Format("2006-01-02 15:04:05"),
		},
	})
}

// GetSignInUser godoc
// @Summary Get Sign In User
// @Description Get Sign In User
// @Tags Auth
// @Produce json
// @Success 200
// @Router /v1/auth/user [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetSignInUser(c echo.Context) error {

	id := c.Get("user_id").(int)
	user, err := repository.GetUserByIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to get user data"))
	}

	var roles models.GlobalRole

	var data reqres.GlobalUserAuthResponse
	data.ID = int(user.ID)

	data.Fullname = user.Fullname
	data.Avatar = user.Avatar
	data.Email = user.Email
	data.Phone = user.Phone
	data.Address = user.Address
	data.Village = user.Village
	data.District = user.District
	data.City = user.City
	data.Province = user.Province
	data.Country = user.Country
	data.ZipCode = user.ZipCode
	data.Status = user.Status
	data.Gender = user.Gender
	data.IdAnggota = user.IdAnggota

	if user.EmailVerifiedAt.Time.IsZero() {
		data.EmailVerified = false
	} else {
		data.EmailVerified = true
	}

	if user.RoleID > 0 {
		roles, _ = repository.GetRoleByIDPlain(user.RoleID)
		data.Role = reqres.GlobalIDNameResponse{
			ID:   int(roles.ID),
			Name: roles.Name,
		}
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Success to get user data",
	})
}
