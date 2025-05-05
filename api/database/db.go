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
		log.Printf("Database file opened: %s", dbPath)
	}

	Database = db

	_, err = db.Exec("pragma journal_mode = wal;")
	if err != nil {
		log.Printf("Error entering WAL mode: %s", err)
	} else {
		log.Println("Database is in WAL mode")
	}

	createRecipesTable()
	createFtsTable()
	insertFtsData()
	createFtsTriggers()
	CreateSessionsTable()
	StartSessionCleanupRoutine()
	return db
}

func createRecipesTable() {
	query := `CREATE TABLE IF NOT EXISTS recipes (
		slug TEXT PRIMARY KEY NOT NULL,
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
	checkError := Database.QueryRow(checkQuery).Scan()

	if checkError == nil {
		log.Println("recipes table already exists")
	} else {
		_, err := Database.Exec(query)
		if err != nil {
			log.Printf("Error creating recipes table: %s", err)
		} else {
			log.Println("recipes table created")
		}
	}
}

func createFtsTable() {
	query := `CREATE VIRTUAL TABLE IF NOT EXISTS recipes_fts USING fts5(
		slug,
		name,
		author,
		source,
		course,
		country,
		ingredients,
		instructions,
		tokenize="trigram"
	);`

	checkQuery := "SELECT name FROM sqlite_master WHERE type='table' AND name='recipes'"
	checkError := Database.QueryRow(checkQuery).Scan()

	if checkError == nil {
		log.Println("fts table already exists")
	} else {
		_, err := Database.Exec(query)
		if err != nil {
			log.Printf("Error creating fts table: %s", err)
		} else {
			log.Println("fts table created")
		}
	}
}

func insertFtsData() {
	query := `INSERT INTO recipes_fts (slug,  name, author, source, course, country, ingredients, instructions)
		SELECT slug, name, author, source, course, country, ingredients, instructions FROM recipes;`

	_, err := Database.Exec(query)
	if err != nil {
		log.Printf("Error inserting data into fts table: %s", err)
	} else {
		log.Println("Data inserted into fts table")
	}
}

func createFtsTriggers() {
	createTriggerIfNotExists("after_insert", `CREATE TRIGGER after_insert AFTER INSERT ON recipes
        BEGIN
            INSERT INTO recipes_fts (slug, name, author, source, course, country, ingredients, instructions)
            VALUES (new.slug, new.name, new.author, new.source, new.course, new.country, new.ingredients, new.instructions);
        END;`)

	createTriggerIfNotExists("after_delete", `CREATE TRIGGER after_delete AFTER DELETE ON recipes
    BEGIN
        DELETE FROM recipes_fts WHERE slug = old.slug;
    END;`)

	createTriggerIfNotExists("after_update", `CREATE TRIGGER after_update AFTER UPDATE ON recipes
    BEGIN
        UPDATE recipes_fts SET 
            name = new.name, 
            author = new.author, 
            source = new.source, 
            course = new.course, 
            country = new.country, 
            ingredients = new.ingredients, 
            instructions = new.instructions
        WHERE slug = old.slug;
    END;`)
}

func createTriggerIfNotExists(triggerName string, triggerSQL string) {
	checkQuery := "SELECT name FROM sqlite_master WHERE type='trigger' AND name=?"
	var existingTrigger string
	checkError := Database.QueryRow(checkQuery, triggerName).Scan(&existingTrigger)

	if checkError == nil {
		log.Printf("%s trigger already exists", triggerName)
	} else if checkError == sql.ErrNoRows {
		_, err := Database.Exec(triggerSQL)
		if err != nil {
			log.Printf("Error creating %s trigger: %s", triggerName, err)
		} else {
			log.Printf("%s trigger created", triggerName)
		}
	} else {
		log.Printf("Error checking for %s trigger: %s", triggerName, checkError)
	}
}
