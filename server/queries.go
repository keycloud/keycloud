package main

import (
	"database/sql"
	"log"
)

func queryUuidFromUsername(db *sql.DB, username string) sql.Result{
	// begin new statement
	log.Println("queryUuidFromUsername: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryUuidFromUsername: prepare statement")
	stmt, err := db.Prepare("SELECT uuid FROM users WHERE name = $1")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryUuidFromUsername: execute statement")
	rs, err := stmt.Exec(username)
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryUuidFromUsername: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryUuidFromUsername: defer")
	defer stmt.Close()
	defer db.Close()

	return rs
}

func queryAddUser(db *sql.DB, user *User) {
	// begin new statement
	log.Println("queryAddUser: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryAddUser: prepare statement")
	stmt, err := db.Prepare("INSERT INTO users (uuid, name, email, masterpasswd, createdat) VALUES($1,$2,$3,$4,CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryAddUser: execute statement")
	_, err = stmt.Exec(newUUID(), user.Name, user.Mail, user.MasterPassword)
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryAddUser: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryAddUser: defer")
	defer stmt.Close()
	defer db.Close()
}

func queryRemoveUser(db *sql.DB, user *User) {
	// begin new statement
	log.Println("queryRemoveUser: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryRemoveUser: prepare statement")
	stmt, err := db.Prepare("DELETE FROM TABLE users WHERE uuid = $1 AND masterpasswd = $2")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryRemoveUser: execute statement")
	_, err = stmt.Exec(user.Name, user.MasterPassword)
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryRemoveUser: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryRemoveUser: defer")
	defer stmt.Close()
	defer db.Close()
}

func queryUpdateUser(db *sql.DB, user *User) {
	// begin new statement
	log.Println("queryUpdateUser: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryUpdateUser: prepare statement")
	stmt, err := db.Prepare("UPDATE TABLE users SET name = $1, mail = $2, masterpasswd = $3 ")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryUpdateUser: execute statement")
	// TODO for each None in [Name, Mail, MasterPassword]: use current value
	_, err = stmt.Exec(user.Name, user.Mail, user.MasterPassword)
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryUpdateUser: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryUpdateUser: defer")
	defer stmt.Close()
	defer db.Close()
}

func queryAddPassword(db *sql.DB, user *User, p *Password) {
	// begin new statement
	log.Println("queryAddPassword: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryAddPassword: prepare statement")
	stmt, err := db.Prepare("INSERT INTO passwds (uuid, url, passwd) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryAddPassword: execute statement")
	_, err = stmt.Exec(getUUID(user.Name), p.Url, p.Password)
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryAddPassword: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryAddPassword: defer")
	defer stmt.Close()
	defer db.Close()
}

func queryGetPassword(db *sql.DB, user *User, url string) sql.Result {
	// begin new statement
	log.Println("queryGetPassword: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryGetPassword: prepare statement")
	stmt, err := db.Prepare("SELECT passwds.passwd WHERE uuid = $1 AND url = $2")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryGetPassword: execute statement")
	rs, err := stmt.Exec(getUUID(user.Name), url)
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryGetPassword: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryGetPassword: defer")
	defer stmt.Close()
	defer db.Close()

	return rs
}

func queryUpdatePassword(db *sql.DB, user *User, p *Password) {
	// begin new statement
	log.Println("queryUpdatePassword: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryUpdatePassword: prepare statement")
	stmt, err := db.Prepare("UPDATE TABLE passwds SET url = $1, passwd = $2 WHERE uuid = $3")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryUpdatePassword: execute statement")
	// TODO for each None in [Url, Passwd]: use current value
	_, err = stmt.Exec(p.Url, p.Password, getUUID(user.Name))
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryUpdatePassword: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryUpdatePassword: defer")
	defer stmt.Close()
	defer db.Close()
}

func queryDeletePassword(db *sql.DB, p *Password) {
	// begin new statement
	log.Println("queryDeletePassword: begin new statement")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare statement
	log.Println("queryDeletePassword: prepare statement")
	stmt, err := db.Prepare("DELETE FROM TABLE passwds WHERE entryid = $1")
	if err != nil {
		log.Fatal(err)
	}

	// execute statement
	log.Println("queryDeletePassword: execute statement")
	_, err = stmt.Exec(p.Id)
	if err != nil {
		log.Fatal(err)
	}

	// end query
	log.Println("queryDeletePassword: end query statement")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// close connection and connection once query is executed
	log.Println("queryDeletePassword: defer")
	defer stmt.Close()
	defer db.Close()
}