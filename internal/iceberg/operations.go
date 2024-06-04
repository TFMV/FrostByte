package iceberg

import (
	"log"
)

// CreateTable creates a new table in Iceberg.
func CreateTable(client *Client) error {
	err := client.CreateTable()
	if err != nil {
		log.Fatalf("Failed to create table in Iceberg: %v", err)
	}
	log.Println("Created Iceberg table.")
	return nil
}

// InsertData inserts data into an Iceberg table.
func InsertData(client *Client) error {
	err := client.InsertData()
	if err != nil {
		log.Fatalf("Failed to insert data into Iceberg: %v", err)
	}
	log.Println("Inserted data into Iceberg table.")
	return nil
}
