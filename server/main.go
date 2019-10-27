package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"webauthn/webauthn"
)

var(
	storage = &Storage{
		authenticators: make(map[string]* Authenticator),
		users:          make(map[string]* User),
	}
	authn *webauthn.WebAuthn
	err error
	store *sessions.CookieStore
	sessionName = "two-factor-authn-session"
)

func main() {
	/*
	Setup secure cookies
	 - initialize new random keys on every startup -> new login required every server restart
	*/
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)
	store =  sessions.NewCookieStore(authKeyOne, encryptionKeyOne)

	authn, err = webauthn.New(&webauthn.Config{
		// A human-readable identifier for the relying party (i.e. your app), intended only for display.
		RelyingPartyName:   "KeyCloud",
		// Storage for the authenticator.
		AuthenticatorStore: storage,
	})
	if err != nil{
		panic("Unable to create 2FA server")
	}

	webauthnRouter := mux.NewRouter()

	webauthnRouter.HandleFunc("/dashboard/index.html", serveFileWithoutCheck)
	webauthnRouter.HandleFunc("/dashboard/signin.css", serveFileWithoutCheck)
	webauthnRouter.HandleFunc("/dashboard/icon.png", serveFileWithoutCheck)

	/*
	Web Authn API implementation for 2FA
	*/
	webauthnRouter.HandleFunc("/webauthn/registration/start", startRegistration)
	webauthnRouter.HandleFunc("/webauthn/registration/finish", finishRegistration)
	webauthnRouter.HandleFunc("/webauthn/login/start", startLogin)
	webauthnRouter.HandleFunc("/webauthn/login/finish", finishLogin)
	log.Fatal(http.ListenAndServe(":8080", webauthnRouter))
}

func serveFileWithoutCheck(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[1:]
	log.Println(path)
	path = "./../" + path
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
		_, _ = writer.Write([]byte("404 What are you doing - " + http.StatusText(404)))
	}
}
