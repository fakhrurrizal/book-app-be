package repository

import (
	"book-app/app/models"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"strconv"
	"time"

	"github.com/lib/pq"
)

func SaveFile(data *models.GlobalFile) (err error) {
	var created bool
	for !created {
		err = config.DB.Create(&data).Error
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					// Handle unique violation error, for example, regenerate token
					data.Token = strconv.Itoa(int(time.Now().Unix())) + utils.GenerateRandomString(5)
				} else {
					created = true
				}
			} else {
				created = true
			}
		} else {
			created = true
		}
	}
	return
}

func GetFile(id int, param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	search := "user_id = " + strconv.Itoa(id)
	if param.Search != "" {
		search = " AND filename ILIKE '%" + param.Search + "%' "
	}
	token := param.Custom.(string)
	if token != "" {
		search += " AND token ILIKE '%" + token + "%' "
	}
	var files []models.GlobalFile
	err = config.DB.Where(search).Order(param.Sort + " " + param.Order).Limit(param.Limit).Offset(param.Offset).Find(&files).Error
	if err != nil {
		return
	}

	var modelTotal []models.GlobalFile

	var totalResult int64
	config.DB.Model(&modelTotal).Count(&totalResult)

	var totalFiltered int64
	config.DB.Model(&modelTotal).Where(search).Count(&totalFiltered)

	for i, file := range files {
		files[i].FullUrl = config.LoadConfig().BaseUrl + "/assets/uploads/" + file.Filename
	}
	data = utils.PopulateResPaging(&param, files, totalResult, totalFiltered)
	return
}

func GetFileByToken(token string, id int64) (data models.GlobalFile, err error) {
	err = config.DB.Where("token = ?", token).First(&data).Error

	return
}
