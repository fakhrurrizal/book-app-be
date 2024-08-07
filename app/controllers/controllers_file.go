package controllers

import (
	"book-app/app/models"
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"net/http"
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
// func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		file, err := c.FormFile("file")
// 		if err != nil {
// 			return err
// 		}
// 		ext := filepath.Ext(file.Filename)
// 		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
// 			src, err := file.Open()
// 			if err != nil {
// 				return err
// 			}
// 			defer src.Close()

// 			// Initialize Cloudinary
// 			cld, err := cloudinary.NewFromParams(config.LoadConfig().CloudinaryName, config.LoadConfig().CloudinaryApiKey, config.LoadConfig().CloudinaryApiSecret)
// 			if err != nil {
// 				fmt.Print("cld", err)
// 				return err
// 			}

// 			var ctx = context.Background()
// 			uploadResult, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "book-app"})
// 			if err != nil {
// 				fmt.Print("uploadResult", err)
// 				return err
// 			}
// 			c.Set("cloudinarySecureURL", uploadResult.SecureURL)

// 			return next(c)
// 		}

// 	}

// }

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
