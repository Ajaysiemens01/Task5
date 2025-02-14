package services

import (
    "strings"
    "errors"
    "Task5/internal/models"
    "Task5/pkg/utils"
)

// GetUsersByFilters filters users based on the provided filters (role, status, or both).
func GetUsersByFilters(usersFile string, filters map[string]string) ([]*models.User, error) {
	// Read users from JSON file
	usersDB, err := utils.ReadUsers(usersFile)
	if err != nil {
		return nil, err
	}

	// Filter users based on the provided filters
	var filteredUsers []*models.User
	for _, user := range usersDB {
		match := true
		// Check role filter if exists
		if role, ok := filters["role"]; ok &&  !strings.EqualFold(user.Role,role) {
			match = false
		}

		// Check status filter if exists
		if status, ok := filters["status"]; ok && !strings.EqualFold(user.Status,status) {
			match = false
		}

		// If the user matches all filters, add them to the result
		if match {
			filteredUsers = append(filteredUsers, user)
		}
	}

	// If no users match the filters, return an error (optional)
	if len(filteredUsers) == 0 {
		return nil, errors.New("no users found matching the filters")
	}

	// Return the filtered list of users
	return filteredUsers, nil
}

// Get a user by ID from DB
func GetUserByID(usersFile string, id string) (*models.User, error) {
    usersDB, err := utils.ReadUsers(usersFile)
    if usersDB == nil {
        return  nil, nil
    }
    if err != nil {
        panic(err)
    }
    user, err := findUserByID(usersDB, id)
    if err != nil {
        return user, err
    }
    
    return user,nil
}
