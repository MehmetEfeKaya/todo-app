package database

import (
	"database/sql"
	"log"
	"yeni/backend/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./todos.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id TEXT PRIMARY KEY,
        title TEXT,
        completed BOOLEAN
    );
    `
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

func GetTodos() []models.Todo {
	rows, err := db.Query("SELECT id,title,completed FROM todos")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed)

		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)

	}
	return todos

}

func CreateTodo(todo models.Todo) {
	query := "INSERT INTO todos (id, title, completed) VALUES (?, ?, ?)"
	_, err := db.Exec(query, todo.ID, todo.Title, todo.Completed)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateTodo(id string, todo models.Todo) {
	query := "UPDATE todos SET id = ? title= ? completed = ?"
	_, err := db.Exec(query, todo.Title, todo.Completed, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTodo(id string) {
	query := "DELETE FROM todos WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
}
