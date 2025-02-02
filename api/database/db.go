package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"recipebook/logic"

	_ "modernc.org/sqlite"
)

var Database *sql.DB

func GetDB() *sql.DB {
	if Database == nil {
		Initialise()
	}
	return Database
}

func Initialise() *sql.DB {
	dbPath := filepath.Join(logic.DataDirectory, "sqlite.db")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Println("Creating database file")

		file, err := os.Create(dbPath)
		if err != nil {
			log.Printf("Error creating database file: %s", err)
			return nil
		} else {
			log.Println("Database file created")
		}
		file.Close()
	} else {
		log.Println("Database already exists")
	}

	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		log.Printf("Error opening database file: %s", err)
		return nil
	} else {
		log.Println("Database file opened")
	}

	Database = db

	_, err = db.Exec("pragma journal_mode = wal;")
	if err != nil {
		log.Printf("Error entering WAL mode: %s", err)
	} else {
		log.Println("Database is in WAL mode")
	}

	createRecipesTable(db)
	return db
}

func createRecipesTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS recipes (
		slug TEXT PRIMARY KEY,
		dateCreated DATETIME,
		name TEXT,
		author TEXT,
		source TEXT,
		course TEXT,
		country TEXT,
		vegetarian INTEGER DEFAULT 0,
		prepTime TEXT,
		cookingTime TEXT,
		calories TEXT,
		servings TEXT,
		ingredients TEXT,
		instructions TEXT,
		imageFilename TEXT,
		imageWidth INTEGER,
		imageHeight INTEGER
	);`

	checkQuery := "SELECT name FROM sqlite_master WHERE type='table' AND name='recipes'"
	var name string
	checkError := db.QueryRow(checkQuery).Scan(&name)

	if checkError == nil {
		log.Println("recipes table already exists")
	} else {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error creating recipes table: %s", err)
		} else {
			log.Println("recipes table created")
		}
	}
}
