package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

//a successful case
func TestWriteTodb1(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := Data{
		id:    5,
		date:  "10-02-2020",
		name_: "some_random_name",
		descr: "some_random_text",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO test_table_37").
		WithArgs(data.id, data.date, data.name_, data.descr).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = WriteTodb(db, data); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// a successful case
func TestReadAllData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM test_table_37")
	mock.ExpectClose()
}
