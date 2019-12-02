package main

import (
	"database/sql"
	"webauthn/webauthn"
)

type Storage struct {
	database *sql.DB
}

/*
	Authenticator operations
*/
func (s *Storage) AddAuthenticator(user webauthn.User, authenticator webauthn.Authenticator) error {
	authr := &Authenticator{
		ID:           authenticator.WebAuthID(),
		CredentialID: authenticator.WebAuthCredentialID(),
		PublicKey:    authenticator.WebAuthPublicKey(),
		AAGUID:       authenticator.WebAuthAAGUID(),
		SignCount:    authenticator.WebAuthSignCount(),
	}
	return CreateAuthenticator(s.database, authr, user.WebAuthID())
}

func (s *Storage) GetAuthenticator(id []byte) (webauthn.Authenticator, error) {
	auth, err := QueryAuthenticator(s.database, id)
	return auth, err
}

func (s *Storage) GetAuthenticators(user webauthn.User) ([]webauthn.Authenticator, error) {
	authn, err := QueryAllAuthenticators(s.database, user.WebAuthID())
	return authn, err
}

/*
	Session operations
*/
func (s *Storage) GetSessionKeyForUser(user *User) ([]byte, error) {
	token, err := QuerySessionForUser(s.database, user)
	if err != nil{
		return []byte(""), nil
	}
	return token, err
}

func (s *Storage) UpdateOrCreateSessionKeyForUser(user *User, b []byte) error {
	return UpdateOrCreateSessionKeyForUser(s.database, user, b)
}

func (s *Storage) DeleteSessionKeyForUser(user *User) error {
	return DeleteSessionKeyForUser(s.database, user)
}

/*
	User operations
*/
func (s *Storage) GetUser(ID string) (*User, error) {
	return QueryUser(s.database, ID)
}
func (s *Storage) GetUserByName(name string) (*User, error) {
	return QueryUserByName(s.database, name)
}

func (s *Storage) CreateUser(u *User) error {
	return CreateUser(s.database, u)
}

func (s *Storage) RemoveUser(u *User) error {
	return RemoveUser(s.database, u)
}

func (s *Storage) UpdateUser(u *User) error {
	return UpdateUser(s.database, u)
}

/*
	Password operations
*/
func (s *Storage) CreatePassword(u *User, st string, p *Password) error {
	return CreatePassword(s.database, u, p)
}

func (s *Storage) GetPassword(user *User, url string, username string) (*Password, error) {
	return QueryPassword(s.database, user, url, username)
}

func (s *Storage) UpdatePassword(u *User, st string, p *Password) error {
	return UpdatePassword(s.database, u, p)
}

func (s *Storage) DeletePassword(user *User, url string, username string) error {
	return DeletePassword(s.database, url, username, string(user.Uuid))
}

func (s *Storage) GetPasswords(u *User) ([]*Password, error){
	passwords, err := QueryAllPasswords(s.database, u)
	if err != nil{
		return make([]*Password, 0), nil
	}
	return passwords, nil
}
