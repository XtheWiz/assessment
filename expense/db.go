package expense

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error while connect to database", err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT []
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Cannot create table expenses... ", err)
	}
	fmt.Println("Exepense table init success")

	return db
}
