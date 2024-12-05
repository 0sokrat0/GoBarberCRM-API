package repositories

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"

	"gorm.io/gorm"
)

var (
	ErrScheduleNotFound = errors.New("расписание не найдено")
)

type ScheduleRepository interface {
	CreateSchedule(schedule *models.Schedule) error
	GetScheduleByID(id int) (*models.Schedule, error)
	GetAllSchedules() ([]models.Schedule, error)
	UpdateSchedule(schedule *models.Schedule) error
	DeleteSchedule(id int) error
	FilterSchedulesByUser(userID int) ([]models.Schedule, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func (r *scheduleRepository) CreateSchedule(schedule *models.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) GetScheduleByID(id int) (*models.Schedule, error) {
	var schedule models.Schedule
	if err := r.db.First(&schedule, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrScheduleNotFound
		}
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) GetAllSchedules() ([]models.Schedule, error) {
	var schedules []models.Schedule
	if err := r.db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) UpdateSchedule(schedule *models.Schedule) error {
	return r.db.Save(schedule).Error
}

func (r *scheduleRepository) DeleteSchedule(id int) error {
	if err := r.db.Delete(&models.Schedule{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepository) FilterSchedulesByUser(userID int) ([]models.Schedule, error) {
	var schedules []models.Schedule
	if err := r.db.Where("user_id = ?", userID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
