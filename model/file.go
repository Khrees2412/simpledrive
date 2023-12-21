package model

type File struct {
	Base
	Name     string  `json:"name" gorm:"unique"`
	UserId   string  `json:"user_id" gorm:"foreignKey:user_id"`
	Url      string  `json:"url"`
	Size     int64   `json:"size"`
	FolderId *string `json:"folder_id" gorm:"foreignKey:folder_id"`
}

type Folder struct {
	Base
	Name   string `json:"name" gorm:"unique"`
	UserId string `json:"user_id" gorm:"foreignKey:user_id"`
	Files  []File `json:"files"`
}
