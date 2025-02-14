package test

import (
    "testing"
    "user-api/internal/services"
    "github.com/stretchr/testify/assert"
)

// TestGetAllUsers tests the GetAllUsers function to ensure it retrieves all users.
func TestGetUsersByFilters(t *testing.T) {
    filters := make(map[string]string)
    filters["role"] = "operator"
    filters["status"] = "suspended"
    users, err := services.GetUsersByFilters("test.json",filters)

    // Assert
    assert.Nil(t, err)
    assert.NotEmpty(t, users)
}

// TestGetUsersByRole tests the GetUsersByFilters function with filter of role
func TestGetUsersByRole(t *testing.T) {
    filters := make(map[string]string)
    filters["role"] = "operator"

    users, err := services.GetUsersByFilters("test.json",filters)

    // Assert
    assert.Nil(t, err)
    assert.NotEmpty(t, users)
}

// TestGetUsersByStatus tests the GetUsersByFilters function with filter of status
func TestGetUsersByStatus(t *testing.T) {
    filters := make(map[string]string)
    filters["status"] = "suspended"
    users, err := services.GetUsersByFilters("test.json",filters)

    // Assert
    assert.Nil(t, err)
    assert.NotEmpty(t, users)
}

// TestGetUsersByFiltersUnexisted tests the GetUsersByFilters function with unexisted filters
func TestGetUsersByFiltersUnexisted(t *testing.T) {
    filters := make(map[string]string)
    filters["role"] = "developer"
    filters["status"] = "suspended"
    users, err := services.GetUsersByFilters("test.json",filters)

    // Assert
    assert.NotNil(t, err)
    assert.Empty(t, users)
}

// TestGetAllUsersEmpty tests the GetUsersByFilters function with empty users
func TestGetAllUsersEmpty(t *testing.T) {
    filters := make(map[string]string)
    users, err := services.GetUsersByFilters("test.json",filters)

    // Assert 
    assert.Nil(t, err)
    assert.Empty(t, users)
}
// TestGetUserByID tests the GetUserByID function with a valid user ID
func TestGetUserByID(t *testing.T) {
    userID := "b4c2f113-2d7e-4813-92f1-1e335652e56d"

    user, err := services.GetUserByID("test.json", userID)

    // Assert 
    assert.Nil(t, err)
    assert.Equal(t, userID, user.ID)
}

// TestGetUserByIDNotFound tests the GetUserByID function with a non-existent user ID.
func TestGetUserByIDNotFound(t *testing.T) {
    userID := "nonexistent"
    _, err := services.GetUserByID("test.json", userID)

    // Assert 
    assert.NotNil(t, err)
    assert.Equal(t, "user not found", err.Error())
}