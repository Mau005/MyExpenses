package router

import (
	"github.com/Mau005/MyExpenses/handler"
	"github.com/Mau005/MyExpenses/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	var userHandler handler.Userhandler
	router := mux.NewRouter()
	router.Use(middleware.CommonMiddleware)
	router.HandleFunc("/signup", userHandler.SignupHandler).Methods("POST")
	router.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", userHandler.LogoutHandler).Methods("GET")

	security := router.PathPrefix("/api/v1").Subrouter()
	security.Use(middleware.CommonMiddleware)
	security.Use(middleware.SessionMiddleware)
	security.HandleFunc("/users", userHandler.UsersHandler).Methods("GET")
	security.HandleFunc("/user/{email}", userHandler.GetUserHandler).Methods("GET")
	security.HandleFunc("/user/{email}", userHandler.DeleteUserHandler).Methods("DELETE")
	security.HandleFunc("/user", userHandler.PatchUserHandler).Methods("PATCH")

	var categoryHandler handler.Categoryhandler
	security.HandleFunc("/categorys", categoryHandler.GetAllCategoryHandler).Methods("GET")
	security.HandleFunc("/category", categoryHandler.CreateCategoryHandler).Methods("POST")
	security.HandleFunc("/category/{id}", categoryHandler.GetCategoryHandler).Methods("GET")
	security.HandleFunc("/category/{id}", categoryHandler.DelCategoryHandler).Methods("DELETE")
	security.HandleFunc("/category", categoryHandler.PatchCategoryHandler).Methods("PATCH")

	return router

}
