package repositories

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"

	"gorm.io/gorm"
)

var (
	ErrClientNotFound      = errors.New("клиент не найден")
	ErrClientAlreadyExists = errors.New("клиент уже существует")
)

type ClientRepository interface {
	CreateClient(client *models.Client) error
	GetClientByID(id int) (*models.Client, error)
	GetAllClients() ([]models.Client, error)
	UpdateClient(client *models.Client) error
	DeleteClient(id int) error
	GetClientByTelegramID(tgID int64) (*models.Client, error)
	FilterClientsByName(name string) ([]models.Client, error)
	QuickAddClient(client *models.Client) error
	SearchClientByEmailOrPhone(email, phone string) (*models.Client, error)
	CheckClientExistence(phoneNumber string, tgID int64) (bool, error)
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{
		db: db,
	}
}

func (r *clientRepository) CreateClient(client *models.Client) error {
	return r.db.Create(client).Error
}

func (r *clientRepository) GetClientByID(id int) (*models.Client, error) {
	var client models.Client
	if err := r.db.First(&client, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	if err := r.db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *clientRepository) UpdateClient(client *models.Client) error {
	return r.db.Save(client).Error
}

func (r *clientRepository) DeleteClient(id int) error {
	if err := r.db.Delete(&models.Client{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *clientRepository) GetClientByTelegramID(tgID int64) (*models.Client, error) {
	var client models.Client
	if err := r.db.Where("tg_id = ?", tgID).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) FilterClientsByName(name string) ([]models.Client, error) {
	var clients []models.Client
	if err := r.db.Where("LOWER(first_name) LIKE LOWER(?) OR LOWER(last_name) LIKE LOWER(?)", "%"+name+"%", "%"+name+"%").Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *clientRepository) QuickAddClient(client *models.Client) error {
	// Проверяем, что хотя бы одно из обязательных полей указано
	if client.PhoneNumber == "" && client.TgID == 0 {
		return errors.New("номер телефона или Telegram ID обязательны")
	}

	// Проверяем на существование клиента по номеру телефона или Telegram ID
	var existingClient models.Client
	if client.PhoneNumber != "" {
		if err := r.db.Where("phone_number = ?", client.PhoneNumber).First(&existingClient).Error; err == nil {
			return ErrClientAlreadyExists
		}
	}
	if client.TgID != 0 {
		if err := r.db.Where("tg_id = ?", client.TgID).First(&existingClient).Error; err == nil {
			return ErrClientAlreadyExists
		}
	}

	return r.db.Create(client).Error
}

func (r *clientRepository) SearchClientByEmailOrPhone(email, phone string) (*models.Client, error) {
	var client models.Client
	if email != "" {
		if err := r.db.Where("email = ?", email).First(&client).Error; err == nil {
			return &client, nil
		}
	}
	if phone != "" {
		if err := r.db.Where("phone_number = ?", phone).First(&client).Error; err == nil {
			return &client, nil
		}
	}
	return nil, ErrClientNotFound
}

func (r *clientRepository) CheckClientExistence(phoneNumber string, tgID int64) (bool, error) {
	var count int64
	if phoneNumber != "" {
		err := r.db.Model(&models.Client{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
		if err != nil {
			return false, err
		}
	} else if tgID != 0 {
		err := r.db.Model(&models.Client{}).Where("tg_id = ?", tgID).Count(&count).Error
		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("необходимо указать номер телефона или Telegram ID")
	}
	return count > 0, nil
}
