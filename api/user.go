package api

import (
	"time"
)

// UserController is a controller interface
type UserController interface {
	SignIn(*UserSignInRequest) (*UserSignInResponse, error)
	SignOut(*UserSignOutRequest) (*UserSignOutResponse, error)
	ChangePassword(*UserChangePasswordRequest) (*UserChangePasswordResponse, error)
	Get(int) (*User, error)
}

// UserSignInRequest is a struct of
// sign in request for UserController
type UserSignInRequest struct {
	Username string
	Password string
}

// UserSignInResponse is a struct of
// sign in response for UserController
type UserSignInResponse struct {
	Token  string
	UserId int
}

// UserSignOutRequest is a struct of
// sign out request for UserController
type UserSignOutRequest struct {
	Token string
}

// UserSignOutResponse is a struct of
// sign out response for UserController
type UserSignOutResponse struct {
}

// UserChangePasswordRequest is a struct of
// change password for UserController
type UserChangePasswordRequest struct {
	Token       string
	OldPassword string
	NewPassword string
}

type UserChangePasswordResponse struct {
}

// User is a struct of get user for UserController
type User struct {
	UserId    int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CurrentUser is a struct for any request
type CurrentUser struct {
	Id   int
	Name string
}
