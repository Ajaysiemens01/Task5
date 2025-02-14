package services

import (
    "errors"
    "Task5/internal/models"
    "Task5/pkg/utils"
	"github.com/google/uuid"
)

 
// Register a new user
func Register(usersFile string, user *models.User) (*models.User, error) {
	//password validation
	if len(user.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}
	if len(user.Username) < 3 {
		return nil, errors.New("username must be at least 3 characters long")
	}

	usersDB, err := utils.ReadUsers(usersFile)
	if err != nil {
		return nil, err
	}
	if userExists(usersDB, user.Username) {
		return nil, errors.New("username already exists")
	}
	if emailExists(usersDB, user.Email) {
		return nil, errors.New("email already exists")
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
    user.Password = hashedPassword
	user.ID =  uuid.New().String()
	user.IsLoggedIn = true
	usersDB = append(usersDB, user)

	if err := utils.WriteUsers(usersFile, usersDB); err != nil {
		return nil, err
	}

	return user, nil
}

// Login a user 
func Login(usersFile string, username, password string) (*models.User, error) {
	usersDB, err := utils.ReadUsers(usersFile)
	if err != nil {
		return nil, err
	}

	user, err := findUserByUserName(usersDB, username)
	if err != nil {
		return nil, err
	}

	//password verification
	if err := checkPassword(user.Password, password); err != nil {
		return nil, errors.New("invalid password")
	}
	user.IsLoggedIn = true
	usersDB = append(usersDB, user)

	if err := utils.WriteUsers(usersFile, usersDB); err != nil {
		return nil, err
	}
	return user, nil
}


// Update updates an existing user's information.
func Update(usersFile string, user *models.User, id string) (*models.User, error) {
    usersDB, err := utils.ReadUsers(usersFile)
    if err != nil {
        return nil, err
    }
    reqUser, err := findUserByID(usersDB, id)
    if err != nil {
        return nil, err 
    }
    if reqUser.IsLoggedIn {
        
		if user.Email != "" && user.Email != reqUser.Email {
			reqUser.Email = user.Email
		}
		if user.Username != "" && user.Username != reqUser.Username {
			reqUser.Username = user.Username
		}
		if user.Password != "" && user.Password != reqUser.Password {
			reqUser.Password = user.Password
		}
		if user.Role != "" && user.Role != reqUser.Role {
			reqUser.Role = user.Role
		}
		if user.Status != "" && user.Status != reqUser.Status {
			reqUser.Status = user.Status
		}

		if err := utils.WriteUsers(usersFile, usersDB); err != nil {
			return nil, err
		}
        return reqUser, nil 
    } 
    return nil, errors.New("user is not logged in")
}



// Delete deletes an existing user's information.
func Delete(usersFile string, id string) (*models.User, error) {
    usersDB, err := utils.ReadUsers(usersFile)
    if err != nil {
        return nil, err
    }
    reqUser, err := findUserByID(usersDB, id)
    if err != nil {
        return nil, err 
    }
    if reqUser.IsLoggedIn {
       usersDB = delete(&usersDB,reqUser)
		if err := utils.WriteUsers(usersFile, usersDB); err != nil {
			return nil, err
		}
        return reqUser, nil 
    } 
    return nil, errors.New("user is not logged in")
}
