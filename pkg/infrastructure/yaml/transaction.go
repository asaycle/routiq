package yaml

import (
	"context"
	"database/sql"
)

type ReadOnlyTx struct{}

func (tx *ReadOnlyTx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return nil
}

func (tx *ReadOnlyTx) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return nil, nil
}

func (tx *ReadOnlyTx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return nil
}

func (tx *ReadOnlyTx) Commit() error {
	return nil
}

func (tx *ReadOnlyTx) Rollback() error {
	return nil
}
