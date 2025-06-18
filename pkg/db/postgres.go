package db

import (
	"bookStore/pkg/logger"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"os"
)

type DB interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

type db struct {
	conn *pgx.Conn
}

func New(ctx context.Context, logger logger.Logger, dbURL string) DB {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		logger.Error("Unable to connect to database:" + err.Error())
		os.Exit(1)
	}
	return &db{conn: conn}
}

func (db *db) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.conn.QueryRow(ctx, query, args...)
}

func (db *db) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return db.conn.Query(ctx, query, args...)
}

func (db *db) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.conn.Exec(ctx, query, args...)
}
