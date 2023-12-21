package service

import (
	"github.com/khrees2412/simpledrive/repositories"
	"mime/multipart"
)

type IFileService interface {
	Upload(userId string, file *multipart.FileHeader) error
	UploadToFolder(userId string, folderName string, file *multipart.FileHeader) error
	Download(userId string, fileName string) error
	Delete(userId string, fileId string) error
	GetAllFiles(userId string) error
	GetFilesByFolder(userId string, folderId string) error
}

type fileService struct {
	fileRepo repositories.IFileRepository
}

// NewFileService will instantiate FileService
func NewFileService() IFileService {
	return &fileService{
		fileRepo: repositories.NewFileRepo(),
	}
}

// kb 204800
var maxByteSize = 209700000 // 200 MB

func (s *fileService) Upload(userId string, file *multipart.FileHeader) error {

	return nil
}

func (s *fileService) UploadToFolder(userId string, folderName string, file *multipart.FileHeader) error {
	return nil
}

func (s *fileService) Download(userId string, fileName string) error {
	return nil
}

func (s *fileService) Delete(userId string, fileId string) error {
	return nil
}

func (s *fileService) GetAllFiles(userId string) error {
	return nil
}

func (s *fileService) GetFilesByFolder(userId string, folderId string) error {
	return nil
}
