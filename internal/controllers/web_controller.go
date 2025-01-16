package controllers

import (
	"net/http"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/templates"
)

type WebController struct {
}

func NewWebController() *WebController {
	return &WebController{}
}

func (wc *WebController) Home(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "home.html", nil)
}

func (wc *WebController) LoginView(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "login.html", nil)
}

func (wc *WebController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the login logic
}

func (wc *WebController) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logout logic
}
