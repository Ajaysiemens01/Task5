package services

import (
	"errors"
	"Task5/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Check if a user exists in the database
func userExists(usersDB []*models.User, username string) bool {
    for _, user := range usersDB {
        if user.Username == username {
            return true
        }
    }
    return false
}
// Check if a user exists in the database
func emailExists(usersDB []*models.User, email string) bool {
    for _, user := range usersDB {
        if user.Email == email {
            return true
        }
    }
    return false
}
//Delete a user
func delete(usersDB *[]*models.User, reqUser *models.User) []*models.User {
    // Create a new slice to hold users that aren't being deleted
    var updatedUsers []*models.User

    for _, user := range *usersDB {
        if user.ID != reqUser.ID {
            updatedUsers = append(updatedUsers, user)
        }
    }

    return updatedUsers
}


// Hash the user's password
func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}


// Find a user by username
func findUserByUserName(usersDB []*models.User, username string) (*models.User, error) {
    for _, user := range usersDB {
        if user.Username == username {
            return user, nil
        }
    }
    return nil, errors.New("invalid username")
}

// Find a user by username
func findUserByID(usersDB []*models.User, id string) (*models.User, error) {
    for _, user := range usersDB {
        if user.ID == id {
            return user, nil
        }
    }
    return nil, errors.New("user not found")
}
// Check if the provided password matches the hashed password
func checkPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}