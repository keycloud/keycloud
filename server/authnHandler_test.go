package main

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestAuthnHandler_standardLogin(t *testing.T) {
	req, err := http.NewRequest("POST", "/standard/login",
		bytes.NewBuffer([]byte(`{"username": "johndoe", "masterpassword": "my-master-passwd"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("johndoe").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID", "johndoe", "@", "my-master-passwd"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO sessions").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(webauthnHandler.standardLogin)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Logged in`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthnHandler_standardRegister(t *testing.T) {
	req, err := http.NewRequest("POST", "/standard/register/start",
		bytes.NewBuffer([]byte(`{"username": "johndoe", "mail": "john@doe.com"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("johndoe").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO users").
		ExpectExec().WithArgs(sqlmock.AnyArg(), "johndoe", "john@doe.com", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO sessions").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(webauthnHandler.standardRegister)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Registered`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthnHandler_startRegistration(t *testing.T) {
	req, err := http.NewRequest("POST", "/webauthn/register",
		bytes.NewBuffer([]byte(`{"username": "johndoe", "mail": "john@doe.com"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT uuid, name, mail, masterpasswd FROM users").
		ExpectQuery().WithArgs("johndoe").WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO users").
		ExpectExec().WithArgs(sqlmock.AnyArg(), "johndoe", "john@doe.com", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT id, credentialid, publickey, aaguid, signcount FROM authenticators").
		ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{"id", "credentialid", "publickey", "aaguid", "signcount"}))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(webauthnHandler.startRegistration)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"publicKey":{"rp":{"name":"KeyCloud"},"user":{"name":"johndoe","id":"[^"]+","displayName":"johndoe"},"challenge":"[^"]+","pubKeyCredParams":[{"type":"public-key","alg":-7}],"timeout":30000,"authenticatorSelection":{"requireResidentKey":false},"attestation":"direct"}}`
	_, ok := rr.Header()["Set-Cookie"]
	if matched, _ := regexp.MatchString(expected, rr.Body.String()); matched && ok {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthnHandler_startLogin(t *testing.T) {
	req, err := http.NewRequest("POST", "/webauthn/login/start",
		bytes.NewBuffer([]byte(`{"username": "johndoe", "mail": "john@doe.com"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a request", err)
	}
	if req.Form == nil {
		req.Form = make(map[string][]string)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users").
		ExpectQuery().WithArgs("johndoe").
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "name", "mail", "masterpasswd"}).
			AddRow("USERID", "johndoe", "@", "my-master-passwd"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT id, credentialid, publickey, aaguid, signcount FROM authenticators").
		ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{"id", "credentialid", "publickey", "aaguid", "signcount"}))
	mock.ExpectCommit()

	// Set global values to mocked one
	initFromDatabaseAndRouter(db)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(webauthnHandler.startLogin)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"publicKey":{"rp":{"name":"KeyCloud"},"user":{"name":"johndoe","id":"[^"]+","displayName":"johndoe"},"challenge":"[^"]+","pubKeyCredParams":[{"type":"public-key","alg":-7}],"timeout":30000,"authenticatorSelection":{"requireResidentKey":false},"attestation":"direct"}}`
	_, ok := rr.Header()["Set-Cookie"]
	if matched, _ := regexp.MatchString(expected, rr.Body.String()); matched && ok {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
