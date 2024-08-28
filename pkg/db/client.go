package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Client interface {
	Get(any, string, ...any) error
	Select(any, string, ...any) error
	ExecuteWithTx(func(*Tx) error) error
}

type client struct {
	db *sqlx.DB
}

func NewClient(cc ClientConfig) Client {
	driver := cc.GetDriver()
	datasource := cc.GetDatasource()

	db, err := sqlx.Connect(driver, datasource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occured during connect %s, driver %s, datasource %s\n", err.Error(), driver, datasource)
		os.Exit(1)
	}

	return &client{db: db}
}

func (c *client) Get(dest any, query string, queryParams ...any) error {
	return sqlx.Get(c.db, dest, query, queryParams...)
}

func (c *client) Select(dest any, query string, queryParams ...any) error {
	return sqlx.Select(c.db, dest, query, queryParams...)
}

type Tx struct {
	*sqlx.Tx
}

func (c *client) ExecuteWithTx(f func(tx *Tx) error) (err error) {
	x, err := c.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "error occured during execute %s\n", err.Error())
			if err := x.Rollback(); err != nil {
				fmt.Fprintf(os.Stderr, "error occured during transaction rollback, %s\n", err.Error())
			}
		} else {
			if err = x.Commit(); err != nil {
				fmt.Fprintf(os.Stderr, "error occured during transaction commit, %s\n", err.Error())
			}
		}
	}()

	err = f(&Tx{x})

	return
}
