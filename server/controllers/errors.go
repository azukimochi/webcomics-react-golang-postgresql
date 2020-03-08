package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/azukimochi/webcomics-react-golang-postgresql/server/models"
)

const (
	// DBQueryFailed is an error code for when the db query failed
	DBQueryFailed = "DB_QUERY_FAILED"
	// DBChangeFailed is an error code for when the app failed to make a change in the database
	DBChangeFailed = "DB_CHANGE_FAILED"
	// InvalidURLParams is an error code for when there was an error when reading the URL parameters
	InvalidURLParams = "INVALID_URL_PARAMS"
	// GeneralDBError is an error code for when there is a general error in the database transaction
	GeneralDBError = "GENERAL_DB_ERROR"
)

// ServeError is a function that handles serving up the json response for the errors
func ServeError(code string, description string, w http.ResponseWriter, r *http.Request) {
	errMsg := models.ErrResponse{
		Code:        code,
		Description: description,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(errMsg)
}
