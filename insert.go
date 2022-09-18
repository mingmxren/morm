package morm

import (
	"context"
	"fmt"
	"strings"

	"github.com/mingmxren/morm/dbtag"
	"github.com/mingmxren/morm/options/insert"
)

func InsertSqlArgs(p PersistentObject, opts ...insert.Option) (sql string, err error) {
	qo := &insert.Options{}
	for _, opt := range opts {
		opt(qo)
	}

	names, _, err := dbtag.NameValuesWithTargetOrIgnore(p, qo.Fields, qo.IgnoreFields)
	if err != nil {
		return "", err
	}
	aliases := make([]string, len(names))
	for i, name := range names {
		aliases[i] = fmt.Sprintf(":%s", name)
	}

	sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		p.TableName(), strings.Join(names, ","), strings.Join(aliases, ","),
	)

	return sql, nil
}

func Insert(ctx context.Context, db NamedExecClient, ps []PersistentObject,
	opts ...insert.Option) (rowsAffected int64, err error) {
	sql, err := InsertSqlArgs(ps[0], opts...)
	if err != nil {
		return 0, err
	}
	result, err := db.NamedExecContext(ctx, sql, ps)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func InsertOne(ctx context.Context, db NamedExecClient, p PersistentObject, opts ...insert.Option) (rowsAffected int64,
	err error) {
	return Insert(ctx, db, []PersistentObject{p}, opts...)
}
