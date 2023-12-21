package service

import (
	"errors"
	"github.com/khrees2412/simpledrive/model"
	"github.com/khrees2412/simpledrive/repositories"
	"github.com/khrees2412/simpledrive/util"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
)

type IFileService interface {
	Upload(userId string, file *multipart.FileHeader) error
	UploadToFolder(userId string, folderName string, file *multipart.FileHeader) error
	Download(userId string, fileName string) error
	Delete(userId string, fileId string) error
	GetAllFiles(userId string) ([]model.File, error)
	GetFilesByFolder(userId string, folderId string) error
}

type fileService struct {
	fileRepo   repositories.IFileRepository
	folderRepo repositories.IFolderRepository
}

// NewFileService will instantiate FileService
func NewFileService() IFileService {
	return &fileService{
		fileRepo:   repositories.NewFileRepo(),
		folderRepo: repositories.NewFolderRepo(),
	}
}

// kb 204800
var maxByteSize = 209700000 // 200 MB

func (s *fileService) Upload(userId string, file *multipart.FileHeader) error {
	if file.Size > int64(maxByteSize) {
		return errors.New("the file size is too large, try something below 200mb")
	}
	resp, err := util.UploadFile(file)
	if err != nil {
		return err
	}
	err = s.fileRepo.Create(&model.File{
		UserId:   userId,
		Name:     file.Filename,
		Size:     file.Size,
		FolderId: nil,
		Url:      resp.Location,
	})
	if err != nil {
		log.Errorf("unable to add file to database: %v", err)
		return errors.New("unable to add file to database")
	}
	return nil
}

func (s *fileService) UploadToFolder(userId string, folderName string, file *multipart.FileHeader) error {
	if file.Size > int64(maxByteSize) {
		return errors.New("the file size is too large, try something below 200mb")
	}
	resp, err := util.UploadFile(file)
	if err != nil {
		return err
	}
	folder, err := s.folderRepo.FindByFolderName(folderName)
	if err != nil {
		return errors.New("unable to find folder")
	}

	err = s.fileRepo.Create(&model.File{
		UserId:   userId,
		Name:     file.Filename,
		Size:     file.Size,
		FolderId: &folder.ID,
		Url:      resp.Location,
	})
	if err != nil {
		log.Errorf("unable to add file to database: %v", err)
		return errors.New("unable to add file to database")
	}
	return nil
}

func (s *fileService) Download(userId string, fileName string) error {

	return nil
}

func (s *fileService) Delete(userId string, fileId string) error {
	err := s.fileRepo.Delete(fileId)
	if err != nil {
		return err
	}
	return nil
}

func (s *fileService) GetAllFiles(userId string) ([]model.File, error) {
	files, err := s.fileRepo.FindAllFilesByUserId(userId)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *fileService) GetFilesByFolder(userId string, folderId string) error {
	return nil
}
