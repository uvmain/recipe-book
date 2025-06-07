package database

import (
	"database/sql"
	"log"
	"time"
)

func CreateSessionsTable() {
	query := `CREATE TABLE IF NOT EXISTS sessions (
		token TEXT PRIMARY KEY,
		expires DATETIME
	);`

	checkQuery := "SELECT name FROM sqlite_master WHERE type='table' AND name='sessions'"
	checkError := Database.QueryRow(checkQuery).Scan()

	if checkError == nil {
		log.Println("recipes table already exists")
	} else {
		_, err := Database.Exec(query)
		if err != nil {
			log.Printf("Error creating sessions table: %s", err)
		} else {
			log.Println("sessions table created")
		}
	}
}

func SaveSessionToken(token string, duration time.Duration) error {
	expiresAt := time.Now().Add(duration) // Set expiration time
	query := `INSERT INTO sessions (token, expires) VALUES (?, ?)`
	_, err := Database.Exec(query, token, expiresAt)
	if err != nil {
		log.Printf("Error saving session token: %s", err)
		return err
	}
	return nil
}

func IsSessionValid(token string) bool {
	var expiresAt time.Time
	query := `SELECT expires FROM sessions WHERE token = ?`
	err := Database.QueryRow(query, token).Scan(&expiresAt)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error checking session token: %s", err)
		}
		return false
	}
	return time.Now().Before(expiresAt)
}

func DeleteSessionToken(token string) error {
	query := `DELETE FROM sessions WHERE token = ?`
	_, err := Database.Exec(query, token)
	if err != nil {
		log.Printf("Error deleting session token: %s", err)
		return err
	}
	return nil
}

func CleanupExpiredSessions() {
	query := `DELETE FROM sessions WHERE expires < ?`
	_, err := Database.Exec(query, time.Now())
	if err != nil {
		log.Printf("Error cleaning up expired sessions: %s", err)
	}
}

func StartSessionCleanupRoutine() {
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			CleanupExpiredSessions()
		}
	}()
}
