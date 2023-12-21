package repositories

import (
	"github.com/khrees2412/simpledrive/database"
	"github.com/khrees2412/simpledrive/model"
	"gorm.io/gorm"
)

type IFileRepository interface {
	Create(file *model.File) error
	FindByFileId(fileId string) (*model.File, error)
	FindAllFilesByUserId(userId string) ([]model.File, error)
	Update(file *model.File) error
	Delete(fileId string) error
}

type fileRepo struct {
	db *gorm.DB
}

// NewFileRepo will instantiate File Repository
func NewFileRepo() IFileRepository {
	return &fileRepo{
		db: database.DB(),
	}
}

func (r *fileRepo) Create(file *model.File) error {
	return r.db.Create(file).Error
}

func (r *fileRepo) FindByFileId(fileId string) (*model.File, error) {
	var file model.File
	if err := r.db.Where("file_id = ?", fileId).First(&file).Error; err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *fileRepo) FindAllFilesByUserId(userId string) ([]model.File, error) {
	var files []model.File
	if err := r.db.Where("user_id = ?", userId).Find(&files).Error; err != nil {
		return nil, err
	}

	return files, nil
}

func (r *fileRepo) Update(file *model.File) error {
	return r.db.Save(file).Error
}

func (r *fileRepo) Delete(fileId string) error {
	return r.db.Where("file_id = ?", fileId).Delete(&model.File{}).Error
}
