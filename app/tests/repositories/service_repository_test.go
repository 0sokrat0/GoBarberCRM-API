package repositories

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServiceRepository_CreateService(t *testing.T) {
	db := setupTestDB(t, &models.Service{})
	repo := repositories.NewServiceRepository(db)

	service := &models.Service{
		Name:        "Test Service",
		Description: "Description",
		Price:       100.0,
		Duration:    60,
		IsActive:    true,
	}

	err := repo.CreateService(service)
	require.NoError(t, err)
	require.NotZero(t, service.ID)
}

func TestServiceRepository_GetServiceByID(t *testing.T) {
	db := setupTestDB(t, &models.Service{})
	repo := repositories.NewServiceRepository(db)

	service := &models.Service{
		Name:        "Test Service",
		Description: "Description",
		Price:       100.0,
		Duration:    60,
		IsActive:    true,
	}
	err := repo.CreateService(service)
	require.NoError(t, err)

	fetchedService, err := repo.GetServiceByID(service.ID)
	require.NoError(t, err)
	assert.Equal(t, service.ID, fetchedService.ID)
	assert.Equal(t, service.Name, fetchedService.Name)
}

func TestServiceRepository_GetAllServices(t *testing.T) {
	db := setupTestDB(t, &models.Service{})
	repo := repositories.NewServiceRepository(db)

	services := []models.Service{
		{Name: "Service 1", Price: 50.0, Duration: 30, IsActive: true},
		{Name: "Service 2", Price: 80.0, Duration: 45, IsActive: true},
	}

	for i := range services {
		err := repo.CreateService(&services[i])
		require.NoError(t, err)
	}

	fetchedServices, err := repo.GetAllServices()
	require.NoError(t, err)
	assert.Len(t, fetchedServices, 2)
}

func TestServiceRepository_UpdateService(t *testing.T) {
	db := setupTestDB(t, &models.Service{})
	repo := repositories.NewServiceRepository(db)

	service := &models.Service{
		Name:        "Old Service",
		Description: "Old Description",
		Price:       100.0,
		Duration:    60,
		IsActive:    true,
	}
	err := repo.CreateService(service)
	require.NoError(t, err)

	service.Name = "Updated Service"
	err = repo.UpdateService(service)
	require.NoError(t, err)

	updatedService, err := repo.GetServiceByID(service.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated Service", updatedService.Name)
}

func TestServiceRepository_DeleteService(t *testing.T) {
	db := setupTestDB(t, &models.Service{})
	repo := repositories.NewServiceRepository(db)

	service := &models.Service{
		Name:     "Test Service",
		Price:    100.0,
		Duration: 60,
		IsActive: true,
	}
	err := repo.CreateService(service)
	require.NoError(t, err)

	err = repo.DeleteService(service.ID)
	require.NoError(t, err)

	_, err = repo.GetServiceByID(service.ID)
	assert.Error(t, err)
	assert.Equal(t, repositories.ErrServiceNotFound, err)
}

func TestServiceRepository_DeactivateService(t *testing.T) {
	db := setupTestDB(t, &models.Service{})
	repo := repositories.NewServiceRepository(db)

	service := &models.Service{
		Name:     "Test Service",
		Price:    100.0,
		Duration: 60,
		IsActive: true,
	}
	err := repo.CreateService(service)
	require.NoError(t, err)

	err = repo.DeactivateService(service.ID)
	require.NoError(t, err)

	updatedService, err := repo.GetServiceByID(service.ID)
	require.NoError(t, err)
	assert.False(t, updatedService.IsActive)
}
