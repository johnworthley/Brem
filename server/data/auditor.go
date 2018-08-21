package data

import (
	"database/sql"
	"errors"
	"strings"

	"../../server/logger"
)

// Auditor structure represents brem ico auditor
type Auditor struct {
	ID       int    `json:"id"`
	Address  string `json:"address"`
	Username string `json:"username"`
}

// AddAuditor insert auditor in auditors table
func (auditor *Auditor) AddAuditor() (err error) {
	auditor.Address = strings.ToLower(auditor.Address)
	statement := "INSERT INTO auditors (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(strings.ToLower(auditor.Address)).Scan(&auditor.ID)
	return
}

// GetAuditors returns auditor by address
func (auditor *Auditor) GetAuditor() (err error) {
	row, err := db.Query("SELECT * FROM auditors WHERE LOWER(address) = LOWER($1)", auditor.Address)
	if err != nil {
		logger.Info(err)
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
	rows, err := db.Query("SELECT * FROM auditors")
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var auditor Auditor
		var username sql.NullString
		err = rows.Scan(&auditor.ID, &auditor.Address, &username)
		if err != nil {
			logger.Info(err)
			return
		}
		if username.Valid {
			auditor.Username = username.String
		}
		auditors = append(auditors, auditor)
	}
	return
}

// GetICOs returns current auditor ICOs
func (auditor *Auditor) GetICOs() (icos []ICO, err error) {
	rows, err := db.Query("SELECT id, address, developerID, closingTime, feePercent, tokenAddress, name, symbol, status FROM ico WHERE status = 'requested' AND"+
		" id IN (SELECT icoID FROM icoAuditors WHERE auditorID = $1)", auditor.ID)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.FeePercent, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status)
		if err != nil {
			logger.Info(err)
			return
		}
		err = ico.Developer.GetDeveloperById()
		if err != nil {
			logger.Info(err)
			return
		}
		icos = append(icos, ico)
	}
	return
}
