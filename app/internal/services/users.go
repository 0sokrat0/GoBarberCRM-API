package services

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id int, input *models.User) error
	DeleteUser(id int) error
	AuthenticateUser(identifier, password string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(user *models.User) error {
	// Проверка обязательных полей
	if user.Username == "" || user.PasswordHash == "" || user.Role == "" {
		return errors.New("обязательные поля: Username, Password и Role")
	}

	// Проверка уникальности Username и Email
	if _, err := s.repo.GetUserByUsername(user.Username); err == nil {
		return errors.New("пользователь с таким Username уже существует")
	}

	if user.Email != "" {
		if _, err := s.repo.GetUserByEmail(user.Email); err == nil {
			return errors.New("пользователь с таким Email уже существует")
		}
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("не удалось хешировать пароль")
	}
	user.PasswordHash = string(hashedPassword)

	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) UpdateUser(id int, input *models.User) error {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	// Обновление полей
	if input.Username != "" && input.Username != user.Username {
		// Проверка уникальности нового Username
		if existingUser, err := s.repo.GetUserByUsername(input.Username); err == nil && existingUser.ID != id {
			return errors.New("пользователь с таким Username уже существует")
		}
		user.Username = input.Username
	}

	if input.Email != "" && input.Email != user.Email {
		// Проверка уникальности нового Email
		if existingUser, err := s.repo.GetUserByEmail(input.Email); err == nil && existingUser.ID != id {
			return errors.New("пользователь с таким Email уже существует")
		}
		user.Email = input.Email
	}

	if input.PasswordHash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("не удалось хешировать пароль")
		}
		user.PasswordHash = string(hashedPassword)
	}

	if input.Role != "" {
		user.Role = input.Role
	}

	if input.PhoneNumber != "" {
		user.PhoneNumber = input.PhoneNumber
	}

	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *userService) AuthenticateUser(identifier, password string) (*models.User, error) {
	var user *models.User
	var err error

	if user, err = s.repo.GetUserByEmail(identifier); err != nil {
		if err != repositories.ErrUserNotFound {
			return nil, err
		}
		// Если не найден по email, пробуем по username
		if user, err = s.repo.GetUserByUsername(identifier); err != nil {
			return nil, errors.New("неверные учетные данные")
		}
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("неверные учетные данные")
	}

	return user, nil
}
