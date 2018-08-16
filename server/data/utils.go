package data

import "log"

func IsSuperuser(address string) (isSuperuser bool, err error) {
	row, err := db.Query("SELECT * FROM superuser WHERE address = $1", address)
	if err != nil {
		log.Println(err)
		return
	}
	defer row.Close()
	isSuperuser = row.Next()
	return
}

func ForcedRepairIcoTable(){
	statement := "DELETE FROM ico WHERE developerid NOT IN (SELECT id FROM developers)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	stmt.QueryRow()
	return
}
func ForcedRepairIcoAuditorsTable(){
	statement := "DELETE FROM icoauditors WHERE auditorid NOT IN (SELECT id FROM auditors) OR icoid NOT IN (SELECT id FROM ico)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	stmt.QueryRow()
	return
}
