package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "tester"
	password = "pivotal"
	dbname   = "testdb"
)

type Data struct {
	id    int
	date  string
	name_ string
	descr string
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
	if _, err = tx.Exec("INSERT INTO test_table_37 (id, date, name_, descr) VALUES (?, ?, ?, ?)",
		data.id, data.date, data.name_, data.descr); err != nil {
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
		err := rows.Scan(&d.id, &d.date, &d.name_, &d.descr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data = append(data, d)
	}
	for _, d := range data {
		fmt.Println(d.id, d.date, d.name_, d.descr)
	}
}

func main() {
	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s target_session_attrs=read-write",
		host, port, dbname, user, password)
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Printf("Unable to connection to database: %v\n\n", err)
	}
	defer db.Close()

	data := Data{
		id:    2,
		date:  "10-03-2019",
		name_: "some_random_name",
		descr: "some_random_text",
	}

	WriteTodb(db, data)
	ReadAllData(db, "select * from test_table_37")
}
