package services

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
)

type BreakService interface {
	CreateBreak(breaks *models.Break) error
	GetBreakByID(id int) (*models.Break, error)
	GetAllBreaks() ([]models.Break, error)
	UpdateBreak(id int, input *models.Break) error
	DeleteBreak(id int) error
}

type breakService struct {
	repo repositories.BreakRepository
}

func NewBreakService(repo repositories.BreakRepository) BreakService {
	return &breakService{
		repo: repo,
	}
}

func (s *breakService) CreateBreak(breaks *models.Break) error {
	return s.repo.CreateBreak(breaks)
}

func (s *breakService) GetBreakByID(id int) (*models.Break, error) {
	return s.repo.GetBreakByID(id)
}

func (s *breakService) GetAllBreaks() ([]models.Break, error) {
	return s.repo.GetAllBreaks()
}

func (s *breakService) UpdateBreak(id int, input *models.Break) error {
	existingBreak, err := s.repo.GetBreakByID(id)
	if err != nil {
		return err
	}

	// Обновляем поля
	existingBreak.UserID = input.UserID
	existingBreak.BreakStart = input.BreakStart
	existingBreak.BreakEnd = input.BreakEnd
	// Другие поля, если есть

	return s.repo.UpdateBreak(existingBreak)
}

func (s *breakService) DeleteBreak(id int) error {
	return s.repo.DeleteBreak(id)
}
