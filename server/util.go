package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"math/rand"
)

func newUUID() string {
	u, err := uuid.NewV4()
	// u is unique
	if err != nil {
		panic(err)
	}
	return u.String()
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GeneratePassword(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
