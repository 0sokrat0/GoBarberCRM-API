package services

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
)

var (
	ErrBookingNotFound  = errors.New("бронирование не найдено")
	ErrTimeSlotOccupied = errors.New("временной слот уже занят")
)

type BookingService interface {
	CreateBooking(booking *models.Bookings) error
	GetBookingByID(id int) (*models.Bookings, error)
	GetAllBookings() ([]models.Bookings, error)
	UpdateBooking(id int, input *models.Bookings) error
	DeleteBooking(id int) error
	CheckAvailability(userID int, bookingTime string) (bool, error)
	GetBookingsByClientID(clientID int) ([]models.Bookings, error)
	GetBookingsByServiceID(serviceID int) ([]models.Bookings, error)
	GetBookingsByUserID(userID int) ([]models.Bookings, error)
}

type bookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) BookingService {
	return &bookingService{
		repo: repo,
	}
}

func (s *bookingService) CreateBooking(booking *models.Bookings) error {
	occupied, err := s.repo.IsTimeSlotOccupied(booking.UserID, booking.BookingTime.String())
	if err != nil {
		return err
	}
	if occupied {
		return ErrTimeSlotOccupied
	}
	return s.repo.CreateBooking(booking)
}

func (s *bookingService) GetBookingByID(id int) (*models.Bookings, error) {
	return s.repo.GetBookingByID(id)
}

func (s *bookingService) GetAllBookings() ([]models.Bookings, error) {
	return s.repo.GetAllBookings()
}

func (s *bookingService) UpdateBooking(id int, input *models.Bookings) error {
	booking, err := s.repo.GetBookingByID(id)
	if err != nil {
		return err
	}

	// Обновляем поля
	booking.UserID = input.UserID
	booking.ClientID = input.ClientID
	booking.ServiceID = input.ServiceID
	booking.BookingTime = input.BookingTime
	// Другие поля...

	return s.repo.UpdateBooking(booking)
}

func (s *bookingService) DeleteBooking(id int) error {
	return s.repo.DeleteBooking(id)
}

func (s *bookingService) CheckAvailability(userID int, bookingTime string) (bool, error) {
	occupied, err := s.repo.IsTimeSlotOccupied(userID, bookingTime)
	if err != nil {
		return false, err
	}
	return !occupied, nil
}

func (s *bookingService) GetBookingsByClientID(clientID int) ([]models.Bookings, error) {
	return s.repo.GetBookingsByClientID(clientID)
}

func (s *bookingService) GetBookingsByServiceID(serviceID int) ([]models.Bookings, error) {
	return s.repo.GetBookingsByServiceID(serviceID)
}

func (s *bookingService) GetBookingsByUserID(userID int) ([]models.Bookings, error) {
	return s.repo.GetBookingsByUserID(userID)
}
