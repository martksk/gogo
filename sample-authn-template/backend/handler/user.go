package handler

import (
	"net/http"

	"github.com/GDSC-KMUTT/totp-session/service"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) userHandler {
	return userHandler{service: service}
}

func (h userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.service.SignUp("","")
	w.Write([]byte("User create!"))
}

func (h userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.service.SignIn("","")
	w.Write([]byte("User sign in!"))
}