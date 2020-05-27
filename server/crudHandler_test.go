package main

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCRUDHandler_GetPassword(t *testing.T) {
	req, err := http.NewRequest("GET", "/password", bytes.NewBuffer([]byte(`{"username": "johndoe", "url": "john.doe"}`)))
	if err != nil{
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}
	req.Form.Add("UserId", string("USERID"))

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("USERID").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID","john","@","password"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM passwds").
		ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{"entryid", "url", "passwd", "username"}).
			AddRow(1,"john.doe","password","johndoe"))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crudHandler.GetPassword)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"password":"password","id":"1","url":"john.doe","username":"johndoe"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCRUDHandler_GetPasswords(t *testing.T) {
	req, err := http.NewRequest("GET", "/passwords", nil)
	if err != nil{
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}
	req.Form.Add("UserId", string("USERID"))

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("USERID").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID","john","@","password"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM passwds").
		ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{"entryid", "url", "passwd", "username"}).
		AddRow(1,"john.doe","password","johndoe"))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crudHandler.GetPasswords)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"password":"password","id":"1","url":"john.doe","username":"johndoe"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCRUDHandler_CreatePassword(t *testing.T) {
	req, err := http.NewRequest("POST", "/password", bytes.NewBuffer([]byte(`{"username": "johndoe", "password": "doejohn", "url": "john.doe"}`)))
	if err != nil{
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}
	req.Form.Add("UserId", string("USERID"))

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("USERID").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID","john","@","password"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO passwds").
		ExpectExec().WithArgs(sqlmock.AnyArg(), "john.doe", "doejohn", "johndoe").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crudHandler.CreatePassword)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Status":"CREATED","Error":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCRUDHandler_RemovePassword(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/password", bytes.NewBuffer([]byte(`{"username": "johndoe", "url": "john.doe"}`)))
	if err != nil{
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}
	req.Form.Add("UserId", string("USERID"))

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("USERID").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID","john","@","password"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("DELETE FROM passwds").
		ExpectExec().WithArgs(sqlmock.AnyArg(), "john.doe", "johndoe").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crudHandler.RemovePassword)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Status":"REMOVED","Error":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCRUDHandler_RemoveUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user", nil)
	if err != nil{
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}
	req.Form.Add("UserId", string("USERID"))

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("USERID").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID","john","@","password"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("DELETE FROM users").
		ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crudHandler.RemoveUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Status":"REMOVED","Error":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCRUDHandler_UpdateUser(t *testing.T) {
	req, err := http.NewRequest("PUT", "/user", bytes.NewBuffer([]byte(`{"name": "johndoe"}`)))
	if err != nil{
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}
	req.Form.Add("UserId", string("USERID"))

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("USERID").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID","john","@","password"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("UPDATE users").
		ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crudHandler.UpdateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Status":"UPDATED","Error":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCRUDHandler_GetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil{
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}
	req.Form.Add("UserId", string("USERID"))

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("USERID").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID","john","@","password"))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crudHandler.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"username":"john","masterpassword":"password","mail":"@"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
