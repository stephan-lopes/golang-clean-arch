package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

// Wrapping do pgxpool para criar mock para os testes
type PoolInterface interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(context.Context, string, []interface{}, []interface{}, func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func GetConnection(context context.Context) *pgxpool.Pool {
	databaseURL := viper.GetString("database.url")

	conn, err := pgxpool.Connect(context, "postgres"+databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func RunMigration() {
	databaseURL := viper.GetString("database.url")
	m, err := migrate.New("file://database/migrations", "pgx"+databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
