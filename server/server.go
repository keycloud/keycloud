package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webauthn/webauthn"
)

type UsernameRequest struct {
	Username string `json:"username"`
}

func startRegistration(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, sessionName)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, ok := storage.users[name]
	if !ok {
		u = &User{
			Name:           name,
			Authenticators: make(map[string]*Authenticator),
		}
		storage.users[name] = u
	}
	options := authn.StartRegistration(request, writer, u, webauthn.WrapMap(session.Values))
	err = session.Save(request, writer)
	if err != nil{
	http.Error(writer, err.Error(), http.StatusInternalServerError)
	return
	}
	authn.Write(request, writer, options)
}

func finishRegistration(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, sessionName)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, ok := storage.users[name]
	if !ok {
		return
	}
	authn.FinishRegistration(request, writer, u, webauthn.WrapMap(session.Values), b)
	err = session.Save(request, writer)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startLogin(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, sessionName)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, ok := storage.users[name]
	if !ok {
		http.Error(writer, "No User", http.StatusInternalServerError)
	}
	options := authn.StartLogin(request, writer, u, webauthn.WrapMap(session.Values))
	err = session.Save(request, writer)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	authn.Write(request, writer, options)
}

func finishLogin(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, sessionName)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var usernameMsg UsernameRequest
	err = json.Unmarshal(b, &usernameMsg)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	name := usernameMsg.Username
	u, ok := storage.users[name]
	if !ok {
		http.Error(writer, "No such user", http.StatusUnauthorized)
		return
	}
	auth_ := authn.FinishLogin(request, writer, u, webauthn.WrapMap(session.Values), b)
	_, ok = auth_.(*Authenticator)
	if !ok{
		http.Error(writer, "Auth error", http.StatusInternalServerError)
	}
	//Set login cookie...
	err = session.Save(request, writer)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(writer, "Logged in")
}
