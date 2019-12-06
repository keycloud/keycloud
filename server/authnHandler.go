package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"webauthn/webauthn"
	"io/ioutil"
	"net/http"
)

type AuthnHandler struct {
	sessionName       string
	authn             *webauthn.WebAuthn
	cookieStore       *sessions.CookieStore
	storage           StorageInterface
	securityTokenName string
	userFieldName     string
	cookieSessionName string
}
type UsernameRequest struct {
	Username string `json:"username"`
	Mail 	 string `json:"mail"`
}

type UsernamePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (handler AuthnHandler) startRegistration(writer http.ResponseWriter, request *http.Request) {
	session, err := handler.cookieStore.Get(request, handler.sessionName)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, err := handler.storage.GetUser(name)
	if u == nil {
		u = &User{
			Name:           name,
			Authenticators: make(map[string]*Authenticator),
			MasterPassword: GeneratePassword(16),
		}
		err = handler.storage.CreateUser(u)
	}
	options := handler.authn.StartRegistration(request, writer, u, webauthn.WrapMap(session.Values))
	err = session.Save(request, writer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.authn.Write(request, writer, options)
}

func (handler AuthnHandler) finishRegistration(writer http.ResponseWriter, request *http.Request) {
	session, err := handler.cookieStore.Get(request, handler.sessionName)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, err := handler.storage.GetUserByName(name)
	if u == nil {
		return
	}
	handler.authn.FinishRegistration(request, writer, u, webauthn.WrapMap(session.Values), b)
	SaveLoginInSession(handler, writer, request, u)
}

func (handler AuthnHandler) startLogin(writer http.ResponseWriter, request *http.Request) {
	session, err := handler.cookieStore.Get(request, handler.sessionName)
	checkError(err, writer)
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, err := handler.storage.GetUserByName(name)
	if err != nil {
		http.Error(writer, "No User", http.StatusInternalServerError)
	}
	options := handler.authn.StartLogin(request, writer, u, webauthn.WrapMap(session.Values))
	err = session.Save(request, writer)
	checkError(err, writer)
	handler.authn.Write(request, writer, options)
}

func (handler AuthnHandler) finishLogin(writer http.ResponseWriter, request *http.Request) {
	session, err := handler.cookieStore.Get(request, handler.sessionName)
	checkError(err, writer)
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, err := handler.storage.GetUserByName(name)
	if err != nil {
		http.Error(writer, "No such user", http.StatusUnauthorized)
		return
	}
	auth_ := handler.authn.FinishLogin(request, writer, u, webauthn.WrapMap(session.Values), b)
	_, ok := auth_.(*Authenticator)
	if !ok {
		http.Error(writer, "Auth Error", http.StatusInternalServerError)
	}
	SaveLoginInSession(handler, writer, request, u)
}

func (handler AuthnHandler) standardLogin(writer http.ResponseWriter, request *http.Request) {
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var userPasswordMsg UsernamePasswordRequest
	err = json.Unmarshal(b, &userPasswordMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	name := userPasswordMsg.Username
	u, err := handler.storage.GetUserByName(name)
	if err != nil {
		http.Error(writer, "401 - Unauthorized - ", http.StatusUnauthorized)
		return
	}
	if !bytes.Equal(u.MasterPassword, []byte(userPasswordMsg.Password)) {
		http.Error(writer, "401 - Unauthorized - ", http.StatusUnauthorized)
		return
	}
	SaveLoginInSession(handler, writer, request, u)
	_, _ = fmt.Fprint(writer, "Logged in")
}

func (handler AuthnHandler) standardRegister(writer http.ResponseWriter, request *http.Request) {
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var userMsg UsernameRequest
	err = json.Unmarshal(b, &userMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	user, err := handler.storage.GetUserByName(userMsg.Username)
	if (user != nil && user.Uuid != nil) || err != nil{
		responseMessageJSON, err := json.Marshal(struct {
			Status string
			Error  string
		}{
			Status: "Not Registered",
			Error:  "Username already exists",
		})
		checkError(err, writer)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(writer, string(responseMessageJSON))
		return
	}
	user = &User{
		Name:           userMsg.Username,
		Authenticators: nil,
		MasterPassword: GeneratePassword(16),
		Mail:           userMsg.Mail,
	}
	err = handler.storage.CreateUser(user)
	SaveLoginInSession(handler, writer, request, user)
	_, _ = fmt.Fprint(writer, "Registered")
}

func (handler AuthnHandler) logout(writer http.ResponseWriter, request *http.Request) {
	//Clear session cookies
	session, err := handler.cookieStore.Get(request, handler.cookieSessionName)
	checkError(err, writer)
	err = handler.storage.DeleteSessionKeyForUser(&User{
		Name:           "",
		Authenticators: nil,
		MasterPassword: nil,
		Mail:           "",
		Uuid:           []byte(request.Form.Get("UserId")),
	})
	checkError(err, writer)
	session.Values = nil
	session.Options.MaxAge = -1
	_ = session.Save(request, writer)
	_, _ = fmt.Fprint(writer, "Logged out")

}

func checkError(err error, writer http.ResponseWriter) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SaveLoginInSession(handler AuthnHandler, writer http.ResponseWriter, request *http.Request, u *User) {
	session, err := handler.cookieStore.New(request, handler.cookieSessionName)
	checkError(err, writer)
	cookieSecurityToken := GeneratePassword(16)
	sessionValues := webauthn.WrapMap(session.Values)
	err = sessionValues.Set(handler.securityTokenName, cookieSecurityToken)
	err = sessionValues.Set(handler.userFieldName, u.WebAuthID())
	checkError(err, writer)
	err = session.Save(request, writer)
	checkError(err, writer)
	err = handler.storage.UpdateOrCreateSessionKeyForUser(u, cookieSecurityToken)
	checkError(err, writer)
}
