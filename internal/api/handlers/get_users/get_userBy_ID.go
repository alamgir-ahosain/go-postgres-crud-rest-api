package getusers

import (
	"database/sql"
	"net/http"

	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/db"
	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/models"
	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/services"
)

func GetUsersByID(w http.ResponseWriter, r *http.Request) {
	id := services.GetID(w, r) //get id from the url
	if id == 0 {
		return
	}
	var user models.Users
	row := db.DB.QueryRow("select * from users where id=$1", id)
	err := row.Scan(&user.ID, &user.SID, &user.Name, &user.CGPA)
	if err != nil {
		if err == sql.ErrNoRows {
			services.HandleHTTPError(w, err, http.StatusNotFound)
		} else {
			services.HandleHTTPError(w, err, http.StatusInternalServerError)
		}
		return
	}
	services.MakeJSONFormatFunc(w, user, 200)
}
