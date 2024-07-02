package middlewares

import (
	"book-app/app/models"
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"context"
	"fmt"
	"net/http"
	"path/filepath"
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
func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var method = c.Request().Method
		file, err := c.FormFile("image")
		if err != nil {
			if (method == "PUT" || method == "POST") && err.Error() == "http: no such file" {
				c.Set("cloudinarySecureURL", "")
				return next(c)
			}
			return c.JSON(http.StatusBadRequest, err)
		}

		ext := filepath.Ext(file.Filename)
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			var ctx = context.Background()
			cld, err := cloudinary.NewFromParams(
				config.LoadConfig().CloudinaryName,
				config.LoadConfig().CloudinaryApiKey,
				config.LoadConfig().CloudinaryApiSecret,
			)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to initialize Cloudinary: %v", err))
			}

			resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "book-app"})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to upload file to Cloudinary: %v", err))
			}

			fmt.Print("resp", resp)

			c.Set("cloudinarySecureURL", resp.SecureURL)
			return next(c)
		} else {
			return c.JSON(http.StatusBadRequest, "The file extension is wrong. Allowed file extensions are images (.png, .jpg, .jpeg, .webp)")
		}
	}
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
