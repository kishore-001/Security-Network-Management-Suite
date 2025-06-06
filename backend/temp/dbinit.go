package dbinit

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func CreateStructure() {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:8500/snsmsdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        name VARCHAR(255) NOT NULL,
        role VARCHAR(50) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(100) NOT NULL
    );`

	_, err = db.ExecContext(context.Background(), createTable)
	if err != nil {
		log.Fatal("creating table:", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	insertUser := `
    INSERT INTO users (name, role, email, password_hash)
	SELECT $1::text, $2::text, $3::text, $4::text
    WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = $3);`

	_, err = db.ExecContext(context.Background(), insertUser, "admin", "admin", "admin@example.com", string(hashedPassword))
	if err != nil {
		log.Fatal("inserting user:", err)
	}

	fmt.Println("Initial setup complete.")
}
