package data

import (
	"log"
)

// Developer represents developer structure
type Developer struct {
	ID      int
	Address string `json:"address"`
}

// Create insert new developer address to db
func (dev *Developer) Create() (err error) {
	statement := "INSERT INTO developer (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(dev.Address).Scan(&dev.ID)
	return
}
