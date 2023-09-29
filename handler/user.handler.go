package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/controller"
	"github.com/Mau005/MyExpenses/models"
	"github.com/gorilla/mux"
)

type Userhandler struct{}

func (ac *Userhandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Leer datos del formulario de inicio de sesión
	email := r.FormValue("email")
	password := r.FormValue("password")

	var userModel controller.UserController
	var api controller.ApiController
	user, token, err := userModel.LoginUser(email, password)
	if err != nil {
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_USER,
			Status:    http.StatusUnauthorized,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	api.SaveSession(&token, w, r)
	json.NewEncoder(w).Encode(user)

}

func (ac *Userhandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Eliminar el token de la sesión
	var api controller.ApiController
	api.SaveSession(nil, w, r)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: "OK",
		Status:  http.StatusOK,
	})
}

func (ac *Userhandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Leer datos del formulario de registro
	email := strings.ToLower(r.FormValue("email"))
	password := r.FormValue("password")

	// Crear nuevo usuario
	var user controller.UserController
	userModel, err := user.CreateUser(email, password)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:         configuration.ERROR_SERVICE_USER,
			Status:        http.StatusNotAcceptable,
			Message:       err.Error(),
			TimeStamp:     time.Now(),
			TransactionId: "1",
			CorrelationId: "1",
		})
		return
	}
	json.NewEncoder(w).Encode(userModel)

}

func (ac *Userhandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	var uc controller.UserController
	user, err := uc.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:         configuration.ERROR_SERVICE_USER,
			Status:        http.StatusNotAcceptable,
			Message:       err.Error(),
			TimeStamp:     time.Now(),
			TransactionId: "1",
			CorrelationId: "1",
		})
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (ac *Userhandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	arguments := mux.Vars(r)

	var uc controller.UserController
	user, err := uc.GetUser(arguments["email"])
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:         configuration.ERROR_SERVICE_USER,
			Status:        http.StatusNotAcceptable,
			Message:       err.Error(),
			TimeStamp:     time.Now(),
			TransactionId: "1",
			CorrelationId: "1",
		})
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (ac *Userhandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	arguments := mux.Vars(r)

	var uc controller.UserController

	err := uc.DelUser(arguments["email"])
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:         configuration.ERROR_SERVICE_USER,
			Status:        http.StatusNotAcceptable,
			Message:       err.Error(),
			TimeStamp:     time.Now(),
			TransactionId: "1",
			CorrelationId: "1",
		})
		return
	}

	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: configuration.DELETE_USER,
	})

}

func (ac *Userhandler) PatchUserHandler(w http.ResponseWriter, r *http.Request) {
	var uc controller.UserController
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	user, err := uc.PutchUser(data)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}

	json.NewEncoder(w).Encode(user)

}
