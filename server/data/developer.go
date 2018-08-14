package data

import (
	"log"
)

// Developer represents developer structure
type Developer struct {
	ID      	int
	Address 	string 	`json:"address"`
	Username	string 	`json:"username"`
	Sign		string	`json:"sign"`
}

// Check for existing of address
func (dev Developer) IsExists() (exists bool, err error)  {
	row, err := db.Query("SELECT * FROM developers WHERE address = $1", dev.Address)
	if err != nil {
		log.Println(err)
		return
	}
	exists = row.Next()
	return
}

// Create insert new developer address to db
func (dev *Developer) Create() (err error) {
	statement := "INSERT INTO developers (address, username) VALUES ($1, $2)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(dev.Address, dev.Username)
	return
}

// GetDeveloper select developer's id from by address
func (dev *Developer) GetDeveloper() (err error) {
	statement := "SELECT * FROM developers WHERE address = $1"
	row, err := db.Query(statement, dev.Address)
	if err != nil {
		log.Println(err)
		return
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&dev.ID, &dev.Address, &dev.Username)
	return
}

// GetICOs return all developer's icos
func (dev *Developer) GetICOs() (icos []ICO, err error)  {
	statement := "SELECT id, address, status FROM ico WHERE developerID = $1 ORDER BY id DESC"
	rows, err := db.Query(statement, dev.ID)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ico ICO
		err = rows.Scan(&ico.ID, &ico.Address, &ico.Status)
		if err != nil {
			log.Println(err)
			return
		}
		icos = append(icos, ico)
	}
	return
}