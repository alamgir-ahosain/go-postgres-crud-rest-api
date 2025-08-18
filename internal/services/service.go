package services

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/models"
)



func HandleHTTPError(w http.ResponseWriter, err error, code int) {
	if err != nil {
		http.Error(w, err.Error(), code)
	}
}

func MakeJSONFormatFunc(w http.ResponseWriter,users []models.Users,code int){
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(users)
}