package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"strings"
	"webauthn/webauthn"
)

var (
	storage = &Storage{
		authenticators: make(map[string]*Authenticator),
		users:          make(map[string]*User),
	}
	authn       *webauthn.WebAuthn
	err         error
	store       *sessions.CookieStore
	sessionName = "two-factor-authn-session"
	fileserver  *FileServer
)

func main() {
	/*
		Setup secure cookies
		 - initialize new random keys on every startup -> new login required every server restart
	*/
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)
	store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)

	authn, err = webauthn.New(&webauthn.Config{
		// A human-readable identifier for the relying party (i.e. your app), intended only for display.
		RelyingPartyName: "KeyCloud",
		// Storage for the authenticator.
		AuthenticatorStore: storage,
	})
	if err != nil {
		panic("Unable to create 2FA server")
	}

	webauthnRouter := mux.NewRouter()

	fileserver = &FileServer{
		baseDir:       "./../",
		Router:        webauthnRouter,
		cookieStore:   store,
		sessionName:   "keycloud_main",
		cookieName:    "keycloud_secure_key",
		storage:       storage,
		userFieldName: "keycloud_user_id",
	}
	//sign in routes
	webauthnRouter.HandleFunc("/dashboard/index.html", fileserver.ServeFileWithoutCheck)
	webauthnRouter.HandleFunc("/dashboard/signin.css", fileserver.ServeFileWithoutCheck)
	webauthnRouter.HandleFunc("/dashboard/icon.png", fileserver.ServeFileWithoutCheck)
	//other routes
	webauthnRouter.NewRoute().MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
		Path := request.URL.Path[1:]
		return !strings.Contains(Path, "index.html") && !strings.Contains(Path, "signin.css") &&
			!strings.Contains(Path, "icon.png")
	}).HandlerFunc(fileserver.ServeFileWithCookieCheck)

	/*
		Web Authn API implementation for 2FA
	*/
	webauthnRouter.HandleFunc("/webauthn/registration/start", startRegistration)
	webauthnRouter.HandleFunc("/webauthn/registration/finish", finishRegistration)
	webauthnRouter.HandleFunc("/webauthn/login/start", startLogin)
	webauthnRouter.HandleFunc("/webauthn/login/finish", finishLogin)

	log.Fatal(http.ListenAndServe(":8080", webauthnRouter))
}
