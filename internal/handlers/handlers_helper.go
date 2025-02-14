package handlers

import (
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/google/jsonapi"
	"errors"
	"v/internal/models"
)

// respondWithJSON writes the response in JSONAPI format
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)
	// Convert the struct or slice to JSONAPI format and write to response
	if err := jsonapi.MarshalPayload(w, data); err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleValidationError formats and sends validation errors in JSON:API format
func handleValidationError(w http.ResponseWriter, messages []string, statusCode int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)
	// Create a slice of jsonapi.ErrorObject
	var errors []*jsonapi.ErrorObject

	for _, msg := range messages {
		errors = append(errors, &jsonapi.ErrorObject{
			Status: fmt.Sprintf("%d", statusCode), 
			Title:  http.StatusText(statusCode),
			Detail: msg,
		})
	}
	// Encode and send JSONAPI error response
	if err := jsonapi.MarshalErrors(w, errors); err != nil {
		http.Error(w, "Error encoding error response", http.StatusInternalServerError)
	}
}

// handleError is a helper function to handle errors
func handleError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)
	var errors []*jsonapi.ErrorObject

		errors = append(errors, &jsonapi.ErrorObject{
			Status: fmt.Sprintf("%d", statusCode), 
			Title:  http.StatusText(statusCode),
			Detail: message,
		})
	if err := jsonapi.MarshalErrors(w, errors); err != nil {
		http.Error(w, "Error encoding error response", http.StatusInternalServerError)
	}
}

// decodeRequestBody is a helper to decode the request body
func decodeRequestBody(r *http.Request, user *models.User) error {
	// ensure Content-Type is JSONAPI
	if r.Header.Get("Content-Type") != "application/vnd.api+json" {
		return errors.New("invalid Content-Type, expected application/vnd.api+json")
	}
	// decode JSONAPI request
	if err := jsonapi.UnmarshalPayload(r.Body, user); err != nil {
		return err
	}

	return nil
}

// BundleValidationErrors maps validation errors to custom messages
func BundleValidationErrors(err error) []string {
	var errorMessages []string
	fmt.Println(err)

	// Loop through validation errors
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Field() {
		case "Username":
			if e.Tag() == "required" {
				errorMessages = append(errorMessages, "Username is required")
			}
			if e.Tag() == "min" {
				errorMessages = append(errorMessages, "Username must be at least 3 characters long")
			}
			if e.Tag() == "max" {
				errorMessages = append(errorMessages, "Username must be at most 20 characters long")
			}
		case "Password":
			if e.Tag() == "required" {
				errorMessages = append(errorMessages, "Password is required")
			}
			if e.Tag() == "min" {
				errorMessages = append(errorMessages, "Password must be at least 8 characters long")
			}
		case "Email":
			if e.Tag() == "required" {
				errorMessages = append(errorMessages, "Email is required")
			}
			if e.Tag() == "email" {
				errorMessages = append(errorMessages, "Invalid email format")
			}
			if e.Tag() == "min" {
				errorMessages = append(errorMessages, "Email must be at least 5 characters long")
			}
			if e.Tag() == "max" {
				errorMessages = append(errorMessages, "Email must be at most 100 characters long")
			}
		case "Role":
			if e.Tag() == "required" {
				errorMessages = append(errorMessages, "Role is required")
			}
			if e.Tag() == "oneof" {
				errorMessages = append(errorMessages, "Role must be one of the following: intern, admin, engineer, manager, operator")
			}
		case "Status":
			if e.Tag() == "oneof" {
				errorMessages = append(errorMessages, "Role must be one of the following: suspended, inleave, active")
			}
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", e.Field()))
		}
	}

	return errorMessages
}
