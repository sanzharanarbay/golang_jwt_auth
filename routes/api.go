package routes

import (
	usersController "jwt_auth_golang/controllers/users"
	authController "jwt_auth_golang/controllers/auth"
	middleware "jwt_auth_golang/middleware"
	"github.com/gorilla/mux"
)

func ApiRoutes(prefix string, r *mux.Router) {

	s := r.PathPrefix(prefix).Subrouter()

	s.HandleFunc("/login", authController.Login).Methods("POST")
	s.HandleFunc("/register", usersController.CreateUser).Methods("POST")
	s.HandleFunc("/users", middleware.ValidateMiddleware(usersController.GetUsers)).Methods("GET")
	s.HandleFunc("/users/{id}", middleware.ValidateMiddleware(usersController.GetUser)).Methods("GET")
	s.HandleFunc("/users", middleware.ValidateMiddleware(usersController.CreateUser)).Methods("POST")
	s.HandleFunc("/users/{id:[0-9]+}", middleware.ValidateMiddleware(usersController.GetUser)).Methods("GET")
	s.HandleFunc("/users/{id:[0-9]+}", middleware.ValidateMiddleware(usersController.UpdateUser)).Methods("PUT")
	s.HandleFunc("/users/{id}", middleware.ValidateMiddleware(usersController.DeleteUser)).Methods("DELETE")
}
