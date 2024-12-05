package repositories

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"gorm.io/gorm"
)

var (
	ErrBookingNotFound  = errors.New("бронирование не найдено")
	ErrTimeSlotOccupied = errors.New("временной слот уже занят")
)

type BookingRepository interface {
	CreateBooking(booking *models.Bookings) error
	GetBookingByID(id int) (*models.Bookings, error)
	GetAllBookings() ([]models.Bookings, error)
	UpdateBooking(booking *models.Bookings) error
	DeleteBooking(id int) error
	IsTimeSlotOccupied(userID int, bookingTime string) (bool, error)
	GetBookingsByClientID(clientID int) ([]models.Bookings, error)
	GetBookingsByServiceID(serviceID int) ([]models.Bookings, error)
	GetBookingsByUserID(userID int) ([]models.Bookings, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

func (r *bookingRepository) CreateBooking(booking *models.Bookings) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingByID(id int) (*models.Bookings, error) {
	var booking models.Bookings
	if err := r.db.Preload("Client").Preload("Service").Preload("User").First(&booking, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookingNotFound
		}
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) GetAllBookings() ([]models.Bookings, error) {
	var bookings []models.Bookings
	if err := r.db.Preload("Client").Preload("Service").Preload("User").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) UpdateBooking(booking *models.Bookings) error {
	return r.db.Save(booking).Error
}

func (r *bookingRepository) DeleteBooking(id int) error {
	if err := r.db.Delete(&models.Bookings{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) IsTimeSlotOccupied(userID int, bookingTime string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Bookings{}).
		Where("user_id = ? AND booking_time = ?", userID, bookingTime).
		Count(&count).Error
	return count > 0, err
}

func (r *bookingRepository) GetBookingsByClientID(clientID int) ([]models.Bookings, error) {
	var bookings []models.Bookings
	if err := r.db.Where("client_id = ?", clientID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) GetBookingsByServiceID(serviceID int) ([]models.Bookings, error) {
	var bookings []models.Bookings
	if err := r.db.Where("service_id = ?", serviceID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) GetBookingsByUserID(userID int) ([]models.Bookings, error) {
	var bookings []models.Bookings
	if err := r.db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
