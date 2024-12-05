package repositories

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"

	"gorm.io/gorm"
)

var (
	ErrServiceNotFound = errors.New("услуга не найдена")
)

type ServiceRepository interface {
	CreateService(service *models.Service) error
	GetServiceByID(id int) (*models.Service, error)
	GetAllServices() ([]models.Service, error)
	UpdateService(service *models.Service) error
	DeleteService(id int) error
	DeactivateService(id int) error
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{
		db: db,
	}
}

func (r *serviceRepository) CreateService(service *models.Service) error {
	return r.db.Create(service).Error
}

func (r *serviceRepository) GetServiceByID(id int) (*models.Service, error) {
	var service models.Service
	if err := r.db.First(&service, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, err
	}
	return &service, nil
}

func (r *serviceRepository) GetAllServices() ([]models.Service, error) {
	var services []models.Service
	if err := r.db.Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func (r *serviceRepository) UpdateService(service *models.Service) error {
	return r.db.Save(service).Error
}

func (r *serviceRepository) DeleteService(id int) error {
	if err := r.db.Delete(&models.Service{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *serviceRepository) DeactivateService(id int) error {
	if err := r.db.Model(&models.Service{}).Where("id = ?", id).Update("is_active", false).Error; err != nil {
		return err
	}
	return nil
}
