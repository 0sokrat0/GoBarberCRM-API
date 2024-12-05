package services

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
)

type ClientService interface {
	CreateClient(client *models.Client) error
	GetClientByID(id int) (*models.Client, error)
	GetAllClients() ([]models.Client, error)
	UpdateClient(id int, input *models.Client) error
	DeleteClient(id int) error
	GetClientByTelegramID(tgID int64) (*models.Client, error)
	FilterClientsByName(name string) ([]models.Client, error)
	QuickAddClient(client *models.Client) error
	SearchClientByEmailOrPhone(email, phone string) (*models.Client, error)
	CheckClientExistence(phoneNumber string, tgID int64) (bool, error)
}

type clientService struct {
	repo repositories.ClientRepository
}

func NewClientService(repo repositories.ClientRepository) ClientService {
	return &clientService{
		repo: repo,
	}
}

func (s *clientService) CreateClient(client *models.Client) error {
	return s.repo.CreateClient(client)
}

func (s *clientService) GetClientByID(id int) (*models.Client, error) {
	return s.repo.GetClientByID(id)
}

func (s *clientService) GetAllClients() ([]models.Client, error) {
	return s.repo.GetAllClients()
}

func (s *clientService) UpdateClient(id int, input *models.Client) error {
	client, err := s.repo.GetClientByID(id)
	if err != nil {
		return err
	}

	// Обновляем поля
	client.FirstName = input.FirstName
	client.LastName = input.LastName
	client.Email = input.Email
	client.PhoneNumber = input.PhoneNumber
	client.TgID = input.TgID
	client.TgNickname = input.TgNickname

	return s.repo.UpdateClient(client)
}

func (s *clientService) DeleteClient(id int) error {
	return s.repo.DeleteClient(id)
}

func (s *clientService) GetClientByTelegramID(tgID int64) (*models.Client, error) {
	return s.repo.GetClientByTelegramID(tgID)
}

func (s *clientService) FilterClientsByName(name string) ([]models.Client, error) {
	return s.repo.FilterClientsByName(name)
}

func (s *clientService) QuickAddClient(client *models.Client) error {
	return s.repo.QuickAddClient(client)
}

func (s *clientService) SearchClientByEmailOrPhone(email, phone string) (*models.Client, error) {
	return s.repo.SearchClientByEmailOrPhone(email, phone)
}

func (s *clientService) CheckClientExistence(phoneNumber string, tgID int64) (bool, error) {
	return s.repo.CheckClientExistence(phoneNumber, tgID)
}
