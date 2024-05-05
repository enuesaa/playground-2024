package presenter

import (
	"fmt"
	"log"

	"github.com/enuesaa/txtsout/repository"
)

func HandleSqlite(filename string) error {
	db := repository.DBRepository{}
	if err := db.Open(filename); err != nil {
		return err
	}
	defer db.Close()
	res, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return err
	}
	for res.Next() {
		var tableName string
		if err := res.Scan(&tableName); err != nil {
			return err
		}
		log.Println(tableName)
		//TODO: DO NOT EMBED SQL. FIX THIS.
		rows, err := db.Query(fmt.Sprintf("SELECT id FROM %s;", tableName))
		if err != nil {
			return err
		}
		for rows.Next() {
			var record string
			if err := rows.Scan(&record); err != nil {
				return err
			}
			log.Println(record)
		}
	}
	
	return nil
}