package controllers

import (
	"book-app/app/models"
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

// UploadFile godoc
// @Summary File Uploader
// @Description File Uploader
// @Tags File
// @Accept mpfd
// @Param file formData file true "File to upload"
// @Produce json
// @Success 200
// @Router /v1/file [post]
// @Security ApiKeyAuth
// @Security JwtToken
func UploadFile(c echo.Context) error {
	acceptedTypes := []string{
		"image/png", "image/jpeg", "image/gif", "video/quicktime", "video/mp4",
		"application/pdf", "text/csv", "application/vnd.ms-excel",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.ms-excel.sheet.macroenabled.12",
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileType := file.Header.Get("Content-Type")
	extension := ".jpg"
	switch fileType {
	case "image/png":
		extension = ".png"
	case "image/jpeg":
		extension = ".jpg"
	case "image/gif":
		extension = ".gif"
	case "video/quicktime":
		extension = ".mov"
	case "video/mp4":
		extension = ".mp4"
	case "application/pdf":
		extension = ".pdf"
	case "application/vnd.ms-excel":
		extension = ".xls"
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		extension = ".xlsx"
	case "application/vnd.ms-excel.sheet.macroenabled.12":
		extension = ".xlsm"
	case "text/csv":
		extension = ".csv"
	}

	var isAccepted bool
	for _, t := range acceptedTypes {
		if t == fileType {
			isAccepted = true
			break
		}
	}

	if !isAccepted {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"accepted_type": acceptedTypes,
		})
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Initialize Cloudinary
	cld, err := cloudinary.NewFromParams(config.LoadConfig().CloudinaryName, config.LoadConfig().CloudinaryApiKey, config.LoadConfig().CloudinaryApiSecret)
	if err != nil {
		return err
	}

	// Upload file to Cloudinary
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	uploadResult, err := cld.Upload.Upload(c.Request().Context(), src, uploader.UploadParams{
		PublicID: fmt.Sprintf("uploads/%s%s", timestamp, extension),
	})
	if err != nil {
		return err
	}

	// Save the file URL to the database
	data, err := SaveFileToDatabase(uploadResult.SecureURL, uploadResult.OriginalFilename)
	if err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}

	// Update FullUrl in the response
	data.FullUrl = uploadResult.SecureURL

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "File uploaded successfully",
	})
}

func SaveFileToDatabase(filename, path string) (data models.GlobalFile, err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
		err = nil
	}
	t := time.Now().In(location).Unix()
	data = models.GlobalFile{
		Token:    strconv.Itoa(int(t)) + utils.GenerateRandomString(5),
		Filename: filename,
		Path:     path,
	}

	err = repository.SaveFile(&data)
	return
}

// GetFile godoc
// @Summary Mendapatkan List Files
// @Description Mendapatkan List Files
// @Tags File
// @Accept json
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param token query string false "token (string)"
// @Produce json
// @Success 200
// @Router /v1/file [get]

func GetFile(c echo.Context) error {
	userID := c.Get("user_id").(int)

	param := utils.PopulatePaging(c, "token")

	data, err := GetFileControl(userID, param)
	if err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}
	return c.JSON(http.StatusOK, data)
}

func GetFileControl(id int, param reqres.ReqPaging) (data reqres.ResPaging, err error) {

	data, err = repository.GetFile(id, param)
	if err != nil {
		return
	}
	return
}
