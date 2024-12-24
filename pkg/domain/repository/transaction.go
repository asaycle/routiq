package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// TODO: database/sqlに依存しているので解消すること
type Tx interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	Commit() error
	Rollback() error
}

type TransactionManager interface {
	RunWithReadOnlyTx(ctx context.Context, fn func(ctx context.Context, tx Tx) error) error
	RunWithReadWriteTx(ctx context.Context, fn func(ctx context.Context, tx Tx) error) error
}

// MEMO: 本来はsqlx.DBがに依存してはいけないが、本実装と離れていると不便なのでここに置いておく
type TransactionManagerImpl struct {
	db *sqlx.DB
}

func NewTransactionManager(db *sqlx.DB) TransactionManager {
	return &TransactionManagerImpl{db: db}
}

func (t *TransactionManagerImpl) RunWithReadOnlyTx(ctx context.Context, fn func(ctx context.Context, tx Tx) error) error {
	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	return t.runWithTx(ctx, tx, fn)
}

func (t *TransactionManagerImpl) RunWithReadWriteTx(ctx context.Context, fn func(ctx context.Context, tx Tx) error) error {
	// トランザクションを開始
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}
	return t.runWithTx(ctx, tx, fn)
}

func (t *TransactionManagerImpl) runWithTx(ctx context.Context, tx Tx, fn func(ctx context.Context, tx Tx) error) error {
	if err := fn(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
