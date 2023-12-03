package user

import (
	"net/http"
)

type JsonFormatUpdateUserRequest struct {
	Email string
	Name  string
}

type JsonFormatCreateUserRequest struct {
	Email    string
	Password []byte
	Name     string
}

type Controller interface {
	GetUsers(http.ResponseWriter, *http.Request)
	GetUsersByRole(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	PutUser(http.ResponseWriter, *http.Request)
	PostUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}
