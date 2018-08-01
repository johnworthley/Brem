package data

import (
	"log"
	"errors"
)

// Auditor structure represents brem ico auditor
type Auditor struct {
	ID      int
	Address string `json:"address" binding:"required"`
}

// AddAuditor insert auditor in auditors table
func (auditor *Auditor) AddAuditor() (err error) {
	statement := "INSERT INTO auditors (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(auditor.Address).Scan(&auditor.ID)
	return
}

// GetAuditors returns auditor by address
func (auditor *Auditor) GetAuditor() (err error) {
	row, err := db.Query("SELECT * FROM auditors WHERE address = $1", auditor.Address)
	if err != nil {
		log.Println(err)
		return
	}
	exists := row.Next()
	if !exists {
		return errors.New("ICO doesn't exists")
	}
	err = row.Scan(&auditor.ID, &auditor.Address)
	return
}

//GetAllAuditors making query and return all auditors addresses
func GetAllAuditors() (auditors []Auditor, err error) {
	rows, err := db.Query("SELECT * FROM auditors")
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
