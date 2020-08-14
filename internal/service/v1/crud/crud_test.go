package crud

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Create(t *testing.T) {

}

func TestAPI_Select(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "age", "create_time"}).AddRow("1", "post 1", 12, time.Now())

	mock.ExpectQuery("SELECT id, name, age, create_time FROM account").WillReturnRows(rows)
	api := API{
		DB: db,
	}
	r := httptest.NewRequest("GET", "/connect", nil)
	w := httptest.NewRecorder()
	api.Select(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}
