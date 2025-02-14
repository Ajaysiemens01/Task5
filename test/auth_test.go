package test

import (
    "testing"
    "Task5/internal/services"
    "github.com/stretchr/testify/assert"
    "Task5/internal/models"

)

// TestRegister tests the Register function for creating a new user.
func TestRegister(t *testing.T) {
	
   // Initialize the user object
   user := &models.User{
	Username: "Alex",
	Email:    "alex@siemens.com",
	Password: "Alex@1234",
	Role:     "operator",
	Status:   "suspended",
}
    requser, err := services.Register("test.json", user)

    // Assert that there is no error and the username is as expected.
    assert.Nil(t, err)
    assert.Equal(t, "Alex", requser.Username)
}

func TestRegisterExistingUsername(t *testing.T) {
    // Assuming the first user has already been registered
    user := &models.User{
        Username: "John",
        Email:    "john@yahoo.com",
        Password: "John@1111",
        Role:     "engineer",
        Status:   "active",
    }

    _, err := services.Register("test.json", user)

    assert.NotNil(t, err)
}
func TestRegisterExistingEmail(t *testing.T) {
      
    user:= &models.User{
        Username: "JohnDoe", 
        Email:    "john@siemens.com", // Same email
        Password: "John@2222",
        Role:     "admin",
        Status:   "active",
    }

    _, err := services.Register("test.json", user)
    assert.NotNil(t, err) // Assert an error is returned because the email already exists
}

// TestRegisterEmptyUser tests the Register function for an empty user.
func TestRegisterEmptyUser(t *testing.T) {
    user := &models.User{
        Username: "",
        Email:    "",
        Password: "",
        Role:     "operator", // assuming a default role is fine
        Status:   "suspended",
    }

    // Try registering an empty user
    _, err := services.Register("test.json", user)
    assert.NotNil(t, err) // Assert that an error is returned due to missing required fields
}


// TestLogin tests the Login function with valid credentials.
func TestLogin(t *testing.T) {
    username := "Ajay"
    password := "Ajay@1111"

    // Assuming this user is already registered, try logging in
    user, err := services.Login("test.json", username, password)
    assert.Nil(t, err) // No error expected
    assert.Equal(t, "Ajay", user.Username) // Ensure the username matches
}

// TestLoginInvalidCredentials tests the Login function with invalid credentials.
func TestLoginInvalidCredentials(t *testing.T) {
    username := "John"
    password := "wrongpassword"

    // Try logging in with invalid credentials
    _, err := services.Login("test.json", username, password)
    assert.NotNil(t, err) // Assert that an error is returned
}
