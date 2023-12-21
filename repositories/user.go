package repositories

import (
	"github.com/khrees2412/simpledrive/database"
	"github.com/khrees2412/simpledrive/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *model.User) error
	FindByUserId(userId string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	DoesEmailExist(email string) (bool, error)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo will instantiate User Repository
func NewUserRepo() IUserRepository {
	return &userRepo{
		db: database.DB(),
	}
}

func (r *userRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByUserId(userId string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) DoesEmailExist(email string) (bool, error) {
	var user model.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *userRepo) Update(user *model.User) error {
	return r.db.Save(user).Error
}
