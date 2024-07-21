package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveMessage(content string) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO messages(content, processed) VALUES($1, $2) RETURNING id", content, false).Scan(&id)
	return id, err
}

func MarkAsProcessed(id int64) error {
	_, err := db.Exec("UPDATE messages SET processed = TRUE WHERE id = $1", id)
	return err
}

func GetStatistics() (int, int, error) {
	var total, processed int
	err := db.QueryRow("SELECT COUNT(*), SUM(CASE WHEN processed THEN 1 ELSE 0 END) FROM messages").Scan(&total, &processed)
	return total, processed, err
}
