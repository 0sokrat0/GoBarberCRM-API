package tests

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScheduleRepository_CreateSchedule(t *testing.T) {
	db := setupTestDB(t, &models.Schedule{})
	repo := repositories.NewScheduleRepository(db)

	schedule := &models.Schedule{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "18:00",
	}

	err := repo.CreateSchedule(schedule)
	require.NoError(t, err)
	require.NotZero(t, schedule.ID)
}

func TestScheduleRepository_GetScheduleByID(t *testing.T) {
	db := setupTestDB(t, &models.Schedule{})
	repo := repositories.NewScheduleRepository(db)

	schedule := &models.Schedule{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "18:00",
	}
	err := repo.CreateSchedule(schedule)
	require.NoError(t, err)

	fetchedSchedule, err := repo.GetScheduleByID(schedule.ID)
	require.NoError(t, err)
	assert.Equal(t, schedule.ID, fetchedSchedule.ID)
	assert.Equal(t, schedule.UserID, fetchedSchedule.UserID)
}

func TestScheduleRepository_GetAllSchedules(t *testing.T) {
	db := setupTestDB(t, &models.Schedule{})
	repo := repositories.NewScheduleRepository(db)

	schedules := []models.Schedule{
		{UserID: 1, ScheduleDay: "Monday", StartTime: "09:00", EndTime: "18:00"},
		{UserID: 2, ScheduleDay: "Tuesday", StartTime: "10:00", EndTime: "19:00"},
	}

	for i := range schedules {
		err := repo.CreateSchedule(&schedules[i])
		require.NoError(t, err)
	}

	fetchedSchedules, err := repo.GetAllSchedules()
	require.NoError(t, err)
	assert.Len(t, fetchedSchedules, 2)
}

func TestScheduleRepository_UpdateSchedule(t *testing.T) {
	db := setupTestDB(t, &models.Schedule{})
	repo := repositories.NewScheduleRepository(db)

	schedule := &models.Schedule{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "18:00",
	}
	err := repo.CreateSchedule(schedule)
	require.NoError(t, err)

	schedule.EndTime = "17:00"
	err = repo.UpdateSchedule(schedule)
	require.NoError(t, err)

	updatedSchedule, err := repo.GetScheduleByID(schedule.ID)
	require.NoError(t, err)
	assert.Equal(t, "17:00", updatedSchedule.EndTime)
}

func TestScheduleRepository_DeleteSchedule(t *testing.T) {
	db := setupTestDB(t, &models.Schedule{})
	repo := repositories.NewScheduleRepository(db)

	schedule := &models.Schedule{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "18:00",
	}
	err := repo.CreateSchedule(schedule)
	require.NoError(t, err)

	err = repo.DeleteSchedule(schedule.ID)
	require.NoError(t, err)

	_, err = repo.GetScheduleByID(schedule.ID)
	assert.Error(t, err)
	assert.Equal(t, repositories.ErrScheduleNotFound, err)
}

func TestScheduleRepository_FilterSchedulesByUser(t *testing.T) {
	db := setupTestDB(t, &models.Schedule{})
	repo := repositories.NewScheduleRepository(db)

	schedules := []models.Schedule{
		{UserID: 1, ScheduleDay: "Monday", StartTime: "09:00", EndTime: "18:00"},
		{UserID: 1, ScheduleDay: "Tuesday", StartTime: "09:00", EndTime: "18:00"},
		{UserID: 2, ScheduleDay: "Monday", StartTime: "10:00", EndTime: "19:00"},
	}

	for i := range schedules {
		err := repo.CreateSchedule(&schedules[i])
		require.NoError(t, err)
	}

	userSchedules, err := repo.FilterSchedulesByUser(1)
	require.NoError(t, err)
	assert.Len(t, userSchedules, 2)
}
