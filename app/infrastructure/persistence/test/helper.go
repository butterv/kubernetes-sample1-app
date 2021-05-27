package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func DbMock(t *testing.T) (sqlmock.Sqlmock, *sqlx.DB) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New() = _, _, %#v; want nil", err)
	}

	sqlxDB := sqlx.NewDb(db, "mysql")

	return mock, sqlxDB
}

func TxMock(t *testing.T) (sqlmock.Sqlmock, *sqlx.Tx) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New() = _, _, %#v; want nil", err)
	}

	mock.ExpectBegin()
	sqlxTx := sqlx.NewDb(db, "mysql").MustBegin()

	return mock, sqlxTx
}
