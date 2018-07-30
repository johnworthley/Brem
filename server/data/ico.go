package data

import "log"

type ICO struct {
	ID 			int
	Address 	string		`json:"ico_address" binding:"required"`
	Developer 	Developer
	Status		string
}

func (ico *ICO) AddICO() (err error) {
	statement := "INSERT INTO ico (address, developerID, status) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(ico.Address, ico.Developer.ID, "created").Scan(&ico.ID)
	return
}

func GetCreatedICOs() (icos []ICO, err error)  {
	rows, err := db.Query("SELECT * FROM ico WHERE status = $1")
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
