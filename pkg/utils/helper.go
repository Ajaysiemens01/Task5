package utils

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"Task5/internal/models"

	"github.com/google/jsonapi"
)

// ReadUsers reads users from a JSON:API formatted file.
func ReadUsers(filename string) ([]*models.User, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []*models.User{}, nil // If file doesn't exist, return empty slice
		}
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Decode JSON:API request into a slice of interfaces
	rawUsers, err := jsonapi.UnmarshalManyPayload(file, reflect.TypeOf([]models.User{}))
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON:API format from file %s: %w", filename, err)
	}

	// Convert []interface{} to []*models.User
	users := make([]*models.User, len(rawUsers))
	for i, rawUser := range rawUsers {
		if user, ok := rawUser.(*models.User); ok {
			users[i] = user
		} else {
			return nil, fmt.Errorf("failed to convert user at index %d to *models.User", i)
		}
	}

	return users, nil
}

// WriteUsers writes users in JSONAPI format to a file.
func WriteUsers(filename string, users []*models.User) error {
	// Create or overwrite the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Create a buffer to store JSONAPI output
	var buf bytes.Buffer
	if err := jsonapi.MarshalPayload(&buf, users); err != nil {
		return fmt.Errorf("failed to marshal JSON:API data: %w", err)
	}

	// Write buffer contents to the file
	if _, err := file.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}

	return nil
}
