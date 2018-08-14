package data

import (
	"log"
	"github.com/pkg/errors"
	"time"
	"database/sql"
)

type ICO struct {
	ID 				int
	Address 		string		`json:"address" binding:"required"`
	Developer 		Developer	`json:"developer"`
	Description 	string
	ClosingTime 	time.Time
	FeePercent  	int
	TokenAddress	string
	Name			string
	Symbol			string
	Status			string
	Location		string
	LocAddress		string
}

// AddICO inserts new ICO to db
func (ico *ICO) AddICO() (err error) {
	statement := "INSERT INTO ico (address, developerID, closingTime, feePercent, tokenAddress, name, symbol, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(ico.Address, ico.Developer.ID, ico.ClosingTime, ico.FeePercent, ico.TokenAddress, ico.Name, ico.Symbol, "created").Scan(&ico.ID)
	return
}

// GetICO returns ICO by ICO's address
func (ico *ICO) GetICO() (err error) {
	row, err := db.Query("SELECT * FROM ico WHERE address = $1", ico.Address)
	if err != nil {
		log.Println(err)
		return
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		return errors.New("ICO doesn't exists")
	}
	var description_string sql.NullString
	var location_string sql.NullString
	var locAddress_string sql.NullString
	err = row.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &description_string, &ico.ClosingTime, &ico.FeePercent, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location_string, &locAddress_string)
	if description_string.Valid {
		ico.Description = description_string.String
	}
	if location_string.Valid {
		ico.Location = location_string.String
	}
	if locAddress_string.Valid {
		ico.LocAddress = locAddress_string.String
	}
	return
}

// GetCreatedICOs show all ICOs with status "created"
func GetCreatedICOs() (icos []ICO, err error)  {
	rows, err := db.Query("SELECT * FROM ico WHERE status = $1 ORDER BY id DESC", "created")
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

// GetOpenedICOs show all ICOs with status "opened"
func GetOpenedICOs() (icos []ICO, err error) {
	rows, err := db.Query("SELECT * FROM ico WHERE status = $1 ORDER BY id DESC", "opened")
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

// GetSuccessICOs show all ICOs with statuses "success" or "requested"
func GetSuccessICOs() (icos []ICO, err error) {
	rows, err := db.Query("SELECT * FROM ico WHERE status = $1 OR status = $2 ORDER BY id DESC", "success", "requested")
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

// GetFailedICOs show all ICOs with status "failed"
func GetFailedICOs() (icos []ICO, err error) {
	rows, err := db.Query("SELECT * FROM ico WHERE status = $1 ORDER BY id DESC", "failed")
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

// GetWithdrawnICOs show all ICOs with status "withdrawn"
func GetWithdrawnICOs() (icos []ICO, err error) {
	rows, err := db.Query("SELECT * FROM ico WHERE status = $1 ORDER BY id DESC", "withdrawn")
	defer rows.Close()
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
	rows, err := db.Query("SELECT * FROM auditors WHERE id IN " +
		"(SELECT auditorID FROM icoAuditors WHERE icoID = $1)", ico.ID)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
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

// SetRequestedStatus change ICO status to "requested"
func (ico *ICO) SetRequestedStatus() (err error) {
	statement := "UPDATE ico SET status = $1 WHERE address = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("requested", ico.Address)
	return
}

// SetSuccessStatus change ICO status to "success"
func (ico *ICO) SetSuccessStatus() (err error) {
	statement := "UPDATE ico SET status = $1 WHERE address = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("success", ico.Address)
	return
}

// SetWithdrawnStatus sets withdrawn status to current ICO
func (ico *ICO) SetWithdrawnStatus() (err error) {
	statement := "UPDATE ico SET status = $1 WHERE address = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("withdrawn", ico.Address)
	return
}
// SetStatusICO set ICO status
func (ico *ICO) SetStatusICO(status string) (err error) {
	statement := "UPDATE ico SET status = $1 WHERE address = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(status, ico.Address)
	return
}
