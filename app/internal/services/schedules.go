package services

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
)

type ScheduleService interface {
	CreateSchedule(schedule *models.Schedule) error
	GetScheduleByID(id int) (*models.Schedule, error)
	GetAllSchedules() ([]models.Schedule, error)
	UpdateSchedule(id int, input *models.Schedule) error
	DeleteSchedule(id int) error
	FilterSchedulesByUser(userID int) ([]models.Schedule, error)
}

type scheduleService struct {
	repo repositories.ScheduleRepository
}

func NewScheduleService(repo repositories.ScheduleRepository) ScheduleService {
	return &scheduleService{
		repo: repo,
	}
}

func (s *scheduleService) CreateSchedule(schedule *models.Schedule) error {
	// Дополнительная бизнес-логика перед созданием расписания (если требуется)
	return s.repo.CreateSchedule(schedule)
}

func (s *scheduleService) GetScheduleByID(id int) (*models.Schedule, error) {
	return s.repo.GetScheduleByID(id)
}

func (s *scheduleService) GetAllSchedules() ([]models.Schedule, error) {
	return s.repo.GetAllSchedules()
}

func (s *scheduleService) UpdateSchedule(id int, input *models.Schedule) error {
	schedule, err := s.repo.GetScheduleByID(id)
	if err != nil {
		return err
	}

	// Обновляем поля
	schedule.ScheduleDay = input.ScheduleDay
	schedule.StartTime = input.StartTime
	schedule.EndTime = input.EndTime

	return s.repo.UpdateSchedule(schedule)
}

func (s *scheduleService) DeleteSchedule(id int) error {
	return s.repo.DeleteSchedule(id)
}

func (s *scheduleService) FilterSchedulesByUser(userID int) ([]models.Schedule, error) {
	return s.repo.FilterSchedulesByUser(userID)
}
