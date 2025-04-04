package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

var DB *sql.DB

func InitializeDatabaseConnection() {
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

	DB, err = sql.Open("postgres", databaseConnection)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("PostgreSQL Database Opened...")

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Succesfully connected to the database...")

	// applyMigrations(db)
}

// func applyMigrations(db *sql.DB) {

// 		creationQuery := `
// 		CREATE TABLE IF NOT EXISTS user (
// 	    id_user SERIAL PRIMARY KEY,
// 	    name_user VARCHAR(50) NOT NULL,
// 	    email_user VARCHAR(100) UNIQUE NOT NULL,
// 	    password_user VARCHAR(255) NOT NULL,
// 	    created_at_user TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 	);

// 	CREATE TABLE IF NOT EXISTS category (
// 	    id_category SERIAL PRIMARY KEY,
// 	    nom_category VARCHAR(100) NOT NULL,
// 	    description_category TEXT
// 	);

// 	CREATE TABLE IF NOT EXISTS expense (
// 	    id_expense SERIAL PRIMARY KEY,
// 	    description_expense TEXT,
// 	    amount DECIMAL(10, 2) NOT NULL,
// 	    date_expense DATE NOT NULL,
// 	    created_at_expense TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	    updated_at_expense TIMESTAMP,
// 	    id_user INT NOT NULL,
// 	    id_category INT,
// 	    FOREIGN KEY (id_user) REFERENCES user(id_user),
// 	    FOREIGN KEY (id_category) REFERENCES category(id_category)
// 	);`

// 	migrationQuery := `

// 	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

// 	ALTER TABLE "users" ADD COLUMN id_user_uuid UUID DEFAULT uuid_generate_v4();

// 	ALTER TABLE expense ADD COLUMN id_user_uuid UUID NOT NULL;

// 	ALTER TABLE expense ADD CONSTRAINT fk_id_user_uuid FOREIGN KEY (id_user_uuid) REFERENCES users(id_user_uuid);

// 	ALTER TABLE "users" DROP COLUMN id_user;

// 	ALTER TABLE "users" ADD CONSTRAINT pk_id_user_uuid PRIMARY KEY(id_user_uuid);

//     `

// 	_, err := db.Exec(migrationQuery)

// 	if err != nil {
// 		log.Fatalf("Error applying migrations: %v", err)
// 	}

// 	fmt.Println("Migrations applied successfully")
// }
