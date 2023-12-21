package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrees2412/simpledrive/service"
	"github.com/khrees2412/simpledrive/types"
	"github.com/khrees2412/simpledrive/util"
)

type IFileController interface {
	Upload(ctx *fiber.Ctx) error
	Download(ctx *fiber.Ctx) error
	RegisterRoutes(app *fiber.App)
}

type fileController struct {
	fileService service.IFileService
}

// NewFileController instantiates File Controller
func NewFileController() IFileController {
	return &fileController{
		fileService: service.NewFileService(),
	}
}

func (ctl *fileController) RegisterRoutes(app *fiber.App) {
	file := app.Group("/v1/file")

	file.Post("/upload", util.SecureAuth(), ctl.Upload)
	file.Post("/download", util.SecureAuth(), ctl.Download)
}

func (ctl *fileController) Upload(ctx *fiber.Ctx) error {
	userId, err := util.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	file, e := ctx.FormFile("file")
	if e != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: e.Error(),
		})
	}
	err = ctl.fileService.Upload(userId, file)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(types.GenericResponse{
		Success: true,
		Message: "File uploaded successfully",
		Data:    file.Filename,
	})
}

func (ctl *fileController) Download(ctx *fiber.Ctx) error {
	return nil
}

//
//import (
//	"fmt"
//
//	db "github.com/khrees2412/simpledrive/database"
//	"github.com/khrees2412/simpledrive/model"
//	"github.com/khrees2412/simpledrive/types"
//	"github.com/khrees2412/simpledrive/util"
//
//	"github.com/gofiber/fiber/v2"
//	log "github.com/sirupsen/logrus"
//)
//
//// kb 204800
//var maxByteSize = 209700000 // 200 MB
//
//func StoreFileInFolder(c *fiber.Ctx) error {
//	userId := fmt.Sprintf("%s", c.Locals("id"))
//	folderId := c.Params("folder")
//	if folderId == "" {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "You must specify a folder",
//		})
//	}
//
//	var folder model.Folder
//	if err := db.DB.Where("user_id = ? AND id = ?", userId, folderId).First(&folder).Error; err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "folder does not exist, you need to create a folder",
//		})
//	}
//
//	file, err := c.FormFile("file")
//	if err != nil {
//		log.Error(err)
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "invalid file",
//		})
//	}
//	filesize := file.Size
//	filename := file.Filename
//	if filesize > int64(maxByteSize) {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "The file size is too large, try something below 200mb",
//		})
//	}
//	data, err := util.UploadFile(file)
//
//	if err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "File upload failed",
//			Data:    err.Error(),
//		})
//	}
//
//	newFile := model.File{
//		Name:     filename,
//		UserID:   userId,
//		Url:      data.Location,
//		FolderID: folderId,
//	}
//
//	if err = db.DB.Create(&newFile).Error; err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "File upload failed",
//		})
//	}
//
//	return c.JSON(types.GenericResponse{
//		Success: true,
//		Message: fmt.Sprintf("successfully uploaded %s", filename),
//		Data:    data.Location,
//	})
//}
//func StoreFile(c *fiber.Ctx) error {
//	userId := fmt.Sprintf("%s", c.Locals("id"))
//
//	file, err := c.FormFile("file")
//	if err != nil {
//		log.Error(err)
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "Invalid file",
//			Data:    err.Error(),
//		})
//	}
//	filesize := file.Size
//	filename := file.Filename

//	data, err := util.UploadFile(file)
//
//	if err != nil {
//		log.Println(err)
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "File upload failed",
//		})
//	}
//
//	newFile := model.File{
//		Name:   filename,
//		UserID: userId,
//		Url:    data.Location,
//	}
//
//	if err = db.DB.Create(&newFile).Error; err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "File upload failed",
//		})
//	}
//
//	return c.JSON(types.GenericResponse{
//		Success: true,
//		Message: fmt.Sprintf("successfully uploaded %s", filename),
//		Data:    data.Location,
//	})
//}
//func GetFiles(c *fiber.Ctx) error {
//	userId := fmt.Sprintf("%s", c.Locals("id"))
//	var files []model.File
//
//	err := db.DB.Where("id = ?", userId).Find(&files).Error
//	if err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "Unable to find files for user",
//		})
//	}
//	return c.JSON(types.GenericResponse{
//		Success: true,
//		Message: "Successfully retrieved files",
//		Data:    files,
//	})
//
//}
//
//func GetFile(c *fiber.Ctx) error {
//	fileName := c.Params("filename")
//	userId := fmt.Sprintf("%s", c.Locals("id"))
//
//	var file model.File
//	err := db.DB.Where("id = ? AND name = ?", userId, fileName).Find(&file).Error
//	if err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "File not found",
//		})
//	}
//	return c.JSON(types.GenericResponse{
//		Success: true,
//		Message: "File successfully retrieved",
//		Data:    file,
//	})
//
//}
//
//func DeleteFile(c *fiber.Ctx) error {
//	fileName := c.Params("filename")
//	userId := fmt.Sprintf("%s", c.Locals("id"))
//
//	var file model.File
//	err := db.DB.Where("id = ? AND name = ?", userId, fileName).Delete(&file).Error
//	if err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "File couldn't be deleted",
//		})
//	}
//	return c.JSON(types.GenericResponse{
//		Success: true,
//		Message: fmt.Sprintf("Deleted %s successfully", fileName),
//	})
//
//}
//
//func DownloadFile(c *fiber.Ctx) error {
//	fileName := c.Params("filename")
//	userId := fmt.Sprintf("%s", c.Locals("id"))
//	var file model.File
//	err := db.DB.Where("id = ? AND name = ?", userId, fileName).First(&file).Error
//	if err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "File couldn't be downloaded",
//		})
//	}
//	f := file.Url
//	return c.Download(f)
//}
