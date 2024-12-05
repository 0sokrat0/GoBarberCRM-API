package repositories

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"

	"gorm.io/gorm"
)

var (
	ErrBreakNotFound = errors.New("перерыв не найден")
)

type BreakRepository interface {
	CreateBreak(breaks *models.Break) error
	GetBreakByID(id int) (*models.Break, error)
	GetAllBreaks() ([]models.Break, error)
	UpdateBreak(breaks *models.Break) error
	DeleteBreak(id int) error
}

type breakRepository struct {
	db *gorm.DB
}

func NewBreakRepository(db *gorm.DB) BreakRepository {
	return &breakRepository{
		db: db,
	}
}

func (r *breakRepository) CreateBreak(breaks *models.Break) error {
	return r.db.Create(breaks).Error
}

func (r *breakRepository) GetBreakByID(id int) (*models.Break, error) {
	var breakModel models.Break
	if err := r.db.First(&breakModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBreakNotFound
		}
		return nil, err
	}
	return &breakModel, nil
}

func (r *breakRepository) GetAllBreaks() ([]models.Break, error) {
	var breaks []models.Break
	if err := r.db.Find(&breaks).Error; err != nil {
		return nil, err
	}
	return breaks, nil
}

func (r *breakRepository) UpdateBreak(breaks *models.Break) error {
	return r.db.Save(breaks).Error
}

func (r *breakRepository) DeleteBreak(id int) error {
	if err := r.db.Delete(&models.Break{}, id).Error; err != nil {
		return err
	}
	return nil
}
