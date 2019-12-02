package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"webauthn/webauthn"
)

// TODO: enter credentials for database / read credentials from file
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "keycloud"	// alter to "KeyCloud"
)

func connectDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	// test database connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateUser(db *sql.DB, user *User) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("INSERT INTO users (uuid, name, mail, masterpasswd, createdate) VALUES($1,$2,$3,$4,CURRENT_TIMESTAMP)")
	if err != nil {
		return err
	}
	// execute statement
	user.Uuid = []byte(newUUID())
	_, err = stmt.Exec(user.Uuid, user.Name, user.Mail, user.MasterPassword)
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	// close connection and connection once query is executed
	defer stmt.Close()
	return err
}

func RemoveUser(db *sql.DB, user *User) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("DELETE FROM users WHERE uuid = $1 AND masterpasswd = $2")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(user.Uuid, user.MasterPassword)
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	// close connection and connection once query is executed
	return err
}

func UpdateUser(db *sql.DB, user *User) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("UPDATE users SET name = $1, mail = $2, masterpasswd = $3 WHERE uuid = $4")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(user.Name, user.Mail, user.MasterPassword, user.Uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func CreatePassword(db *sql.DB, user *User, p *Password) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("INSERT INTO passwds (uuid, url, passwd, username) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(user.Uuid, p.Url, p.Password, p.Username)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func QueryPassword(db *sql.DB, user *User, url string, username string) (password *Password, err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// prepare statement
	stmt, err := db.Prepare("SELECT entryid, url, passwd, username FROM passwds WHERE uuid = $1 AND url = $2 AND username = $3")
	if err != nil {
		return nil, err
	}
	// execute statement
	row := stmt.QueryRow(user.Uuid, url, username)
	// end query
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	// close connection and connection once query is executed
	defer stmt.Close()
	password = &Password{}
	err = row.Scan(&password.Id, &password.Url, &password.Password, &password.Username)
	if err != nil {
		return nil, err
	}
	return
}

func QuerySessionForUser(db *sql.DB, user *User) (token []byte, err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// prepare statement
	stmt, err := db.Prepare("SELECT session_token FROM sessions WHERE uuid = $1")
	if err != nil {
		return nil, err
	}
	// execute statement
	row := stmt.QueryRow(user.Uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	// end query
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	err = row.Scan(&token)
	if err != nil {
		return nil, err
	}
	return
}

func UpdatePassword(db *sql.DB, user *User, p *Password) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("UPDATE passwds SET url = $1, passwd = $2 WHERE uuid = $3 AND url = $4")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(p.Url, p.Password, user.Uuid, p.Url)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func DeletePassword(db *sql.DB, url string, username string, uuid string) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("DELETE FROM passwds WHERE uuid = $1 AND url = $2 AND username = $3")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(uuid, url, username)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func QueryUser(db *sql.DB, uuid string) (user *User, err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// prepare statement
	stmt, err := db.Prepare("SELECT uuid, name, mail, masterpasswd FROM users WHERE uuid = $1")
	if err != nil {
		return nil, err
	}
	// execute statement
	row := stmt.QueryRow(uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	// end query
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	user = &User{}
	err = row.Scan(&user.Uuid, &user.Name, &user.Mail, &user.MasterPassword)
	return user, nil
}

func QueryUserByName(db *sql.DB, name string) (user *User, err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// prepare statement
	stmt, err := db.Prepare("SELECT uuid, name, mail, masterpasswd FROM users WHERE name = $1")
	if err != nil {
		return nil, err
	}
	// execute statement
	row := stmt.QueryRow(name)
	// close connection and connection once query is executed
	defer stmt.Close()
	// end query
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	user = &User{}
	err = row.Scan(&user.Uuid, &user.Name, &user.Mail, &user.MasterPassword)
	return user, nil
}
func UpdateOrCreateSessionKeyForUser(db *sql.DB, u *User, token []byte) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("INSERT INTO sessions VALUES ($1, $2) ON CONFLICT ON CONSTRAINT sessions_pk DO UPDATE SET session_token = $3 WHERE sessions.uuid = $4")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(u.Uuid, token, token, u.Uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func DeleteSessionKeyForUser(db *sql.DB, u *User) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("DELETE FROM sessions WHERE uuid = $1")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(u.Uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func ClearAllSessionKeys(db *sql.DB) (err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("DELETE FROM sessions")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec()
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func QueryAllPasswords(db *sql.DB, u *User) (passwords []*Password, err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// prepare statement
	stmt, err := db.Prepare("SELECT entryid, url, passwd, username FROM passwds WHERE uuid = $1")
	if err != nil {
		return nil, err
	}
	// execute statement
	rows, err := stmt.Query(u.Uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		psw := Password{}
		err = rows.Scan(&psw.Id, &psw.Url, &psw.Password, &psw.Username)
		passwords = append(passwords, &psw)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return
}

func CreateAuthenticator(db *sql.DB, auth *Authenticator, uuid []byte) error {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// prepare statement
	stmt, err := db.Prepare("INSERT INTO authenticators VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	// execute statement
	_, err = stmt.Exec(auth.ID, auth.CredentialID, auth.PublicKey, auth.AAGUID, auth.SignCount, uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return err
	}
	// end query
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func QueryAuthenticator(db *sql.DB, id []byte) (auth *Authenticator, err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// prepare statement
	stmt, err := db.Prepare("SELECT id, credentialid, publickey, aaguid, signcount FROM authenticators WHERE id = $1")
	if err != nil {
		return nil, err
	}
	// execute statement
	row := stmt.QueryRow(id)
	// close connection and connection once query is executed
	defer stmt.Close()
	// end query
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	auth = &Authenticator{}
	err = row.Scan(&auth.ID, &auth.CredentialID, &auth.PublicKey, &auth.AAGUID, &auth.SignCount)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func QueryAllAuthenticators(db *sql.DB, uuid []byte) (auths []webauthn.Authenticator, err error) {
	// begin new statement
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// prepare statement
	stmt, err := db.Prepare("SELECT id, credentialid, publickey, aaguid, signcount FROM authenticators WHERE userid = $1")
	if err != nil {
		return nil, err
	}
	// execute statement
	rows, err := stmt.Query(uuid)
	// close connection and connection once query is executed
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		auth := &Authenticator{}
		err = rows.Scan(&auth.ID, &auth.CredentialID, &auth.PublicKey, &auth.AAGUID, &auth.SignCount)
		if err != nil {
			return nil, err
		}
		auths = append(auths, auth)
	}
	defer rows.Close()
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return
}
