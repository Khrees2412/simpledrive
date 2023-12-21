package repositories

import (
	"github.com/khrees2412/simpledrive/database"
	"github.com/khrees2412/simpledrive/model"
	"gorm.io/gorm"
)

type IFolderRepository interface {
	Create(folder *model.Folder) error
	FindByFolderId(folderId string) (*model.Folder, error)
	FindByFolderName(folderName string) (*model.Folder, error)
	FindAllFoldersByUserId(userId string) ([]model.Folder, error)
	Update(folder *model.Folder) error
	Delete(folderId string) error
}

type folderRepo struct {
	db *gorm.DB
}

// NewFolderRepo will instantiate Folder Repository
func NewFolderRepo() IFolderRepository {
	return &folderRepo{
		db: database.DB(),
	}
}

func (r *folderRepo) Create(folder *model.Folder) error {
	return r.db.Create(folder).Error
}

func (r *folderRepo) FindByFolderId(folderId string) (*model.Folder, error) {
	var folder model.Folder
	if err := r.db.Where("folder_id = ?", folderId).First(&folder).Error; err != nil {
		return nil, err
	}

	return &folder, nil
}
func (r *folderRepo) FindByFolderName(folderName string) (*model.Folder, error) {
	var folder model.Folder
	if err := r.db.Where("folder_name = ?", folderName).First(&folder).Error; err != nil {
		return nil, err
	}

	return &folder, nil
}

func (r *folderRepo) FindAllFoldersByUserId(userId string) ([]model.Folder, error) {
	var folders []model.Folder
	if err := r.db.Where("user_id = ?", userId).Find(&folders).Error; err != nil {
		return nil, err
	}

	return folders, nil
}

func (r *folderRepo) Update(folder *model.Folder) error {
	return r.db.Save(folder).Error
}

func (r *folderRepo) Delete(folderId string) error {
	return r.db.Where("folder_id = ?", folderId).Delete(&model.Folder{}).Error
}
