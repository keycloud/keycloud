package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/keycloud/webauthn/webauthn"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const(
	secureTokenName     string = "keycloud-secure-key"
	sessionName         string = "keycloud-main"
	userFieldName       string = "keycloud-user-id"
	webAuthnSessionName string = "two-factor-authn-session"
)

var (
	authn               *webauthn.WebAuthn
	err                 error
	store               *sessions.CookieStore
	fileServer          *FileServer
	webauthnHandler     *AuthnHandler
	crudHandler         *CRUDHandler
	database 			*sql.DB
	storage				StorageInterface
)

func initFromDatabaseAndRouter(db *sql.DB){
	database = db
	rand.Seed(time.Now().UnixNano())
	//authKeyOne := securecookie.GenerateRandomKey(64)
	//encryptionKeyOne := securecookie.GenerateRandomKey(32)
	// TODO: change for productive server again
	store = sessions.NewCookieStore([]byte("aaaaaaaaaaaaaaaa"), []byte("aaaaaaaaaaaaaaaa"))

	storage = &Storage{
		database: db,
	}

	authn, err = webauthn.New(&webauthn.Config{
		RelyingPartyName: "KeyCloud",
		AuthenticatorStore: storage,
	})
	if err != nil {
		panic("Unable to create 2FA server")
	}

	fileServer = &FileServer{
		baseDir:       "./../",
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
}


func main() {


	// Connect to database
	database, err = connectDatabase()
	defer database.Close()
	if err != nil{
		panic(err)
	}
	
	// Create Tables if they do not exist yet
	data, _ := ioutil.ReadFile("init.sql")
	_, err = database.Exec(string(data))

	// Delete all previous stored Sessions
	err = ClearAllSessionKeys(database)
	if err != nil{
		fmt.Println(err)
	}

	initFromDatabaseAndRouter(database)

	webauthnRouter := mux.NewRouter()

	webauthnRouter.HandleFunc("/.well-known/assetlinks.json", assetLinksHandler)
	webauthnRouter.PathPrefix("/dashboard/").Handler(http.StripPrefix("/dashboard/", http.FileServer(http.Dir("./dashboard/"))))

	/*
		Web Authn API implementation for 2FA and standard login calls
	*/
	// Registration (adding a new 2FA) of a new authenticator should only be allowed when already logged in
	// -> therefore also only available after cookie check
	webauthnRouter.Handle("/webauthn/registration/start", checkCookiePermissionsMiddleware(
		http.HandlerFunc(webauthnHandler.startRegistration))).Methods(http.MethodPost)
	webauthnRouter.Handle("/webauthn/registration/finish", checkCookiePermissionsMiddleware(
		http.HandlerFunc(webauthnHandler.finishRegistration))).Methods(http.MethodPost)

	webauthnRouter.HandleFunc("/webauthn/login/start", webauthnHandler.startLogin).Methods(http.MethodPost)
	webauthnRouter.HandleFunc("/webauthn/login/finish", webauthnHandler.finishLogin).Methods(http.MethodPost)
	webauthnRouter.HandleFunc("/standard/login", webauthnHandler.standardLogin).Methods(http.MethodPost)
	webauthnRouter.HandleFunc("/standard/register", webauthnHandler.standardRegister).Methods(http.MethodPost)

	webauthnRouter.Handle("/logout", checkCookiePermissionsMiddleware(http.HandlerFunc(webauthnHandler.logout))).Methods(http.MethodPost)

	/*
		CRUD operations for the Users and the user's passwords
	*/
	webauthnRouter.Handle("/user", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.GetUser))).Methods(http.MethodGet)
	webauthnRouter.Handle("/user", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.RemoveUser))).Methods(http.MethodDelete)
	webauthnRouter.Handle("/user", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.UpdateUser))).Methods(http.MethodPut)
	webauthnRouter.Handle("/password", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.GetPassword))).Methods(http.MethodGet)
	webauthnRouter.Handle("/passwords", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.GetPasswords))).Methods(http.MethodGet)
	webauthnRouter.Handle("/password", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.CreatePassword))).Methods(http.MethodPost)
	webauthnRouter.Handle("/password", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.RemovePassword))).Methods(http.MethodDelete)
	webauthnRouter.Handle("/password-by-url", checkCookiePermissionsMiddleware(http.HandlerFunc(crudHandler.GetPasswordByUrl))).Methods(http.MethodGet, http.MethodPost)

	panic(http.ListenAndServe(":8080", webauthnRouter))
}

func assetLinksHandler(writer http.ResponseWriter, request *http.Request) {
	println("Assetlinks file requested")
	data, _ := ioutil.ReadFile("assetlinks.json")
	writer.Header().Add("Content-Type", "application/json")
	_, _ = writer.Write(data)
}

func checkCookiePermissionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session, err := store.Get(request, sessionName)
		if err != nil {
			//Unauthorized
			writer.WriteHeader(401)
			_, _ = writer.Write([]byte("401 - Unauthorized - "))
			return
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
		user, err1 := storage.GetUser(string(userId))
		sessionKey, err2 := storage.GetSessionKeyForUser(user)
		if !ok1 || !ok2 || user == nil || err1 != nil || err2 != nil || !bytes.Equal(secureCookie, sessionKey) {
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
