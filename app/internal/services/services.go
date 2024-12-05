package services

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
)

type ServiceService interface {
	CreateService(service *models.Service) error
	GetServiceByID(id int) (*models.Service, error)
	GetAllServices() ([]models.Service, error)
	UpdateService(id int, input *models.Service) error
	DeleteService(id int) error
	DeactivateService(id int) error
}

type serviceService struct {
	repo repositories.ServiceRepository
}

func NewServiceService(repo repositories.ServiceRepository) ServiceService {
	return &serviceService{
		repo: repo,
	}
}

func (s *serviceService) CreateService(service *models.Service) error {
	// Дополнительная бизнес-логика перед созданием услуги (если требуется)
	return s.repo.CreateService(service)
}

func (s *serviceService) GetServiceByID(id int) (*models.Service, error) {
	return s.repo.GetServiceByID(id)
}

func (s *serviceService) GetAllServices() ([]models.Service, error) {
	return s.repo.GetAllServices()
}

func (s *serviceService) UpdateService(id int, input *models.Service) error {
	service, err := s.repo.GetServiceByID(id)
	if err != nil {
		return err
	}

	// Обновляем поля
	service.Name = input.Name
	service.Description = input.Description
	service.Price = input.Price
	service.Duration = input.Duration
	service.IsActive = input.IsActive

	return s.repo.UpdateService(service)
}

func (s *serviceService) DeleteService(id int) error {
	return s.repo.DeleteService(id)
}

func (s *serviceService) DeactivateService(id int) error {
	return s.repo.DeactivateService(id)
}
