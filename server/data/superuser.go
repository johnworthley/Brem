package data

import (
	"strings"

	"../../server/logger"
)

type Superuser struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

func (superuser *Superuser) AddSuperuser() (err error) {
	superuser.Address = strings.ToLower(superuser.Address)
	statement := "INSERT INTO superuser (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(strings.ToLower(superuser.Address)).Scan(&superuser.ID)
	return
}

func (superuser *Superuser) GetSuperuser() (err error) {
	row, err := db.Query("SELECT * FROM superuser WHERE LOWER(address) = LOWER($1)", superuser.Address)
	if err != nil {
		logger.Info(err)
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&superuser.ID, &superuser.Address)

	return
}

func ClearSuperuser() (err error) {
	statement := "DELETE FROM superuser"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	stmt.QueryRow()
	return
}

func IsSuperuser(address string) (isSuperuser bool, err error) {
	row, err := db.Query("SELECT * FROM superuser WHERE address = $1", address)
	if err != nil {
		logger.Info(err)
		return
	}
	defer row.Close()
	isSuperuser = row.Next()
	return
}

func GetSuperuserICOs() (icos []ICO, err error) {
	rows, err := db.Query("SELECT id, address, closingTime, feePercent, tokenAddress, name, symbol, status, locAddress FROM ico WHERE status = 'created'")
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		err = rows.Scan(&ico.ID, &ico.Address, &ico.ClosingTime, &ico.FeePercent, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &ico.LocAddress)
		if err != nil {
			logger.Info(err)
			return
		}
		icos = append(icos, ico)
	}
	return
}