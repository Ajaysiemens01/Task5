package models

// User represents a user in the system following JSON:API specification.
type User struct {
	ID       string `jsonapi:"primary,users"`
	Email    string `jsonapi:"attr,email" validate:"required,email,min=5,max=100"`
	Username string `jsonapi:"attr,username" validate:"required,min=3,max=20"`
	Password string `jsonapi:"attr,password" validate:"required,min=8"`
	Role     string `jsonapi:"attr,role" validate:"required,oneof=intern admin engineer manager operator"`
	Status   string `jsonapi:"attr,status" validate:"required,oneof=suspended inleave active"`
	IsLoggedIn bool `jsonapi:"attr,isloggedin"`
}

// ResponseUser is a struct used to return user data without the password.
type ResponseUser struct {
	ID       string `jsonapi:"primary,users"`
	Email    string `jsonapi:"attr,email"`
	Username string `jsonapi:"attr,username"`
	Role     string `jsonapi:"attr,role"`
	Status   string `jsonapi:"attr,status"`
	IsLoggedIn bool `jsonapi:"attr,isloggedin"`
}

// Method to return User without the password in the response
func (u *User) ToResponse() *ResponseUser {
	return &ResponseUser{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
		Role: u.Role,
		Status: u.Status,
		IsLoggedIn:u.IsLoggedIn,
	}
}
