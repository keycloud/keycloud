package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"webauthn/webauthn"
)

var (
	storage = &Storage{
		authenticators: make(map[string]*Authenticator),
		users:          make(map[string]*User),
	}
	authn               *webauthn.WebAuthn
	err                 error
	store               *sessions.CookieStore
	webAuthnSessionName string = "two-factor-authn-session"
	fileServer          *FileServer
	webauthnHandler     *AuthnHandler
	crudHandler         *CRUDHandler
	secureTokenName     string = "keycloud-secure-key"
	sessionName         string = "keycloud-main"
	userFieldName       string = "keycloud-user-id"
)

func main() {
	/*
		Setup secure cookies
		 - initialize new random keys on every startup -> new login required every server restart
	*/
	rand.Seed(time.Now().UnixNano())
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

	fileServer = &FileServer{
		baseDir:       "./../",
		Router:        webauthnRouter,
		cookieStore:   store,
		sessionName:   sessionName,
		cookieName:    secureTokenName,
		storage:       storage,
		userFieldName: userFieldName,
	}

	webauthnHandler = &AuthnHandler{
		sessionName:       webAuthnSessionName,
		authn:             authn,
		cookieStore:       store,
		storage:           storage,
		securityTokenName: secureTokenName,
		userFieldName:     userFieldName,
		cookieSessionName: sessionName,
	}

	crudHandler = &CRUDHandler{
		cookieStore: store,
		storage:     storage,
	}
	//sign in routes
	webauthnRouter.HandleFunc("/dashboard/login.html", fileServer.ServeFileWithoutCheck).Methods("GET")
	webauthnRouter.HandleFunc("/dashboard/login.css", fileServer.ServeFileWithoutCheck).Methods("GET")
	webauthnRouter.HandleFunc("/dashboard/login.js", fileServer.ServeFileWithoutCheck).Methods("GET")
	webauthnRouter.HandleFunc("/dashboard/icon.png", fileServer.ServeFileWithoutCheck).Methods("GET")
	//other static file routes with permission check middleware
	webauthnRouter.NewRoute().MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
		Path := request.URL.Path[1:]
		return !strings.Contains(Path, "login.html") && !strings.Contains(Path, "login.css") &&
			!strings.Contains(Path, "login.js") &&
			!strings.Contains(Path, "icon.png") && strings.Contains(Path, "dashboard")
	}).Handler(checkCookiePermissionsMiddleware(http.HandlerFunc(fileServer.ServeFileWithoutCheck))).Methods("GET")

	/*
		Web Authn API implementation for 2FA and standard login calls
	*/
	webauthnRouter.HandleFunc("/webauthn/registration/start", webauthnHandler.startRegistration).Methods("POST")
	webauthnRouter.HandleFunc("/webauthn/registration/finish", webauthnHandler.finishRegistration).Methods("POST")
	webauthnRouter.HandleFunc("/webauthn/login/start", webauthnHandler.startLogin).Methods("POST")
	webauthnRouter.HandleFunc("/webauthn/login/finish", webauthnHandler.finishLogin).Methods("POST")
	webauthnRouter.HandleFunc("/standard/login", webauthnHandler.standardLogin).Methods("POST")
	webauthnRouter.HandleFunc("/standard/register", webauthnHandler.standardRegister).Methods("POST")
	webauthnRouter.HandleFunc("/logout", webauthnHandler.logout).Methods("POST")

	/*
		CRUD operations for the Users and the user's passwords
	*/
	webauthnRouter.Handle("/user", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.GetUser))).Methods(http.MethodGet)
	webauthnRouter.Handle("/user", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.RemoveUser))).Methods(http.MethodDelete)
	webauthnRouter.Handle("/user", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.UpdateUser))).Methods(http.MethodPut)
	webauthnRouter.Handle("/password", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.GetPassword))).Methods(http.MethodGet)
	webauthnRouter.Handle("/password", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.AddPassword))).Methods(http.MethodPost)
	webauthnRouter.Handle("/password", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.RemovePassword))).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", webauthnRouter))
}

func checkCookiePermissionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session, err := store.Get(request, sessionName)
		if err != nil {
			//Unauthorized
			writer.WriteHeader(401)
			_, _ = writer.Write([]byte(err.Error()))

		}
		secureCookieRaw, err1 := webauthn.WrapMap(session.Values).Get(secureTokenName)
		userIdRaw, err2 := webauthn.WrapMap(session.Values).Get(userFieldName)
		if err1 != nil || err2 != nil {
			//Unauthorized
			writer.WriteHeader(401)
			_, _ = writer.Write([]byte("401 - Unauthorized - "))
			return
		}
		secureCookie, ok1 := secureCookieRaw.([]byte)
		userId, ok2 := userIdRaw.([]byte)
		user := storage.GetUser(string(userId))
		if !ok1 || !ok2 || user == nil || !bytes.Equal(secureCookie, storage.GetSessionKeyForUser(user)) {
			//Unauthorized
			writer.WriteHeader(401)
			_, _ = writer.Write([]byte("401 - Unauthorized - "))
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if request.Form == nil {
			request.Form = make(map[string][]string)
		}
		request.Form.Add("UserId", string(userId))
		next.ServeHTTP(writer, request)
	})
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateMasterPassword() []byte {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
