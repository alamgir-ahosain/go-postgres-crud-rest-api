package routes

import (
	getusers "github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/api/handlers/get_users"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", getusers.GetUsers).Methods("GET")
}
