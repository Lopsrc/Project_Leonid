package postgresql

import (
	"context"
	"fmt"
	"log"
	"pckg/postgresql/utils"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	//"github.com/jackc/pgx/v5/pgxpool"
)

//хранение конфигурационных данных
type StorageConf struct{
	username, password, host, port, database string
	maxAttempts int
}
type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions)(pgx.Tx, error)
	//BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error

}
//Подключение к БД ... Реализовать через структуру конфигурационные данные
func newClient(ctx context.Context, maxAttempts int, /*sc StorageConf*/ username, password, host, port, database string ) (pool *pgxpool.Pool,err error){

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", /* sc.username, sc.password, sc.host, sc.port, sc.database*/ username, 
		password, 
		host, 
		port, 
		database)

	// for maxAttempts > 0 {
	// 	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// 	defer cancel()

	// 	pool, err := pgxpool.Connect(ctx, dsn)
	// 	if err != nil {
	// 		fmt.Print("Failed to connect...")
	// 		return
	// 	}
	// 	return nil
	// }
	err = utils.DoWithTries(func() error{
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
	 	if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return pool, nil
		
}
