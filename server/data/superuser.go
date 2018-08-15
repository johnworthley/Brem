package data

import (
"log"
"strings"
)

type Superuser struct {
	ID       int
	Address  string
}


func (superuser *Superuser) AddSuperuser() (err error) {
	superuser.Address = strings.ToLower(superuser.Address)
	statement := "INSERT INTO superuser (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(strings.ToLower(superuser.Address)).Scan(&superuser.ID)
	return
}

func (superuser *Superuser) GetSuperuser() (err error) {
	row, err := db.Query("SELECT * FROM superuser WHERE LOWER(address) = LOWER($1)", superuser.Address)
	if err != nil {
		log.Println(err)
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&superuser.ID, &superuser.Address)

	return
}

func ClearSuperuser() (err error){
	statement := "DELETE FROM superuser"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	stmt.QueryRow()
	return
}