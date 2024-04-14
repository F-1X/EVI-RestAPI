package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresDB interface {
	Close(ctx context.Context) error
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

type PostgresClientWrapper struct {
	db *pgx.Conn
}

func (w *PostgresClientWrapper) Close(ctx context.Context) error {
	return w.db.Close(ctx)
}

func (w *PostgresClientWrapper) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return w.db.QueryRow(ctx, query, args...)
}

func (w *PostgresClientWrapper) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return w.db.Query(ctx, query, args...)
}

func InitPostgresClient(ctx context.Context, connString string) (PostgresDB, error) {
	conn, err := initPostgresClient(ctx, connString)
	if err != nil {
		return nil, err
	}

	return &PostgresClientWrapper{db: conn}, nil
}

func initPostgresClient(ctx context.Context, connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (w *PostgresClientWrapper) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return w.db.Exec(ctx, query, args...)
}
