package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
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
	date  pgtype.Date
	name_ string
	descr string
}

func main() {

	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s target_session_attrs=read-write",
		host, port, dbname, user, password)

	pool, err := pgxpool.Connect(context.Background(), connstring)
	if err != nil {
		fmt.Printf("Unable to connection to database: %v\n\n", err)
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), "select * from test_table_37")

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	data := []Data{}

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
