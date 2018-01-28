package api

import (
	"fmt"
	"time"
	"unicode/utf8"
)

// UserController is a controller interface
type UserController interface {
	SignIn(*UserSignInRequest) (*UserSignInResponse, error)
	SignOut(*UserSignOutRequest) (*UserSignOutResponse, error)
	ChangePassword(*UserChangePasswordRequest) error
	GetUser(int) (*User, error)
}

// UserSignInRequest is a struct of
// sign in request for UserController
type UserSignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate UserSignInRequest
func (u *UserSignInRequest) Validate() error {
	if len(u.Username) == 0 {
		return fmt.Errorf("Username has required")
	}
	if len(u.Password) == 0 {
		return fmt.Errorf("Password has required")
	}
	if utf8.RuneCountInString(u.Password) < 4 {
		return fmt.Errorf("Password must be at least 8 char")
	}
	return nil
}

// UserSignInResponse is a struct of
// sign in response for UserController
type UserSignInResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

// UserSignOutRequest is a struct of
// sign out request for UserController
type UserSignOutRequest struct {
	Token string `json:"token"`
}

// Validate UserSignOutRequest
func (u *UserSignOutRequest) Validate() error {
	if len(u.Token) == 0 {
		return fmt.Errorf("Token has required")
	}
	return nil
}

// UserSignOutResponse is a struct of
// sign out response for UserController
type UserSignOutResponse struct {
}

// UserChangePasswordRequest is a struct of
// change password for UserController
type UserChangePasswordRequest struct {
	Token       string `json:"token"`
	OldPassword string `jsom:"old_password"`
	NewPassword string `json:"new_password"`
}

// Validate UserChangePasswordRequest
func (u *UserChangePasswordRequest) Validate() error {
	if len(u.OldPassword) == 0 {
		return fmt.Errorf("Password has required")
	}
	if len(u.NewPassword) == 0 {
		return fmt.Errorf("Password has required")
	}
	if utf8.RuneCountInString(u.NewPassword) < 4 {
		return fmt.Errorf("Password must be at least 8 char")
	}
	return nil
}

// Validate get user by id
func (u *User) Validate() error {
	if u.UserID == 0 {
		return fmt.Errorf("user id not correct")
	}
	return nil
}

// User is a struct of get user for UserController
type User struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
