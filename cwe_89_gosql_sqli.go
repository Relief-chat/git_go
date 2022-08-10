package main

import "database/sql"
import "fmt"

func bad1() {
    db, err := sql.Open("mysql", "theUser:thePassword@/theDbName")
    if err != nil {
        panic(err)
    }
    query = "SELECT name FROM users WHERE age=" + req.FormValue("age")
    // bad
    db.Query(query)
}

func bad2(db *sql.DB) {
    query = "SELECT name FROM users WHERE age="
    query += req.FormValue("age")
    // bad
    db.QueryRow(query)
}

func bad3(db *sql.DB) {
    query = fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email)
    // bad
    db.Exec(query)
}

func bad4(db *sql.DB) {
    // bad
    db.Exec("SELECT name FROM users WHERE age=" + req.FormValue("age"))
}

func bad5(db *sql.DB) {
    // bad
    db.Exec(fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email))
}

func ok1(db *sql.DB) {
    query = fmt.Sprintf("SELECT * FROM users WHERE email=hello;")
    // ok 
    db.Exec(query)
}

func ok2(db *sql.DB) {
    query = "SELECT name FROM users WHERE age=" + "3"
    // ok 
    db.Query(query)
}

func ok3(db *sql.DB) {
    query = "SELECT name FROM users WHERE age="
    query += "3"
    // ok 
    db.Query(query)
}

func ok4(db *sql.DB) {
    // ok 
    db.Exec("INSERT INTO users(name, email) VALUES($1, $2)",
  "Jon Calhoun", "jon@calhoun.io")
}

func ok5(db *sql.DB) {
    // ok 
    db.Exec("SELECT name FROM users WHERE age=" + "3")
}

func ok6(db *sql.DB) {
    // ok 
    db.Exec(fmt.Sprintf("SELECT * FROM users WHERE email=hello;"))
}
