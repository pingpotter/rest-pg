package crud

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

type API struct {
	DB *sql.DB
}

type account struct {
	ID         string
	Name       string
	Age        int16
	CreateTime time.Time
}

// Create create db
func (a API) Create(w http.ResponseWriter, r *http.Request) {

	u := account{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	const query = `INSERT INTO account (id, name, age, create_time)VALUES ($1, $2, $3, $4)`

	_, err := a.DB.Exec(query, u.ID, u.Name, u.Age, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Select select db
func (a API) Select(w http.ResponseWriter, _ *http.Request) {

	const query = `SELECT id, name, age, create_time FROM account`

	allUser := []account{}
	rows, err := a.DB.Query(query)
	if err != nil {
		log.Error("error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {

		var age int16
		var id, name string
		var timestamp time.Time

		if err := rows.Scan(&id, &name, &age, &timestamp); err != nil {
			log.Error(err)
			continue
		}

		u := account{
			ID:         id,
			Name:       name,
			Age:        age,
			CreateTime: timestamp,
		}

		allUser = append(allUser, u)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(allUser)
}
