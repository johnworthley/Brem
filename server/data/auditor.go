package data

import (
	"log"
	"errors"
	"strings"
	"database/sql"
)

// Auditor structure represents brem ico auditor
type Auditor struct {
	ID      	int
	Address 	string
	Username	string
}

// TODO: Change
// AddAuditor insert auditor in auditors table
func (auditor *Auditor) AddAuditor() (err error) {
	statement := "INSERT INTO auditors (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(strings.ToLower(auditor.Address)).Scan(&auditor.ID)
	return
}

// TODO: Change
// GetAuditors returns auditor by address
func (auditor *Auditor) GetAuditor() (err error) {
	row, err := db.Query("SELECT * FROM auditors WHERE LOWER(address) = LOWER($1)", auditor.Address)
	if err != nil {
		log.Println(err)
		return
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		return errors.New("Auditor doesn't exists")
	}
	var username_string sql.NullString
	err = row.Scan(&auditor.ID, &auditor.Address, &username_string)
	if username_string.Valid {
		auditor.Username = username_string.String
	}
	return
}

//GetAllAuditors making query and return all auditors addresses (with usernames)
func GetAllAuditors() (auditors []Auditor, err error) {
	rows, err := db.Query("SELECT * FROM auditors WHERE username LENGTH(username) > 0")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var auditor Auditor
		err = rows.Scan(&auditor.ID, &auditor.Address, &auditor.Username)
		if err != nil {
			log.Println(err)
			return
		}
		auditors = append(auditors, auditor)
	}
	return
}

// GetICOs returns current auditor ICOs
func (auditor *Auditor) GetICOs() (icos []ICO, err error) {
	rows, err := db.Query("SELECT * FROM ico WHERE status = 'requested' AND" +
		" id IN (SELECT icoID FROM icoAuditors WHERE auditorID = $1)", auditor.ID)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.Status)
		if err != nil {
			log.Println(err)
			return
		}
		icos = append(icos, ico)
	}
	return
}
