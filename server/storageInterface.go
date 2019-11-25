package main

import "webauthn/webauthn"

type User struct {
	Name           string                    `json:"name"`
	Authenticators map[string]*Authenticator `json:"-"`
	MasterPassword []byte
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
	Password string
	Id       string
	//TODO: Add further attributes
}

func (u *User) WebAuthID() []byte {
	return []byte(u.Name)
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
	GetUser(webauthnID string) *User
	AddUser(*User) error
	RemoveUser(*User) error
	UpdateUser(*User) error
	GetSessionKeyForUser(*User) []byte
	SetSessionKeyForUser(*User, []byte) error
	DeleteSessionKeyForUser(*User) error
	GetPassword(*User, string) (*Password, error)
	AddPassword(*User, string, *Password) error
	UpdatePassword(*User, string, *Password) error
	DeletePassword(*User, string) error
}
