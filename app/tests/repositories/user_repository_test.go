package repositories

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db := setupTestDB(t, &models.User{})
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Role:         "admin",
		Email:        "test@example.com",
	}

	err := repo.CreateUser(user)
	require.NoError(t, err)
	require.NotZero(t, user.ID)
}

func TestUserRepository_GetUserByID(t *testing.T) {
	db := setupTestDB(t, &models.User{})
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Role:         "admin",
		Email:        "test@example.com",
	}
	err := repo.CreateUser(user)
	require.NoError(t, err)

	fetchedUser, err := repo.GetUserByID(user.ID)
	require.NoError(t, err)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.Username, fetchedUser.Username)
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	db := setupTestDB(t, &models.User{})
	repo := repositories.NewUserRepository(db)

	users := []models.User{
		{Username: "user1", PasswordHash: "pass1", Role: "user", Email: "user1@example.com"},
		{Username: "user2", PasswordHash: "pass2", Role: "user", Email: "user2@example.com"},
	}

	for i := range users {
		err := repo.CreateUser(&users[i])
		require.NoError(t, err)
	}

	fetchedUsers, err := repo.GetAllUsers()
	require.NoError(t, err)
	assert.Len(t, fetchedUsers, 2)
}

func TestUserRepository_UpdateUser(t *testing.T) {
	db := setupTestDB(t, &models.User{})
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Role:         "admin",
		Email:        "test@example.com",
	}
	err := repo.CreateUser(user)
	require.NoError(t, err)

	user.Email = "newemail@example.com"
	err = repo.UpdateUser(user)
	require.NoError(t, err)

	updatedUser, err := repo.GetUserByID(user.ID)
	require.NoError(t, err)
	assert.Equal(t, "newemail@example.com", updatedUser.Email)
}

func TestUserRepository_DeleteUser(t *testing.T) {
	db := setupTestDB(t, &models.User{})
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Role:         "admin",
		Email:        "test@example.com",
	}
	err := repo.CreateUser(user)
	require.NoError(t, err)

	err = repo.DeleteUser(user.ID)
	require.NoError(t, err)

	_, err = repo.GetUserByID(user.ID)
	assert.Error(t, err)
	assert.Equal(t, repositories.ErrUserNotFound, err)
}
