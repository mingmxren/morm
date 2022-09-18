package morm

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type QueryxClient interface {
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
}

type NamedExecClient interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type ExecClient interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}
