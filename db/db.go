package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "api.db") 
    if err != nil {
        panic("can't connect to database")
    }

    DB.SetMaxOpenConns(10)
    DB.SetConnMaxIdleTime(5)

    CreateTable()
}
func CreateTable() {
	createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL unique,
        password  TEXT NOT NULL
        
    );`

    _, err := DB.Exec(createUserTable)
    if err != nil {
        log.Fatalf("can't create table: %v", err)
    }
    createEventTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER,
		foreign key(user_id) references users(id)
    );`

    _, err = DB.Exec(createEventTable)
    if err != nil {
        log.Fatalf("can't create table: %v", err)
    }


   createRegistrationTable :=`
   CREATE TABLE IF NOT EXISTS registration (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   event_id INTEGER,
   user_id INTEGER,
   FOREIGN KEY(event_id) REFERENCES events(id),
   FOREIGN KEY(user_id) REFERENCES users(id)
);

   `
  _, err = DB.Exec(createRegistrationTable)
  if err != nil{
	panic("cant create regis table")
  }

}
