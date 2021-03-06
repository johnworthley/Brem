package data

import "../../server/logger"

func ForcedRepairIcoTable() {
	statement := "DELETE FROM ico WHERE developerid NOT IN (SELECT id FROM developers)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	stmt.QueryRow()
	return
}

func ForcedRepairIcoAuditorsTable() {
	statement := "DELETE FROM icoauditors WHERE auditorid NOT IN (SELECT id FROM auditors) OR icoid NOT IN (SELECT id FROM ico)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		logger.Info(err)
		return
	}
	defer stmt.Close()
	stmt.QueryRow()
	return
}
