package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

func InitializeDatabaseConnection() {
	var db *sql.DB
	//Loading environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v ", err)
	}

	fmt.Println(".env file successfully loaded...")

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	databaseConnection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)

	db, err = sql.Open("postgres", databaseConnection)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("PostgreSQL Database Opened...")

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Succesfully connected to the database...")

	applyMigrations(db)

	fmt.Println("migrations applied...")
}

func applyMigrations(db *sql.DB) {
	migrationQuery := `
    CREATE TABLE user (
    id_user SERIAL PRIMARY KEY,
    name_user VARCHAR(50) NOT NULL,
    email_user VARCHAR(100) UNIQUE NOT NULL,
    password_user VARCHAR(255) NOT NULL,
    created_at_user TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE category (
    id_category SERIAL PRIMARY KEY,
    nom_category VARCHAR(100) NOT NULL,
    description_category TEXT
);

CREATE TABLE expense (
    id_expense SERIAL PRIMARY KEY,
    description_expense TEXT,
    amount DECIMAL(10, 2) NOT NULL,
    date_expense DATE NOT NULL,
    created_at_expense TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at_expense TIMESTAMP, 
    id_user INT NOT NULL,
    id_category INT,
    FOREIGN KEY (id_user) REFERENCES user(id_user),
    FOREIGN KEY (id_category) REFERENCES category(id_category)
);
    `

	_, err := db.Exec(migrationQuery)

	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
		return
	}

	fmt.Println("Migrations successfully done")
}
