package main

import (
	"encoding/hex"
	"fmt"
	"webauthn/webauthn"
)

type Storage struct {
	users          map[string]*User
	authenticators map[string]*Authenticator
}

func (s *Storage) AddAuthenticator(user webauthn.User, authenticator webauthn.Authenticator) error {
	authr := &Authenticator{
		ID:           authenticator.WebAuthID(),
		CredentialID: authenticator.WebAuthCredentialID(),
		PublicKey:    authenticator.WebAuthPublicKey(),
		AAGUID:       authenticator.WebAuthAAGUID(),
		SignCount:    authenticator.WebAuthSignCount(),
	}
	key := hex.EncodeToString(authr.ID)

	u, ok := s.users[string(user.WebAuthID())]
	if !ok {
		return fmt.Errorf("user not found")
	}

	if _, ok := s.authenticators[key]; ok {
		return fmt.Errorf("authenticator already exists")
	}

	authr.User = u

	u.Authenticators[key] = authr
	s.authenticators[key] = authr

	return nil
}

func (s *Storage) GetAuthenticator(id []byte) (webauthn.Authenticator, error) {
	authr, ok := s.authenticators[hex.EncodeToString(id)]
	if !ok {
		return nil, fmt.Errorf("authenticator not found")
	}
	return authr, nil
}

func (s *Storage) GetUser(webauthnID string) (user *User) {
	u, ok := s.users[webauthnID]
	if !ok {
		return nil
	} else {
		return u
	}
}

func (s *Storage) GetAuthenticators(user webauthn.User) ([]webauthn.Authenticator, error) {
	u, ok := s.users[string(user.WebAuthID())]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	var authrs []webauthn.Authenticator
	for _, v := range u.Authenticators {
		authrs = append(authrs, v)
	}
	return authrs, nil
}

/*
DUMMY IMPLEMENTATION
*/
func (s *Storage) GetSessionKeyForUser(user *User) []byte {
	return []byte("abc")
}
func (s *Storage) SetSessionKeyForUser(user *User, b []byte) error {
	return nil
}
func (s *Storage) DeleteSessionKeyForUser(user *User) error {
	return nil
}
func (s *Storage) AddUser(u *User) error {
	s.users[string(u.WebAuthID())] = u
	return nil
}
func (s *Storage) RemoveUser(u *User) error {
	return nil
}
func (s *Storage) UpdateUser(u *User) error {
	return nil
}

func (s *Storage) AddPassword(u *User, st string, p *Password) error {
	return nil
}

func (s *Storage) GetPassword(u *User, st string) (*Password, error) {
	return &Password{
		Password: "test",
		Id:       "test",
	}, nil
}

func (s *Storage) UpdatePassword(u *User, st string, p *Password) error {
	return nil
}
func (s *Storage) DeletePassword(u *User, st string) error {
	return nil
}
