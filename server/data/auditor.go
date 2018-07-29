package data

import "log"

// Auditor structure represents brem ico auditor
type Auditor struct {
	ID      int
	Address string `json:"address" binding:"required"`
}

// AddAuditor insert auditor in auditors table
func (auditor *Auditor) AddAuditor() (err error) {
	statement := "INSERT INTO auditor (address) VALUES ($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(auditor.Address).Scan(&auditor.ID)
	return
}
