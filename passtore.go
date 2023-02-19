package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide host port username password db")
		return
	}

	host := arguments[1]
	p := arguments[2]
	username := arguments[3]
	password := arguments[4]
	dbname := arguments[5]

	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number", p, err)
		return
	}
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Println("Open connection error:", err)
		return
	}

	defer db.Close()

	// list all databases
	rows, err := db.Query(`select "datname" from "pg_database" where datistemplate=false`)
	if err != nil {
		fmt.Println("list databases error:", err)
		return
	}

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan error:", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()

	// list tables
	query := `select table_name from information_schema.tables where table_schema='public' order by table_name`
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println("Query error", err)
		return
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan error:", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
