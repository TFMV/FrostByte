package polaris

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateTable(db *sqlx.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS sample_table (id INT PRIMARY KEY, name TEXT)")
	if err != nil {
		return err
	}
	log.Println("Created Polaris table.")
	return nil
}

func InsertData(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO sample_table (id, name) VALUES (1, 'Alice'), (2, 'Bob')")
	if err != nil {
		return err
	}
	log.Println("Inserted data into Polaris table.")
	return nil
}
