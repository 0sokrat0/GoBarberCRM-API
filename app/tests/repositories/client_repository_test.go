package repositories

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientRepository_CreateClient(t *testing.T) {
	db := setupTestDB(t, &models.Client{})
	repo := repositories.NewClientRepository(db)

	client := &models.Client{
		FirstName:   "John Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "+1234567890",
	}

	err := repo.CreateClient(client)
	require.NoError(t, err)
	require.NotZero(t, client.ID)
}

func TestClientRepository_GetClientByID(t *testing.T) {
	db := setupTestDB(t, &models.Client{})
	repo := repositories.NewClientRepository(db)

	client := &models.Client{
		FirstName:   "John Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "+1234567890",
	}
	err := repo.CreateClient(client)
	require.NoError(t, err)

	fetchedClient, err := repo.GetClientByID(client.ID)
	require.NoError(t, err)
	assert.Equal(t, client.ID, fetchedClient.ID)
	assert.Equal(t, client.FirstName, fetchedClient.FirstName)
}

func TestClientRepository_UpdateClient(t *testing.T) {
	db := setupTestDB(t, &models.Client{})
	repo := repositories.NewClientRepository(db)

	client := &models.Client{
		FirstName:   "John Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "+1234567890",
	}
	err := repo.CreateClient(client)
	require.NoError(t, err)

	client.Email = "new.email@example.com"
	err = repo.UpdateClient(client)
	require.NoError(t, err)

	updatedClient, err := repo.GetClientByID(client.ID)
	require.NoError(t, err)
	assert.Equal(t, "new.email@example.com", updatedClient.Email)
}

func TestClientRepository_DeleteClient(t *testing.T) {
	db := setupTestDB(t, &models.Client{})
	repo := repositories.NewClientRepository(db)

	client := &models.Client{
		FirstName:   "John Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "+1234567890",
	}
	err := repo.CreateClient(client)
	require.NoError(t, err)

	err = repo.DeleteClient(client.ID)
	require.NoError(t, err)

	_, err = repo.GetClientByID(client.ID)
	assert.Error(t, err)
	assert.Equal(t, repositories.ErrClientNotFound, err)
}
