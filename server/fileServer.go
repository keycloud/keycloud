package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
	"strings"
	"webauthn/webauthn"
)

type FileServer struct {
	baseDir       string // "./../"
	Router        *mux.Router
	cookieStore   *sessions.CookieStore
	sessionName   string
	cookieName    string
	storage       *Storage
	userFieldName string
}

func (server FileServer) ServeFileWithoutCheck(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[1:]
	path = server.baseDir + path
	data, err := ioutil.ReadFile(path)

	if err == nil {
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/plain"
		}
		writer.Header().Add("Content-Type", contentType)
		_, _ = writer.Write(data)
	} else {
		writer.WriteHeader(404)
		_, _ = writer.Write([]byte("404 - File not found - "))
	}
}

func (server FileServer) ServeFileWithCookieCheck(writer http.ResponseWriter, request *http.Request) {
	session, err := server.cookieStore.Get(request, server.sessionName)
	if err != nil {
		//Unauthorized
		writer.WriteHeader(401)
		_, _ = writer.Write([]byte("401 - Unauthorized - "))
		return
	}
	secureCookieRaw, err1 := webauthn.WrapMap(session.Values).Get(server.cookieName)
	userIdRaw, err2 := webauthn.WrapMap(session.Values).Get(server.userFieldName)
	if err1 != nil || err2 != nil {
		//Unauthorized
		writer.WriteHeader(401)
		_, _ = writer.Write([]byte("401 - Unauthorized - "))
		return
	}
	secureCookie, ok1 := secureCookieRaw.([]byte)
	userId, ok2 := userIdRaw.(string)
	user := server.storage.GetUser(userId)
	if !ok1 || !ok2 || user == nil || !bytes.Equal(secureCookie, user.GetSecurityToken()) {
		//Unauthorized
		writer.WriteHeader(401)
		_, _ = writer.Write([]byte("401 - Unauthorized - "))
		return
	}
	path := request.URL.Path[1:]
	path = server.baseDir + path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/plain"
		}
		writer.Header().Add("Content-Type", contentType)
		_, _ = writer.Write(data)
	} else {
		writer.WriteHeader(404)
		_, _ = writer.Write([]byte("404 - File not found - " + http.StatusText(404)))
	}
}
