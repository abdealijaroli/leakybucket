package db

import (
	"log"

	_ "github.com/lib/pq"
)

func SeedDB(db *DB) error {
	err := CreateTables(db) 
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func CreateTables(db *DB) error {
	query := `CREATE TABLE IF NOT EXISTS scraped_data (
			id SERIAL PRIMARY KEY,
            link TEXT NOT NULL,
            raw_data TEXT,
            refined_data TEXT
		)
	`
	
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
