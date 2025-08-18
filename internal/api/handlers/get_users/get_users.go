package getusers

import (
	"net/http"
	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/db"
	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/models"
	"github.com/alamgir-ahosain/go-postgres-crud-rest-api/internal/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	row, err := db.DB.Query("select * from users")
	services.HandleHTTPError(w, err, http.StatusInternalServerError)

	defer row.Close() // close  when main exist

	// Add users to the slice which returned by the query
	for row.Next() {
		var user models.Users
		err := row.Scan(&user.ID, &user.SID, &user.Name, &user.CGPA)
		services.HandleHTTPError(w, err, http.StatusInternalServerError)

		users = append(users, user)
	}
	services.MakeJSONFormatFunc(w, users, 200)
}
