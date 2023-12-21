package model

type File struct {
	Base
	Name     string `json:"name"`
	UserID   string `json:"user_id" gorm:"foreignKey:user_id"`
	Url      string `json:"url"`
	FolderID string `json:"folder_id" gorm:"foreignKey:folder_id"`
}

type Folder struct {
	Base
	Name   string `json:"name" gorm:"unique"`
	UserID string `json:"user_id" gorm:"foreignKey:user_id"`
	Files  []File `json:"files"`
}
