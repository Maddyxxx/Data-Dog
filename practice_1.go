package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tester"
	password = "pivotal"
	dbname   = "testdb"
	sslmode  = "disable"
)

type Data struct {
	id   int
	text string
	date string
}

func WriteTodb(db *sql.DB, data Data) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	if _, err = tx.Exec("INSERT INTO test_table (f1, f2, f3) VALUES (?, ?, ?)",
		data.id, data.text, data.date); err != nil {
		return
	}
	return
}

func ReadAllData(db *sql.DB, sql_ string) {

	rows, err := db.Query(sql_)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var data []Data

	for rows.Next() {
		d := Data{}
		err := rows.Scan(&d.id, &d.text, &d.date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data = append(data, d)
	}
	for _, d := range data {
		fmt.Println(d.id, d.text, d.date)
	}
}

func main() {
	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbname, user, password, sslmode)
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Printf("Unable to connection to database: %v\n\n", err)
	}
	defer db.Close()

	data := Data{
		id:   2,
		text: "some text",
		date: "10-03-2019",
	}

	WriteTodb(db, data)
	ReadAllData(db, "select * from test_table")
}
