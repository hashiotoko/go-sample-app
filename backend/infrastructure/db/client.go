package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/hashiotoko/go-sample-app/backend/config"
	"github.com/hashiotoko/go-sample-app/backend/database/sqlc"

	_ "github.com/go-sql-driver/mysql"
)

type Client interface {
	Conn() *sqlc.Queries
	WithTx(ctx context.Context, f func(ctx context.Context, tx *sqlc.Queries) error) error
	CloseDB()
}

type client struct {
	db   *sql.DB
	conn *sqlc.Queries
}

func NewClient() Client {
	slog.Info("Start Connecting to DataBase.....")
	dbCfg := config.Config.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		slog.Error("Failed to connect to DB", "error", err)
	}
	checkConnection(sqlDB, 3)

	if dbCfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	}

	slog.Info("Connected to DataBase!!")
	return &client{
		db:   sqlDB,
		conn: sqlc.New(sqlDB),
	}
}

func checkConnection(db *sql.DB, retry_count int) {
	if err := db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		retry_count--
		fmt.Printf("retry... retry_count:%v\n", retry_count)
		checkConnection(db, retry_count)
	}
}

func (c *client) Conn() *sqlc.Queries {
	return c.conn
}

func (c *client) WithTx(ctx context.Context, f func(ctx context.Context, tx *sqlc.Queries) error) (err error) {
	var tx *sql.Tx
	tx, err = c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				panic(rerr)
			}
			panic(p)
		}

		select {
		case <-ctx.Done():
			rerr := tx.Rollback()

			err = ctx.Err()
			if rerr != nil {
				err = errors.Join(err, rerr)
			}
		default:
			if err != nil {
				rerr := tx.Rollback()
				if rerr != nil {
					err = errors.Join(err, rerr)
				}
			} else {
				err = tx.Commit()
			}
		}
	}()
	qtx := c.conn.WithTx(tx)
	err = f(ctx, qtx)
	return err
}

func (c *client) CloseDB() {
	if c.conn == nil {
		return
	}

	err := c.conn.Close()
	if err != nil {
		panic(err)
	}
}
