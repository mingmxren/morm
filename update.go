package morm

import (
	"context"
	"fmt"
	"strings"

	"github.com/mingmxren/morm/condition"
	"github.com/mingmxren/morm/dbtag"
	"github.com/mingmxren/morm/limit"
	"github.com/mingmxren/morm/options/update"
)

func QuestionUpdateSqlArgs(p PersistentObject, opts ...update.Option) (sql string, args []interface{}, err error) {
	qo := &update.Options{}
	for _, opt := range opts {
		opt(qo)
	}

	fields, args, err := dbtag.NameValuesWithTargetOrIgnore(p, qo.Fields, qo.IgnoreFields)
	if err != nil {
		return "", nil, err
	}
	fieldPairs := make([]string, len(fields))
	for i, f := range fields {
		fieldPairs[i] = fmt.Sprintf("%s=?", f)
	}

	where := ""
	if qo.Condition != nil {
		where = fmt.Sprintf("WHERE %s", qo.Condition.Sql())
		args = append(args, qo.Condition.Args()...)
	}

	sql = fmt.Sprintf("UPDATE %s SET %s %s %s",
		p.TableName(), strings.Join(fieldPairs, ","), where, qo.Limit)

	return sql, args, nil
}

func Update(ctx context.Context, db ExecClient, p PersistentObject, opts ...update.Option) (rowsAffected int64,
	err error) {
	sql, args, err := QuestionUpdateSqlArgs(p, opts...)
	if err != nil {
		return 0, err
	}
	result, err := db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func UpdateByPrimaryKey(ctx context.Context, db ExecClient, p PersistentObjectWithPrimaryKey,
	opts ...update.Option) (rowsAffected int64, err error) {
	pkNames, pkValues, err := dbtag.NameValuesByFieldPtr(p, p.PrimaryKey()...)
	if err != nil {
		return 0, err
	}

	cond := make([]condition.Condition, 0, len(pkNames))
	for i, name := range pkNames {
		cond = append(cond, condition.Equal(name, pkValues[i]))
	}
	newOpts := make([]update.Option, 0)
	newOpts = append(newOpts,
		update.WithCondition(condition.And(cond...)),
		update.WithLimit(limit.LimitBy(2)),
		update.WithIgnoreFields(pkNames...),
	)

	newOpts = append(newOpts, opts...)
	return Update(ctx, db, p, newOpts...)
}
