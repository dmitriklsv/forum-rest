package sqlite3

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	users_table = `CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"email" TEXT,
		"username" TEXT,
		"password" TEXT
	);`

	sessions_table = `CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER,
		"session_token" TEXT,
		"expire_time" DATETIME,
		FOREIGN KEY (user_id) REFERENCES users (id)
	);`

	posts_table = `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER,
		"title" TEXT,
		"text" TEXT,
		FOREIGN KEY (user_id) REFERENCES users (id)
	);`

	categories_table = `CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"post_id" INTEGER,
		"name" TEXT,
		FOREIGN KEY (post_id) REFERENCES posts (id)
	);`

	// postcategory_table = `CREATE TABLE IF NOT EXISTS postcategory (
	// 	"post_id" INTEGER,
	// 	"category_id" INTEGER,
	// 	FOREIGN KEY (post_id) REFERENCES posts (id),
	// 	FOREIGN KEY (category_id) REFERENCES categories (id)
	// );`

	comment_table = `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER,
		"post_id" INTEGER,
		"text" TEXT,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (post_id) REFERENCES posts (id)
	);`

	post_reaction_table = `CREATE TABLE IF NOT EXISTS post_reactions (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"post_id" INTEGER,
		"user_id" INTEGER,
		"reaction" INTEGER,
		FOREIGN KEY (post_id) REFERENCES posts (id),
		FOREIGN KEY (user_id) REFERENCES users (id)
	);`

	comment_reaction_table = `CREATE TABLE IF NOT EXISTS comment_reactions (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"comment_id" INTEGER,
		"user_id" INTEGER,
		"reaction" INTEGER,
		FOREIGN KEY (comment_id) REFERENCES comments (id),
		FOREIGN KEY (user_id) REFERENCES users (id)
	);`
)

type DB struct {
	Collection *sql.DB
}

func Connect() (*DB, error) {
	log.Println("| | opening database...")
	db, err := sql.Open("sqlite3", "file:./database/database.db?_auth&_auth_user=admin&_auth_pass=admin&_auth_crypt=sha1")
	if err != nil {
		return nil, fmt.Errorf("couldn't open database due to %v", err)
	}
	log.Println("| | preparing database tables...")
	if _, err := prepareTables(db); err != nil {
		return nil, err
	}
	log.Println("| | database check done!")

	return &DB{
		db,
	}, nil
}

func prepareTables(db *sql.DB) (sql.Result, error) {
	log.Println("| | | checking database for existing tables...")

	arr := []string{users_table, sessions_table, posts_table, categories_table, comment_table, post_reaction_table, comment_reaction_table}

	for i := 0; i < len(arr); i++ {
		st, err := db.Prepare(arr[i])
		if err != nil {
			return nil, fmt.Errorf("couldn't create new table due to %v", err)
		}
		if i == len(arr)-1 {
			return st.Exec()
		}
		st.Exec()
	}

	// st, err := db.Prepare(fmt.Sprint(users_table + "\n" + sessions_table + "\n" + posts_table + "\n" + categories_table + "\n" + postcategory_table))
	// if err != nil {
	// 	return nil, err
	// }
	// return st.Exec()
	return nil, nil
}
