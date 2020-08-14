package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/pingpotter/rest-pg/internal/middleware"
	"github.com/pingpotter/rest-pg/internal/service/v1/crud"
)

// Router ...
func Router(db *sql.DB) *mux.Router {

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/healthz", healthz).Methods("GET")

	curlAPI := crud.API{
		DB: db,
	}

	r.HandleFunc("/user", curlAPI.Select).Methods("GET")
	r.HandleFunc("/user", curlAPI.Create).Methods("POST")

	return r
}
