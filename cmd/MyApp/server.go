package myapp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/api/routes"
	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/db"
	"github.com/gorilla/mux"
)

func Serve() {

	db.Connect() //connect with Database
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	fmt.Println("server running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting the server", err)

	}

}
