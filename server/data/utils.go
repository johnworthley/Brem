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
