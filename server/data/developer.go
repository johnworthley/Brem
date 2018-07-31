package data

import (
	"log"
)

// Developer represents developer structure
type Developer struct {
	ID      int
	Address string `json:"address" binding:"required"`
}

// Create insert new developer address to db
func (dev *Developer) Create() (err error) {
	statement := "INSERT INTO developers (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(dev.Address).Scan(&dev.ID)
	return
}

// GetDeveloper select developer's id from by address
func (dev *Developer) GetDeveloper() (err error) {
	statement := "SELECT id FROM developers WHERE address = $1"
	row, err := db.Query(statement, dev.Address)
	row.Next()
	if err != nil {
		log.Println(err)
		return
	}
	err = row.Scan(&dev.ID)
	return
}