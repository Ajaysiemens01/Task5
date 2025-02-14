package handlers

import (
    //"encoding/json"
    "net/http"
    "Task5/internal/services"
    "github.com/gorilla/mux"
    "Task5/internal/models"
)

func  usersToResponse(users []*models.User) []*models.ResponseUser {
    var responseUsers []*models.ResponseUser
	for _, user := range users {
		responseUsers = append(responseUsers, user.ToResponse())
	}
	return responseUsers
}
// GetUsersHandler handles the request to get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map to store dynamic filters
	filters := make(map[string]string)

	// Extract query parameters and add to filters map
	if role := r.URL.Query().Get("role"); role != "" {
		filters["role"] = role
	}
	if status := r.URL.Query().Get("status"); status != "" {
		filters["status"] = status
	}

	// Fetch users with applied filters
	users, err := services.GetUsersByFilters(userFileName,filters)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	respondWithJSON(w, http.StatusOK, usersToResponse(users))
}


// GetUserByIDHandler handles the request to get a user by ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    user, err := services.GetUserByID(userFileName, id)
    if err != nil {
        handleError(w, err.Error(), http.StatusNotFound)
        return
    }
    respondWithJSON(w, http.StatusOK, user.ToResponse())
}

