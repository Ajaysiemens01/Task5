package handlers

import (
	"net/http"
	"Task5/internal/services"
	"github.com/go-playground/validator/v10"
	"Task5/internal/models"
	"github.com/gorilla/mux"
)
// validator instance
var validate = validator.New()
const userFileName = "user.json"


// Home page handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the User API"))
}

// Register handler
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User
	if err := decodeRequestBody(r, &req); err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate input with custom error handling
	if err := validate.Struct(req); err != nil {
		// map the validation error to custom error messages
		customErrors := BundleValidationErrors(err)
		handleValidationError(w, customErrors, http.StatusBadRequest)
		return
	}

	user, err := services.Register(userFileName, &req)
	if err != nil {
		handleError(w, err.Error(), http.StatusConflict)
		return
	}
	respondWithJSON(w, http.StatusCreated, user.ToResponse())
}


// Login handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User
	if err := decodeRequestBody(r, &req); err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if req.Username == "" || req.Password == "" {
		handleError(w, "invalid credentials",  http.StatusBadRequest)
		return
	}

	user, err := services.Login(userFileName,req.Username,req.Password)
	if err != nil {
		handleError(w, err.Error(), http.StatusUnauthorized)
		return
	}
	respondWithJSON(w, http.StatusOK, user.ToResponse())
}

// Update handler
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User
	vars := mux.Vars(r)
    id := vars["id"]
	if err := decodeRequestBody(r, &req); err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.Update(userFileName, &req, id)
	if err != nil {
		handleError(w, err.Error(),http.StatusUnauthorized)
		return
	}
	respondWithJSON(w, http.StatusOK, user.ToResponse())
}

// Update handler
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User
	vars := mux.Vars(r)
    id := vars["id"]
	if err := decodeRequestBody(r, &req); err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.Delete(userFileName, id)
	if err != nil {
		handleError(w, err.Error(),http.StatusUnauthorized)
		return
	}
	respondWithJSON(w, http.StatusOK, user.ToResponse())
}




