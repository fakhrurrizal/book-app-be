package repository

import (
	"book-app/app/middlewares"
	"book-app/app/models"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/guregu/null"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func CreateUser(status int, data *reqres.GlobalUserRequest, userID int) (response models.GlobalUser, err error) {
	if userID == 0 {
		var location *time.Location
		location, err = time.LoadLocation("Asia/Jakarta")
		if err != nil {
			location = time.Local
			err = nil
		}
		_, errUser := GetUserByEmail(strings.ToLower(data.Email))
		if errUser == nil {
			err = errors.New("email has been registered")
			return
		}

		response = models.GlobalUser{
			Avatar:    data.Avatar,
			Fullname:  data.Fullname,
			Email:     strings.ToLower(data.Email),
			Password:  middlewares.BcryptPassword(data.Password),
			Phone:     data.Phone,
			Address:   data.Address,
			IdAnggota: data.IdAnggota,
			Village:   data.Village,
			District:  data.District,
			City:      data.City,
			Province:  data.Province,
			Country:   data.Country,
			ZipCode:   data.ZipCode,
			Gender:    data.Gender,
			RoleID:    data.RoleID,
			Status:    status,
		}

		response.EmailVerifiedAt = null.TimeFrom(time.Now().In(location))

		var created bool
		for !created {
			respectiveID, _ := config.GetRespectiveID(config.DB, response.TableName(), true)
			response.ID = respectiveID
			err = config.DB.Create(&response).Error

			if err != nil {
				if !config.LoadConfig().EnableIDDuplicationHandling {
					return
				}
				if pqErr, ok := err.(*pq.Error); ok {
					if pqErr.Code != "23505" {
						return
					}
				}
			} else {
				created = true
			}
		}
	}

	return
}

func GetLastUserID() (int, error) {
	var user models.GlobalUser
	err := config.DB.Order("id DESC").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return int(user.ID), nil
}

func BuildUserResponse(data models.GlobalUser) (response reqres.GlobalUserResponse) {

	var roles models.GlobalRole

	response.CustomGormModel = data.CustomGormModel
	response.Avatar = data.Avatar
	response.Fullname = data.Fullname
	response.Email = strings.ToLower(data.Email)
	response.Phone = data.Phone
	response.Address = data.Address
	response.Village = data.Village
	response.District = data.District
	response.City = data.City
	response.Province = data.Province
	response.Country = data.Country
	response.ZipCode = data.ZipCode
	response.IdAnggota = data.IdAnggota
	response.Status = data.Status
	response.Gender = data.Gender

	if data.RoleID > 0 {
		roles, _ = GetRoleByIDPlain(data.RoleID)
		response.Role = reqres.GlobalIDNameResponse{
			ID:   int(roles.ID),
			Name: roles.Name,
		}
	}

	return response
}

func GetUsers(roleID, appID int, createdAtMarginTop, createdAtMarginBottom string, param reqres.ReqPaging) (data reqres.ResPaging) {
	var responses []models.GlobalUser
	where := "deleted_at IS NULL"

	if roleID > 0 {
		where += " AND role_id = " + strconv.Itoa(roleID)
	}

	var modelTotal []models.GlobalUser

	type TotalResult struct {
		Total       int64
		LastUpdated time.Time
	}
	var totalResult TotalResult
	config.DB.Model(&modelTotal).Select("COUNT(*) AS total, MAX(updated_at) AS last_updated").Scan(&totalResult)

	if createdAtMarginTop != "" {
		where += " AND created_at <= '" + createdAtMarginTop + "'"
	}
	if createdAtMarginBottom != "" {
		where += " AND created_at >= '" + createdAtMarginBottom + "'"
	}
	if param.Custom != "" {
		where += " AND status = " + param.Custom.(string)
	}

	if param.Search != "" {
		where += " AND (fullname ILIKE '%" + param.Search + "%' OR email ILIKE '%" + param.Search + "%')"
	}

	var totalFiltered int64
	config.DB.Model(&modelTotal).Where(where).Count(&totalFiltered)

	config.DB.Limit(param.Limit).Offset(param.Offset).Order(param.Sort + " " + param.Order).Where(where).Find(&responses)

	var responsesRefined []reqres.GlobalUserResponse
	for _, item := range responses {
		responseRefined := BuildUserResponse(item)

		responsesRefined = append(responsesRefined, responseRefined)
	}

	data = utils.PopulateResPaging(&param, responsesRefined, totalResult.Total, totalFiltered)

	return
}

func GetAllUsers(companyID, status int) (users []models.GlobalUser, err error) {
	where := "deleted_at IS NULL"
	// if appID > 0 {
	// 	where += " AND app_id = " + strconv.Itoa(appID)
	// }
	if companyID > 0 {
		where += " AND company_id = " + strconv.Itoa(companyID)
	}
	// if roleID > 0 {
	// 	where += " AND role_id = " + strconv.Itoa(roleID)
	// }
	if status > 0 {
		where += " AND status = " + strconv.Itoa(status)
	}

	err = config.DB.Where(where).Find(&users).Error

	return
}

func GetUserByID(id int) (response reqres.GlobalUserResponse, err error) {
	var data models.GlobalUser
	err = config.DB.First(&data, id).Error

	response = BuildUserResponse(data)

	return
}

func GetUserByIDPlain(id int) (response models.GlobalUser, err error) {
	err = config.DB.First(&response, id).Error

	return
}

func GetUserByEmail(email string) (response models.GlobalUser, err error) {
	err = config.DB.Where("email = ?", strings.ToLower(email)).First(&response).Error

	return
}

func DeleteUser(request models.GlobalUser) (models.GlobalUser, error) {
	var err error
	err = config.DB.Delete(&request).Error

	return request, err
}

func UpdateUser(request models.GlobalUser) (response models.GlobalUser, err error) {
	err = config.DB.Save(&request).Scan(&response).Error

	return
}
