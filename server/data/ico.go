package data

import (
	"log"
	"github.com/pkg/errors"
)

type ICO struct {
	ID 			int
	Address 	string		`json:"address" binding:"required"`
	Developer 	Developer	`json:"developer"`
	Status		string
}

// AddICO inserts new ICO to db
func (ico *ICO) AddICO() (err error) {
	statement := "INSERT INTO ico (address, developerID, status) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(ico.Address, ico.Developer.ID, "created").Scan(&ico.ID)
	return
}

// GetICO returns ICO by ICO's address
func (ico *ICO) GetICO() (err error) {
	row, err := db.Query("SELECT * FROM ico WHERE address = $1", ico.Address)
	if err != nil {
		log.Println(err)
		return
	}
	exists := row.Next()
	if !exists {
		return errors.New("ICO doesn't exists")
	}
	err = row.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.Status)
	return
}

// GetCreatedICOs show all ICOs with status "created"
func GetCreatedICOs() (icos []ICO, err error)  {
	rows, err := db.Query("SELECT * FROM ico WHERE status = $1", "created")
	if err != nil {
		log.Println(err)
		return
	}
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

// AddAuditorToICO connect ICO id and auditor id
func (ico *ICO) AddAuditorToICO(auditor Auditor) (err error) {
	statement := "INSERT INTO icoAuditors (icoID, auditorID) VALUES ($1, $2)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(ico.ID, auditor.ID)
	return
}

// GetICOAuditors returns all current ICO auditors
func (ico *ICO) GetICOAuditors() (auditors []Auditor, err error) {
	rows, err := db.Query("SELECT * FROM auditors WHERE id = " +
		"(SELECT auditorID FROM icoAuditors WHERE icoID = $1)", ico.ID)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var auditor Auditor
		err = rows.Scan(&auditor.ID, &auditor.Address)
		if err != nil {
			log.Println(err)
			return
		}
		auditors = append(auditors, auditor)
	}
	return
}

// PublishICO change ICO status to "opened"
func (ico *ICO) PublishICO() (err error) {
	statement := "UPDATE ico SET status = $1 WHERE address = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("opened", ico.Address)
	return
}
