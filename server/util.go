package main

import (
	"database/sql"
	uuid "github.com/nu7hatch/gouuid"
)

func newUUID() string {
	u, err := uuid.NewV4()
	// u is unique
	if err != nil {
		panic(err)
	}
	return u.String()
}

func getUUID(username string) sql.Result {
	uuid := queryUuidFromUsername(connectDatabase(), username)
	// TODO parse uuid from rs to string
	return uuid
}