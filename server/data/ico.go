package data

import (
	"database/sql"
	"github.com/pkg/errors"
	"strings"
	"time"

	"../../server/logger"
)

type ICO struct {
	ID           int       `json:"id"`
	Address      string    `json:"address" binding:"required"`
	Developer    Developer `json:"developer"`
	Description  string    `json:"description"`
	ClosingTime  time.Time `json:"closing_time"`
	FeePercent   int       `json:"fee_percent"`
	TokenAddress string    `json:"token_address"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	Status       string    `json:"status"`
	Location     string    `json:"location"`
	LocAddress   string    `json:"loc_address"`
}

const (
	CREATED   = "created"
	OPENED    = "opened"
	SUCCESS   = "success"
	FAILED    = "failed"
	WITHDRAWN = "withdrawn"
	OVERDUE   = "overdue"
)

// AddICO inserts new ICO to db
func (ico *ICO) AddICO() (err error) {
	ico.Address = strings.ToLower(ico.Address)
	ico.TokenAddress = strings.ToLower(ico.TokenAddress)
	statement := "INSERT INTO ico (address, developerID, closingTime, feePercent, tokenAddress, name, symbol, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(ico.Address, ico.Developer.ID, ico.ClosingTime, ico.FeePercent, ico.TokenAddress, ico.Name, ico.Symbol, "created").Scan(&ico.ID)
	return
}

// CreateICO from request
func (ico *ICO) CreateICO() (err error) {
	ico.Address = strings.ToLower(ico.Address)
	ico.TokenAddress = strings.ToLower(ico.TokenAddress)
	statement := "INSERT INTO ico (address, developerID, description, closingTime, feePercent, tokenAddress, name, symbol, status, location, locAddress) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(ico.Address, ico.Developer.ID, ico.Description, ico.ClosingTime, ico.FeePercent, ico.TokenAddress, ico.Name, ico.Symbol, CREATED, ico.Location, ico.LocAddress).Scan(&ico.ID)
	return
}

// GetICO returns ICO by ICO's address
func (ico *ICO) GetICO() (err error) {
	ico.Address = strings.ToLower(ico.Address)
	ico.TokenAddress = strings.ToLower(ico.TokenAddress)
	row, err := db.Query("SELECT * FROM ico WHERE address = $1", ico.Address)
	if err != nil {
		logger.Info(err)
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
	if err != nil {
		logger.Info(err)
	}
	if description_string.Valid {
		ico.Description = description_string.String
	}
	if location_string.Valid {
		ico.Location = location_string.String
	}
	if locAddress_string.Valid {
		ico.LocAddress = locAddress_string.String
	}
	err = ico.Developer.GetDeveloperById()
	if err != nil {
		logger.Info(err)
		return
	}
	return
}

const ICOs_AMOUNT_IN_PAGE = 6

// GetAllICOs returns ICOs in one page
func GetAllICOs(page int) (icos []ICO, err error) {
	start := ICOs_AMOUNT_IN_PAGE * page
	rows, err := db.Query("SELECT id, address, developerid, closingtime, tokenaddress, name, symbol, status, location, locaddress FROM ico  ORDER BY id DESC LIMIT $1 OFFSET $2",
		ICOs_AMOUNT_IN_PAGE, start)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		var location, locAddress sql.NullString
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location, &locAddress)
		if err != nil {
			logger.Info(err)
			return
		}
		if location.Valid {
			ico.Location = location.String
		}
		if locAddress.Valid {
			ico.LocAddress = locAddress.String
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

// GetCreatedICOs show all ICOs with status "created"
func GetCreatedICOs(page int) (icos []ICO, err error) {
	start := ICOs_AMOUNT_IN_PAGE * page
	rows, err := db.Query("SELECT id, address, developerid, closingtime, tokenaddress, name, symbol, status, location, locaddress FROM ico WHERE status = $1 ORDER BY id DESC LIMIT $2 OFFSET $3",
		CREATED, ICOs_AMOUNT_IN_PAGE, start)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		var location, locAddress sql.NullString
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location, &locAddress)
		if location.Valid {
			ico.Location = location.String
		}
		if locAddress.Valid {
			ico.LocAddress = locAddress.String
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

// GetOpenedICOs show all ICOs with status "opened"
func GetOpenedICOs(page int) (icos []ICO, err error) {
	start := ICOs_AMOUNT_IN_PAGE * page
	rows, err := db.Query("SELECT id, address, developerid, closingtime, tokenaddress, name, symbol, status, location, locaddress FROM ico WHERE status = $1 ORDER BY id DESC LIMIT $2 OFFSET $3",
		OPENED, ICOs_AMOUNT_IN_PAGE, start)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		var location, locAddress sql.NullString
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location, &locAddress)
		if location.Valid {
			ico.Location = location.String
		}
		if locAddress.Valid {
			ico.LocAddress = locAddress.String
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

// GetSuccessICOs show all ICOs with statuses "success" or "requested"
func GetSuccessICOs(page int) (icos []ICO, err error) {
	start := ICOs_AMOUNT_IN_PAGE * page
	rows, err := db.Query("SELECT id, address, developerid, closingtime, tokenaddress, name, symbol, status, location, locaddress FROM ico WHERE status = $1 ORDER BY id DESC LIMIT $2 OFFSET $3",
		SUCCESS, ICOs_AMOUNT_IN_PAGE, start)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		var location, locAddress sql.NullString
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location, &locAddress)
		if location.Valid {
			ico.Location = location.String
		}
		if locAddress.Valid {
			ico.LocAddress = locAddress.String
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

// GetFailedICOs show all ICOs with status "failed"
func GetFailedICOs(page int) (icos []ICO, err error) {
	start := ICOs_AMOUNT_IN_PAGE * page
	rows, err := db.Query("SELECT id, address, developerid, closingtime, tokenaddress, name, symbol, status, location, locaddress FROM ico WHERE status = $1 ORDER BY id DESC LIMIT $2 OFFSET $3",
		FAILED, ICOs_AMOUNT_IN_PAGE, start)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		var location, locAddress sql.NullString
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location, &locAddress)
		if location.Valid {
			ico.Location = location.String
		}
		if locAddress.Valid {
			ico.LocAddress = locAddress.String
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

// GetWithdrawnICOs show all ICOs with status "withdrawn"
func GetWithdrawnICOs(page int) (icos []ICO, err error) {
	start := ICOs_AMOUNT_IN_PAGE * page
	rows, err := db.Query("SELECT id, address, developerid, closingtime, tokenaddress, name, symbol, status, location, locaddress FROM ico WHERE status = $1 ORDER BY id DESC LIMIT $2 OFFSET $3",
		WITHDRAWN, ICOs_AMOUNT_IN_PAGE, start)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		var location, locAddress sql.NullString
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location, &locAddress)
		if location.Valid {
			ico.Location = location.String
		}
		if locAddress.Valid {
			ico.LocAddress = locAddress.String
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

// GetOverdueICOs show all ICOs with status "overdue"
func GetOverdueICOs(page int) (icos []ICO, err error) {
	start := ICOs_AMOUNT_IN_PAGE * page
	rows, err := db.Query("SELECT id, address, developerid, closingtime, tokenaddress, name, symbol, status, location, locaddress FROM ico WHERE status = $1 ORDER BY id DESC LIMIT $2 OFFSET $3",
		OVERDUE, ICOs_AMOUNT_IN_PAGE, start)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		var location, locAddress sql.NullString
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Developer.ID, &ico.ClosingTime, &ico.TokenAddress, &ico.Name, &ico.Symbol, &ico.Status, &location, &locAddress)
		if location.Valid {
			ico.Location = location.String
		}
		if locAddress.Valid {
			ico.LocAddress = locAddress.String
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

// AddAuditorToICO connect ICO id and auditor id
func (ico *ICO) AddAuditorToICO(auditor Auditor) (err error) {
	statement := "INSERT INTO icoAuditors (icoID, auditorID) VALUES ($1, $2)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(ico.ID, auditor.ID)
	return
}

// GetICOAuditors returns all current ICO auditors
func (ico *ICO) GetICOAuditors() (auditors []Auditor, err error) {
	rows, err := db.Query("SELECT * FROM auditors WHERE id IN "+
		"(SELECT auditorID FROM icoAuditors WHERE icoID = $1)", ico.ID)
	if err != nil {
		logger.Info(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var auditor Auditor
		err = rows.Scan(&auditor.ID, &auditor.Address)
		if err != nil {
			logger.Info(err)
			return
		}
		auditors = append(auditors, auditor)
	}
	return
}

// PublishICO change ICO status to "opened"
func (ico *ICO) PublishICO() (err error) {
	ico.Address = strings.ToLower(ico.Address)
	statement := "UPDATE ico SET status = $1 WHERE address = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(OPENED, ico.Address)
	return
}

// SetRequestedStatus change ICO status to "requested"
func (ico *ICO) SetRequestedStatus() (err error) {
	statement := "UPDATE ico SET status = $1 WHERE address = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
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
		logger.Info(err)
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
		logger.Info(err)
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
		logger.Info(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(status, ico.Address)
	return
}
