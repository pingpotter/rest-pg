package crud

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Create(t *testing.T) {

	timeNow := time.Now()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	const query = `INSERT INTO account`

	mock.ExpectExec(query).WithArgs("1", "test product", 11, timeNow).WillReturnResult(sqlmock.NewResult(1, 1))

	api := API{
		DB:      db,
		TimeNow: timeNow,
	}

	jsonStr := []byte(`{ "id":"1", "name":"test product", "age": 11}`)

	r := httptest.NewRequest("POST", "/connect", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	r.Header.Set("Content-Type", "application/json")

	api.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "", w.Body.String())
}

// func TestAPI_Select(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	rows := sqlmock.NewRows([]string{"id", "name", "age", "create_time"}).
// 		AddRow("1", "post 1", 12, time.Now()).
// 		AddRow("2", "post 2", 32, time.Now())

// 	mock.ExpectQuery("SELECT id, name, age, create_time FROM account").WillReturnRows(rows)

// 	api := API{
// 		DB: db,
// 	}

// 	r := httptest.NewRequest("GET", "/connect", nil)
// 	w := httptest.NewRecorder()

// 	api.Select(w, r)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, "", w.Body.String())

// }

// func TestAPI_Select_Error_FROM_Query(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectQuery("SELECT id, name, age, create_time FROM account").WillReturnError(errors.New("Query error"))
// 	api := API{
// 		DB: db,
// 	}

// 	r := httptest.NewRequest("GET", "/connect", nil)
// 	w := httptest.NewRecorder()

// 	api.Select(w, r)

// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	assert.Equal(t, "Query error\n", w.Body.String())

// }

func TestAPI_Select_Error_FROM_Scan(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "age", "create_time"}).
		AddRow("1", "post 1", 12, time.Now()).
		RowError(0, fmt.Errorf("error form rows"))

	mock.ExpectQuery("SELECT id, name, age, create_time FROM account").WillReturnRows(rows)

	api := API{
		DB: db,
	}

	r := httptest.NewRequest("GET", "/connect", nil)
	w := httptest.NewRecorder()

	api.Select(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, "[]", w.Body.String())
}
