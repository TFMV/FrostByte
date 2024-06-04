package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TFMV/FrostByte/internal/iceberg"
	"github.com/TFMV/FrostByte/internal/polaris"
	"github.com/TFMV/FrostByte/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	utils.InitLogger()

	config, err := utils.LoadConfig("/Users/thomasmcgeehan/FrostByte/FrostByte/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	polarisClient, err := polaris.NewClient(
		config.Polaris.Host,
		config.Polaris.Port,
		config.Polaris.User,
		config.Polaris.Password,
		config.Polaris.DbName,
	)
	if err != nil {
		log.Fatalf("Failed to create Polaris client: %v", err)
	}

	err = CreateTable(polarisClient)
	if err != nil {
		log.Fatalf("Failed to create table in Polaris: %v", err)
	}

	err = InsertData(polarisClient)
	if err != nil {
		log.Fatalf("Failed to insert data into Polaris: %v", err)
	}

	icebergClient := iceberg.NewClient()

	err = iceberg.CreateTable(icebergClient)
	if err != nil {
		log.Fatalf("Failed to create table in Iceberg: %v", err)
	}

	err = iceberg.InsertData(icebergClient)
	if err != nil {
		log.Fatalf("Failed to insert data into Iceberg: %v", err)
	}

	log.Println("Successfully showcased Iceberg and Polaris power!")
}

func CreateTable(pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS sample_table (
			id INT PRIMARY KEY,
			name TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	log.Println("Created Polaris table.")
	return nil
}

func InsertData(pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), `
		INSERT INTO sample_table (id, name) VALUES
		(1, 'Alice'),
		(2, 'Bob')
		ON CONFLICT (id) DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	log.Println("Inserted data into Polaris table.")
	return nil
}
