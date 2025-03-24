package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

func InitializeDatabase() {
	var db *sql.DB
	databaseConnection := "user=expense_user password=password123 dbname=expense_tracker sslmode=disable"

	db, err := sql.Open("postgres", databaseConnection)

	fmt.Println("PostgreSQL Database Opened")

	if err != nil {
		log.Fatalf("Error connecting to database : %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Succesfully connected to the database")

	applyMigrations(db)
}

func applyMigrations(db *sql.DB) {
	migrationQuery := `
    CREATE TABLE IF NOT EXISTS "user" (
        id_user SERIAL PRIMARY KEY,
        name_user VARCHAR(50) NOT NULL,
        email_user VARCHAR(100) UNIQUE NOT NULL,
        password_user VARCHAR(255) NOT NULL,
        created_at_user TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE IF NOT EXISTS "category" (
        id_category SERIAL PRIMARY KEY,
        name_category VARCHAR(100) NOT NULL,
        description_category TEXT
    );

    CREATE TABLE IF NOT EXISTS "expense" (
        id_expense SERIAL PRIMARY KEY,
        amount DECIMAL(10, 2) NOT NULL,
        date_expense DATE NOT NULL, 
        description_expense TEXT,
        id_user INT NOT NULL,
        id_category INT,
        FOREIGN KEY (id_user) REFERENCES "user"(id_user),
        FOREIGN KEY (id_category) REFERENCES "category"(id_category)
    );
    `

	_, err := db.Exec(migrationQuery)

	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
		return
	}

	fmt.Println("Migrations successfully done")
}
