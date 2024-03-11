package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Connect() *sql.DB {
	// this should live in a .env
	dsn := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	println("Database connected")

	// init tables
	usersTable(db)
	createTodoTable(db)

	println("Database initialized")
	return db
}

// Tables
func usersTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
		    			id SERIAL PRIMARY KEY,
		    			email TEXT NOT NULL,
		    			created_at TIMESTAMP DEFAULT NOW()
		)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// seed
	//adminUser := user.User{Email: "adminUser@adminUser.com"}
	//testUser := user.User{Email: "test@test.com"}
	//query = `SELECT id FROM users WHERE email = $1`
	//err = db.QueryRow(query, adminUser.Email).Scan(&adminUser.ID)
	//if err != nil {
	//	adminUser.ID = insertUser(db, adminUser.Email)
	//}
	//
	//err = db.QueryRow(query, testUser.Email).Scan(&testUser.ID)
	//if err != nil {
	//	testUser.ID = insertUser(db, testUser.Email)
	//}

	println("Table users created")
}

func createTodoTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS todos (
		    			id SERIAL PRIMARY KEY,
		    			task TEXT NOT NULL,
		    			status TEXT NOT NULL,
		    			user_id INT REFERENCES users(id),
		    			created_at TIMESTAMP DEFAULT NOW(),
						updated_at TIMESTAMP DEFAULT NOW()
		)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// seed
	//todos := []todo.Todo{
	//	{Task: "Groceries", Status: todo.StatusIncomplete},
	//	{Task: "Laundry", Status: todo.StatusInProgress},
	//	{Task: "Dishes", Status: todo.StatusComplete},
	//	{Task: "Retro meeting", Status: todo.StatusCancelled},
	//}
	//
	//for i, t := range todos {
	//	if i%2 != 0 {
	//		id := insertTodo(db, t.Task, t.Status, 1)
	//		println("Inserted todo with id:", id)
	//	} else {
	//		id := insertTodo(db, t.Task, t.Status, 2)
	//		println("Inserted todo with id:", id)
	//	}
	//}

	println("Table todos created")
}

func insertUser(db *sql.DB, email string) int {
	query := `INSERT INTO users (email) VALUES ($1) RETURNING id`

	var id int
	err := db.QueryRow(query, email).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func insertTodo(db *sql.DB, task string, status string, userID int) int {
	query := `INSERT INTO todos (task, status, user_id) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := db.QueryRow(query, task, status, userID).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}
