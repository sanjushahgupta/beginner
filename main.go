package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbname := "task.db"

	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("open", err)

	}
	defer db.Close()

	Sqlstmt := `create table Tabletask1 (Name TEXT, Age INT,Email TEXT);
	create table Newtabletask1 (Name TEXT,Email TEXT)`

	_, err = db.Exec(Sqlstmt)
	if err != nil {
		fmt.Printf("somethingwrong")
	}

	insertStatement := `
INSERT INTO Tabletask1 (Name,Age,Email) VALUES
("Amit",24,"amit@gmail.com")`

	_, err = db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}
