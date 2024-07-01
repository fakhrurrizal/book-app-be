package controllers

import (
	"book-app/app/models"
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	// Load location
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
	}

	// Define the accepted MIME types
	acceptedTypes := []string{
		"image/png", "image/jpeg", "image/gif", "video/quicktime", "video/mp4",
		"application/pdf", "text/csv", "application/vnd.ms-excel",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.ms-excel.sheet.macroenabled.12",
	}

	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.Logger().Error("Error retrieving the file: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error retrieving the file"})
	}

	// Get the MIME type of the file
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
		c.Logger().Error("Unsupported file type: ", fileType)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"accepted_type": acceptedTypes,
		})
	}

	src, err := file.Open()
	if err != nil {
		c.Logger().Error("Error opening the file: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error opening the file"})
	}
	defer src.Close()

	// Create temp directory if not exists
	err = os.MkdirAll("/temp", os.ModePerm)
	if err != nil {
		c.Logger().Error("Error creating temp directory: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating temp directory"})
	}

	// Save file to temp directory
	tempFilePath := filepath.Join("/temp", file.Filename)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		c.Logger().Error("Error creating temp file: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating temp file"})
	}
	defer tempFile.Close()

	if _, err = io.Copy(tempFile, src); err != nil {
		c.Logger().Error("Error copying to temp file: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error copying to temp file"})
	}

	// Create directory if not exists
	t := time.Now().In(location)
	folder := t.Format("2006-01")
	uploadDir := filepath.Join(config.RootPath(), "assets/uploads", folder)
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		c.Logger().Error("Error creating directory: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating directory"})
	}

	timestamp := strconv.Itoa(int(t.Unix()))
	finalFilePath := filepath.Join(uploadDir, timestamp+extension)
	if err = os.Rename(tempFilePath, finalFilePath); err != nil {
		c.Logger().Error("Error moving file to final destination: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error moving file to final destination"})
	}

	data, err := SaveFileToDatabase(folder+"/"+timestamp+extension, finalFilePath)
	if err != nil {
		c.Logger().Error("Error saving file to database: ", err)
		return c.JSON(utils.ParseHttpError(err))
	}

	data.FullUrl = config.LoadConfig().BaseUrl + "/assets/uploads/" + folder + "/" + timestamp + extension
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Upload File Berhasil",
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
