package repositories

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"gorm.io/gorm"
)

type AuthUserRepository struct {
	db *gorm.DB
}

func NewAuthUserRepository(db *gorm.DB) *AuthUserRepository {
	return &AuthUserRepository{db: db}
}

func (r *AuthUserRepository) FindByUsername(username string) (*models.AuthUser, error) {
	var user models.AuthUser
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *AuthUserRepository) Create(user *models.AuthUser) error {
	return r.db.Create(user).Error
}
