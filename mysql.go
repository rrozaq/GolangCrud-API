package main

/* import package 
	"database/sql untuk mengonekan ke database sql"
	"log untuk menampilkan log di terminal pada saat debug"
*/
import (
	"database/sql"
	"log"
)

// membuat fungsi connect()
// db, err := sql.Open("mysql", "user:password@tcp(server:3306)/database")
func connect() *sql.DB {
	db, err := sql.Open("mysql", "admin:rooz@tcp(localhost:3306)/amikom")

	// jika database error
	if err != nil {
		log.Fatal(err)
	}

	return db
}
