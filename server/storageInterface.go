package main

import (
	"github.com/keycloud/webauthn/webauthn"
)

type User struct {
	Name           string                    `json:"name"`
	Authenticators map[string]*Authenticator `json:"-"`
	MasterPassword []byte
	Mail           string
	Uuid           []byte
}

type Authenticator struct {
	User         *User
	ID           []byte
	CredentialID []byte
	PublicKey    []byte
	AAGUID       []byte
	SignCount    uint32
}

type Password struct {
	Password string `json:"password"`
	Id       string `json:"id"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

func (u *User) WebAuthID() []byte {
	return []byte(u.Uuid)
}

func (u *User) WebAuthName() string {
	return u.Name
}

func (u *User) WebAuthDisplayName() string {
	return u.Name
}

func (a *Authenticator) WebAuthID() []byte {
	return a.ID
}

func (a *Authenticator) WebAuthCredentialID() []byte {
	return a.CredentialID
}

func (a *Authenticator) WebAuthPublicKey() []byte {
	return a.PublicKey
}

func (a *Authenticator) WebAuthAAGUID() []byte {
	return a.AAGUID
}

func (a *Authenticator) WebAuthSignCount() uint32 {
	return a.SignCount
}

type StorageInterface interface {
	AddAuthenticator(webauthn.User, webauthn.Authenticator) error
	GetAuthenticator([]byte) (webauthn.Authenticator, error)
	GetAuthenticators(webauthn.User) ([]webauthn.Authenticator, error)
	// User operations
	GetUser(webauthnID string) (*User, error)
	GetUserByName(name string) (*User, error)
	CreateUser(*User) error
	RemoveUser(*User) error
	UpdateUser(*User) error
	// Session operations
	GetSessionKeyForUser(*User) ([]byte, error)
	UpdateOrCreateSessionKeyForUser(*User, []byte) error
	DeleteSessionKeyForUser(*User) error
	// Password operations
	GetPassword(user *User, url string, username string) (*Password, error)
	GetPasswordByUrl(user *User, url string) ([] *Password, error)
	GetPasswords(*User) ([] *Password, error)
	CreatePassword(*User, string, *Password) error
	UpdatePassword(*User, string, *Password) error
	DeletePassword(user *User, url string, username string) error
}
