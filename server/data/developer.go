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

func (dev *Developer) GetDeveloper() (err error) {
	statement := "SELECT * FROM developers WHERE address = $1"
	row, err := db.Query(statement, dev.Address)
	if err != nil {
		log.Println(err)
		return
	}
	err = row.Scan(&dev.ID)
	return
}