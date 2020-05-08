package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
)

type CRUDHandler struct {
	cookieStore *sessions.CookieStore
	storage     StorageInterface
}

type GetPasswordRequest struct {
	Url      string `json:"url"`
	Username string `json:"username"`
}

type UserRequest struct {
	Name string `json:"name"`
}

type PasswordRequest struct {
	Password string `json:"password"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

func (handler CRUDHandler) GetPassword(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.storage.GetUser(request.Form.Get("UserId"))
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var passwordId GetPasswordRequest
	err = json.Unmarshal(b, &passwordId)
	password, err := handler.storage.GetPassword(user, passwordId.Url, passwordId.Username)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	/*
		Send the password "plain" as received from the database, Encryption and Decryption in frontend
	*/
	passwordJson, err := json.Marshal(password)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(writer, string(passwordJson))
}

func (handler CRUDHandler) GetPasswordByUrl(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.storage.GetUser(request.Form.Get("UserId"))
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var passwordId GetPasswordRequest
	err = json.Unmarshal(b, &passwordId)
	passwords, err := handler.storage.GetPasswordByUrl(user, passwordId.Url)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	/*
		Send the password "plain" as received from the database, Encryption and Decryption in frontend
	*/
	passwordJson, err := json.Marshal(passwords)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(writer, string(passwordJson))
}

func (handler CRUDHandler) GetPasswords(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.storage.GetUser(request.Form.Get("UserId"))
	passwords, err := handler.storage.GetPasswords(user)
	responseMessagJSON, err := json.Marshal(passwords)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	_, _ = fmt.Fprint(writer, string(responseMessagJSON))
}

func (handler CRUDHandler) CreatePassword(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.storage.GetUser(request.Form.Get("UserId"))
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var password PasswordRequest
	err = json.Unmarshal(b, &password)
	err = handler.storage.CreatePassword(user, password.Url, &Password{
		Password: password.Password,
		Url:      password.Url,
		Username: password.Username,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("CREATED", "", writer)
}

func (handler CRUDHandler) RemovePassword(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.storage.GetUser(request.Form.Get("UserId"))
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var passwordId GetPasswordRequest
	err = json.Unmarshal(b, &passwordId)
	err = handler.storage.DeletePassword(user, passwordId.Url, passwordId.Username)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("REMOVED", "", writer)
}

func (handler CRUDHandler) RemoveUser(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.storage.GetUser(request.Form.Get("UserId"))
	//TODO: remove all passwords before removing the User either as go implementation or as db constraint etc.
	err = handler.storage.RemoveUser(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("REMOVED", "", writer)
}

func (handler CRUDHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var user UserRequest
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.storage.UpdateUser(&User{
		Name: user.Name,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("UPDATED", "", writer)
}

func (handler CRUDHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.storage.GetUser(request.Form.Get("UserId"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// Wrap struct around internal User struct to enforce string encoding of master password
	userObject := struct {
		Name           string
		MasterPassword string
	}{
		user.Name,
		string(user.MasterPassword),
	}
	userJson, err := json.Marshal(userObject)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(writer, string(userJson))
}

func sendCRUDAnswer(statusMessage string, errMessage string, writer http.ResponseWriter) {
	responseMessageJSON, err := json.Marshal(struct {
		Status string
		Error  string
	}{
		Status: statusMessage,
		Error:  errMessage,
	})
	checkError(err, writer)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(writer, string(responseMessageJSON))
}
